/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg/fpc"
	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg/k8s"
)

func main() {
	startServer()
}

var (
	flagPort string
)

func init() {
	flag.StringVar(&flagPort, "port", "4000", "Port to listen on")
}

func startServer() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "x-user")

	r := gin.Default()
	r.Use(cors.New(config))

	// controller
	r.POST("/api/proposal", submitProposal)
	r.POST("/api/launch", launch)
	r.POST("/api/kill", kill)
	r.POST("/api/execute", execute)

	r.Run(":" + flagPort)
}

func submitProposal(c *gin.Context) {

	fmt.Println("new proposal")
	// submit new proposal via FPC

	// return answer
	c.IndentedJSON(http.StatusOK, "proposalSubmitted")
}

func launch(c *gin.Context) {

	fmt.Println("Launching worker ...")

	// create instance
	r, err := k8s.Launch()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Printf("K8s: %s\n", r)

	// TODO this json hack is not ideal but it works for now
	type attestation struct {
		PublicKey   string
		Attestation string
	}
	var a attestation

	// do attestation call worker/api/attestation
	fmt.Printf("Requesting attestation from worker ...\n")
	client := resty.New()
	_, err = client.R().SetResult(&a).Post("http://localhost:5000/api/attestation")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Printf("Attestation: %s\n", a)

	// TODO do some attestation evidence transformation here
	// Take the quote and send it to IAS
	// Return attestation evidence from IAS to client
	fmt.Printf("Perform attestation evidence transformation\n")

	// return answer
	c.IndentedJSON(http.StatusOK, a)
}

func kill(c *gin.Context) {
	fmt.Printf("Killing worker")
	resp, err := k8s.Kill()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, resp)
}

func execute(c *gin.Context) {

	fmt.Printf("Requesting evaluation pack from FPC Chaincode ...\n")
	// retrieve evaluation pack from FPC
	evaluationPack := fpc.FetchEvaluationPack()

	serializedEvaluationPack, err := json.MarshalIndent(evaluationPack, "", "    ")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	type ExecutionRequest struct {
		EncryptedEvaluationPack []byte `json:"data" binding:"required"`
	}

	fmt.Printf("Sending evaluation pack to worker and execute ...\n")

	// call worker/api/execute
	client := resty.New()
	resp, err := client.R().
		SetBody(ExecutionRequest{serializedEvaluationPack}).
		Post("http://localhost:5000/api/execute")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Printf("Received result: %s\n", string(resp.Body()))

	// return answer
	c.IndentedJSON(http.StatusOK, string(resp.Body()))
}

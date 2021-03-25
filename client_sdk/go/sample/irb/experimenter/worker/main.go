/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"

	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg"
	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	startServer()
}

var (
	flagPort     string
	privateKey   *ecdsa.PrivateKey
	publicKeyPem []byte
)

func init() {
	flag.StringVar(&flagPort, "port", "5000", "Port to listen on")

	// create public/private key pair
	fmt.Println("Worker started")
	fmt.Println("Create public/private key pair")
	createIdentity()

	fmt.Printf("%s\n", publicKeyPem)
}

func createIdentity() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	serializedPublicKey, err := x509.MarshalPKIXPublicKey(privateKey.Public())
	if err != nil {
		panic(err)
	}

	publicKeyPem = pem.EncodeToMemory(&pem.Block{
		Type:  "EC PUBLIC KEY",
		Bytes: serializedPublicKey,
	})
}

func startServer() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "x-user")

	r := gin.Default()
	r.Use(cors.New(config))

	// controller
	r.POST("/api/attestation", attestation)
	r.POST("/api/execute", execute)
	r.GET("/api/alive", alive)

	r.Run(":" + flagPort)
}

func attestation(c *gin.Context) {

	fmt.Println("run attestation")
	type attestation struct {
		PublicKey   string
		Attestation string
	}

	// TODO call graphene low-level attestation API
	// https://graphene.readthedocs.io/en/latest/attestation.html#low-level-dev-attestation-interface

	// hash the public key
	//hash := sha256.Sum256(publicKeyPem)

	// write the hash to /dev/attestation/user_report_data
	//err := ioutil.WriteFile("/dev/attestation/user_report_data", hash[:], 0644)
	//if err != nil {
	//	c.IndentedJSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	// read the quote from /dev/attestation/quote
	//quote, err := ioutil.ReadFile("/dev/attestation/quote")
	//if err != nil {
	//	c.IndentedJSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//quoteBase64 := base64.StdEncoding.EncodeToString(quote)

	quoteBase64 := "some attestation proof"
	a := attestation{
		PublicKey:   string(publicKeyPem),
		Attestation: quoteBase64,
	}

	// return answer
	c.IndentedJSON(http.StatusOK, a)
}

// Binding from JSON
type ExecutionRequest struct {
	EncryptedEvaluationPack []byte `json:"data" binding:"required"`
}

func execute(c *gin.Context) {

	fmt.Println("run execute")
	var req ExecutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// decrypt evaluation pack
	fmt.Println("Decrypting evaluation pack ...")
	// TODO
	evaluationPackPlaintext := req.EncryptedEvaluationPack

	// retrieve metadata from FPC
	evaluationPack, err := pkg.EvaluationPackFromJson(evaluationPackPlaintext)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	// for all items
	for _, item := range evaluationPack.Items {
		// load data from storage
		fmt.Printf("Downloading file %s\n", item.Handler)
		ciphertext := storage.Download(item.Handler)

		// decrypt data
		sk := item.SecretKey.KeyMaterial
		data := pkg.Decrypt(ciphertext, sk)
		fmt.Printf("Decrypting file %s\n", item.Handler)

		// store data on local storage
		// note that we write all items to input.jpg since the pytorch example only accepts a single input file.
		outputPath := filepath.Join(".", "input.jpg")
		err := ioutil.WriteFile(filepath.Clean(outputPath), data, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Stored downloaded data at %s\n", filepath.Clean(outputPath))
	}

	// trigger python script
	pythonFile := "pytorchexample.py"
	fmt.Printf("Executing %s ...\n", pythonFile)

	cmd := exec.Command("python", pythonFile)
	out, err := cmd.Output()
	fmt.Println(string(out))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	// read result.txt
	content, err := ioutil.ReadFile("result.txt")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Printf("Result: %s\n", content)

	// TODO sign result before returning

	// return answer
	c.IndentedJSON(http.StatusOK, string(content))
}

func alive(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "up and running!")
}

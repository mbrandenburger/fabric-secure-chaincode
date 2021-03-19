package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg"
	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg/fpc"
	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg/storage"
)

func main() {
	startServer()
}

var flagPort string

func init() {
	flag.StringVar(&flagPort, "port", "3000", "Port to listen on")
}

func startServer() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "x-user")

	r := gin.Default()
	r.Use(cors.New(config))

	// controller
	r.POST("/api/upload", upload)

	r.Run(":" + flagPort)
}

// Binding from JSON
type UploadRequest struct {
	Data       []byte   `json:"data" binding:"required"`
	DataName   string   `json:"dataName" binding:"required"`
	Domain     string   `json:"domain" binding:"required"`
	AllowedUse []string `json:"allowedUse" binding:"required"`
}

func upload(c *gin.Context) {

	var req UploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// encrypt with new random key
	sk := pkg.NewEncryptionKey()
	encryptedData := pkg.Encrypt(req.Data, sk)
	fmt.Printf("encrypt data: %s\n")

	// upload encrypted data
	handle := storage.Upload(encryptedData)
	fmt.Printf("stored under key: %s\n", handle)

	// create metadata
	dat := &pkg.PatientData{
		Metadata: &pkg.Metadata{
			Name:   req.DataName,
			Domain: req.Domain,
		},
		Consent: &pkg.Consent{
			AllowedUsageType: req.AllowedUse,
		},
		Handler: handle,
		SecretKey: &pkg.SecretKey{
			KeyMaterial: sk,
			KeyAlg:      "256AES-GCM",
		},
	}

	// register patient data at FPC IRB
	fpc.RegisterConsent(dat)

	// return answer
	c.IndentedJSON(http.StatusOK, "uploaded")
}

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg"
	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg/fpc"
	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg/storage"
)

func main() {

	// retrieve metadata from FPC
	evaluationPack := fpc.FetchEvaluationPack()

	// for all items
	for _, patientData := range evaluationPack {
		// load data from storage
		ciphertext := storage.Download(patientData.Handler)

		// decrypt data
		sk := patientData.SecretKey.KeyMaterial
		data := pkg.Decrypt(ciphertext, sk)

		// store data on local storage
		outputPath := filepath.Join(".", patientData.Metadata.Name)
		err := ioutil.WriteFile(filepath.Clean(outputPath), data, 0644)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Stored downloaded data at %s\n", filepath.Clean(outputPath))
	}

}

package fpc

import (
	"fmt"

	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg"
	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg/storage"
)

const dummyPatientId = "patientXYZ"

func RegisterConsent(data *pkg.PatientData) {

	jsonData, err := data.ToJson()
	if err != nil {
		panic(err)
	}

	patientId := dummyPatientId

	storage.Store(patientId, string(jsonData))
	fmt.Printf("patient data stored: %s\n", patientId)
}

func FetchEvaluationPack() []*pkg.PatientData {

	// for now we fetch just a single item

	patientId := dummyPatientId

	data := storage.Load(patientId)
	patientData, err := pkg.FromJson([]byte(data))
	if err != nil {
		panic(err)
	}

	return []*pkg.PatientData{patientData}
}

func ProposeExperiment(desc *pkg.Experiment) {

}

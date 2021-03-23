/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fpc

import (
	"fmt"

	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg"
	"github.com/hyperledger-labs/fabric-private-chaincode/client_sdk/go/sample/irb/pkg/storage"
)

const dummyItemtId = "patientXYZ"

func RegisterConsent(data *pkg.ConsentData) {

	jsonData, err := data.ToJson()
	if err != nil {
		panic(err)
	}

	patientId := dummyItemtId

	fmt.Printf("Registering consent with id: %s\n", patientId)
	storage.Store(patientId, string(jsonData))
}

func FetchEvaluationPack() *pkg.EvaluationPack {

	// for now we fetch just a single item

	itemId := dummyItemtId

	data := storage.Load(itemId)
	item, err := pkg.ConsentDataFromJson([]byte(data))
	if err != nil {
		panic(err)
	}

	return &pkg.EvaluationPack{
		Items: []*pkg.ConsentData{item},
	}
}

func ProposeExperiment(desc *pkg.Experiment) {

}

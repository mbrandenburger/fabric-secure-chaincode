/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package ercc

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"

	"github.com/hyperledger-labs/fabric-private-chaincode/ercc/attestation"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var spid = [16]byte{0x25, 0x42, 0xF9, 0x0D, 0x63, 0x31, 0x2C, 0x1D, 0xA2, 0xF6, 0xE2, 0x18, 0x76, 0xF0, 0xE3, 0x89}

// MockEnclaveRegistryStub implements EnclaveRegistryStub interface and calls ercc
type MockEnclaveRegistryStub struct {
}

// GetSPID return SPID from ercc
func (t *MockEnclaveRegistryStub) GetSPID(stub shim.ChaincodeStubInterface, chaincodeName, channel string) ([]byte, error) {
	return spid[:], nil
}

// RegisterEnclave registers enclave at ercc
func (t *MockEnclaveRegistryStub) RegisterEnclave(stub shim.ChaincodeStubInterface, chaincodeName, channel string, enclavePk, enclaveQuote []byte) error {
	//fmt.Println("Register: " + base64.StdEncoding.EncodeToString(enclaveID) + " : " + base64.StdEncoding.EncodeToString(enclaveQuote))

	attestationReport := attestation.IASAttestationReport{
		EnclavePk:                   enclavePk,
		IASReportSignature:          "dummySignature",
		IASReportSigningCertificate: "dummyCert",
		IASReportBody:               []byte("dummyBody"),
	}

	// store attestation report under enclavePk hash in state
	attestationReportAsBytes, err := json.Marshal(attestationReport)
	if err != nil {
		return err
	}

	// create hash of enclave pk
	enclavePkHash := sha256.Sum256(enclavePk)
	enclavePkHashBase64 := base64.StdEncoding.EncodeToString(enclavePkHash[:])
	err = stub.PutState(enclavePkHashBase64, attestationReportAsBytes)

	return nil
}

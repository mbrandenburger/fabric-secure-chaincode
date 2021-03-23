/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package pkg

import "encoding/json"

type ConsentData struct {
	Metadata  *Metadata
	Consent   *Consent
	Handler   string
	SecretKey *SecretKey
}

type Metadata struct {
	Name   string
	Domain string
}

type Consent struct {
	AllowedUsageType []string
}

type SecretKey struct {
	KeyMaterial []byte
	KeyAlg      string
}

func (pd *ConsentData) ToJson() ([]byte, error) {
	return json.MarshalIndent(pd, "", "    ")
}

type EvaluationPack struct {
	Items []*ConsentData
}

func EvaluationPackFromJson(data []byte) (*EvaluationPack, error) {
	pack := &EvaluationPack{}
	err := json.Unmarshal(data, pack)
	if err != nil {
		return nil, err
	}

	return pack, nil
}

func ConsentDataFromJson(data []byte) (*ConsentData, error) {
	consent := &ConsentData{}
	err := json.Unmarshal(data, consent)
	if err != nil {
		return nil, err
	}

	return consent, nil
}

type Experiment struct {
	Id            string
	Title         string
	Description   string
	TypeInputData []string
	Manifest      *Manifest
	Approvals     []*Approval
}

type Manifest struct {
	CodeIdentity string
}

type Approval struct {
	ExperimentId   string
	Signature      []byte
	SignerIdentity string
}

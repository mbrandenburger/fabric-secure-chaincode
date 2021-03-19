package pkg

import "encoding/json"

type PatientData struct {
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

func (pd *PatientData) ToJson() ([]byte, error) {
	return json.MarshalIndent(pd, "", "    ")
}

func FromJson(data []byte) (*PatientData, error) {
	patientData := &PatientData{}
	err := json.Unmarshal(data, patientData)
	if err != nil {
		return nil, err
	}

	return patientData, nil
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

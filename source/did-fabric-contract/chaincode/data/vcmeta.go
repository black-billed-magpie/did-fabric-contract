// Copyright 2024 Raonsecure

package data

const (
	VCMETA_PREFIX = "open:vcmeta:"
)

type VC_STATUS string

const (
	VC_ACTIVE   VC_STATUS = "ACTIVE"
	VC_INACTIVE VC_STATUS = "INACTIVE"
	VC_REVOKED  VC_STATUS = "REVOKED"
)

type VcMeta struct {
	Id               string           `validate:"required" json:"id"`
	Issuer           Provider         `validate:"required" json:"issuer"`
	Subject          string           `validate:"required" json:"subject"`
	CredentialSchema CredentialSchema `validate:"required" json:"credentialSchema"`
	Status           VC_STATUS        `validate:"required" json:"status"`
	IssuanceDate     string           `validate:"required" json:"issuanceDate"`
	ValidFrom        string           `validate:"required" json:"validFrom"`
	ValidUntil       string           `validate:"required" json:"validUntil"`
	FormatVersion    string           `validate:"required" json:"formatVersion"`
	Language         string           `validate:"required" json:"language"`
}

func (v VcMeta) Key() ([]string, error) {
	return []string{VCMETA_PREFIX, v.Id}, nil
}
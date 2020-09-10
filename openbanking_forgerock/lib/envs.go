package lib

import (
	"fmt"
	"os"
)

// Envs stores all the required envs.
type Envs struct {
	SSID           string
	OrganisationID string
	SSA            string
	KID            string
}

// GetEnvs returns an `Envs` struct pre-populated from the current environment.
func GetEnvs() (*Envs, error) {
	SSID := os.Getenv("SSID")
	organisationID := os.Getenv("ORGANISATION_ID")
	SSA := os.Getenv("SSA")
	KID := os.Getenv("KID")

	envs := &Envs{
		SSID:           SSID,
		OrganisationID: organisationID,
		SSA:            SSA,
		KID:            KID,
	}

	if err := envs.validate(); err != nil {
		return nil, err
	}

	return envs, nil
}

// TODO: better error logging, log the purpose of the variable.
func (e *Envs) validate() error {
	if e.SSID == "" {
		return missingEnvError("SSID", "Software Statement ID")
	}
	if e.OrganisationID == "" {
		return missingEnvError("ORGANISATION_ID", "Organisation ID")
	}
	if e.SSA == "" {
		return missingEnvError("SSA", "Software Statement Assertion")
	}
	if e.KID == "" {
		return missingEnvError("KID", "Key ID")
	}

	return nil
}

// TODO: better error logging take in the purpose for the env. variable.
func missingEnvError(varName, varPurpose string) error {
	return fmt.Errorf("missing %s environmental variable: %s=<value>", varPurpose, varName)
}

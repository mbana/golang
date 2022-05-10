package config

import (
	"io/ioutil"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Config holds configuration details.
//
// We specify both a `yaml` and `json` even though the config is in `yaml` because specifying a json tag for the field
// displays better errors messages.
// TODO(mbana): enhance and document better.
type Config struct {
	SSID                    string `yaml:"SSID" json:"SSID"`
	OrganisationID          string `yaml:"ORGANISATION_ID" json:"ORGANISATION_ID"`
	SSA                     string `yaml:"SSA" json:"SSA"`
	KID                     string `yaml:"KID" json:"KID"`
	ClientID                string `yaml:"client_id" json:"client_id"`
	RequestObjectSigningAlg string `yaml:"request_object_signing_alg" json:"request_object_signing_alg"`
	EncryptionPrivate       string `yaml:"encryption_private" json:"encryption_private"`
	EncryptionPublic        string `yaml:"encryption_public" json:"encryption_public"`
	SignaturePrivate        string `yaml:"signature_private" json:"signature_private"`
	SignaturePublic         string `yaml:"signature_public" json:"signature_public"`
	TransportPrivate        string `yaml:"transport_private" json:"transport_private"`
	TransportPublic         string `yaml:"transport_public" json:"transport_public"`
	RegisterResponse        string `yaml:"register_response" json:"register_response"`
}

func NewConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "failed on ioutil.ReadFile in NewConfig: filename=%q", filename)
	}

	config := &Config{}
	if err := yaml.Unmarshal(file, config); err != nil {
		return nil, errors.Wrapf(err, "failed on yaml.Unmarshal in NewConfig: filename=%q", filename)
	}

	if err := config.Validate(); err != nil {
		return nil, errors.Wrapf(err, "failed on config.Validate in NewConfig: filename=%q", filename)
	}

	return config, nil
}

func (c Config) Validate() error {
	requestObjectSigningAlgValuesSupported := RequestObjectSigningAlgValuesSupported.ToInterface()
	return validation.ValidateStruct(&c,
		validation.Field(&c.SSID, validation.Required),
		validation.Field(&c.OrganisationID, validation.Required),
		validation.Field(&c.SSA, validation.Required),
		validation.Field(&c.KID, validation.Required),
		validation.Field(&c.ClientID, validation.Required),
		validation.Field(&c.RequestObjectSigningAlg, validation.Required, validation.In(requestObjectSigningAlgValuesSupported...)),
		validation.Field(&c.EncryptionPrivate, validation.Required, filePath),
		validation.Field(&c.EncryptionPublic, validation.Required, filePath),
		validation.Field(&c.SignaturePrivate, validation.Required, filePath),
		validation.Field(&c.SignaturePublic, validation.Required, filePath),
		validation.Field(&c.TransportPrivate, validation.Required, filePath),
		validation.Field(&c.TransportPublic, validation.Required, filePath),
		validation.Field(&c.RegisterResponse, validation.Required, filePath),
	)
}

// func (c Config) String() string {
// 	buffer := bytes.NewBufferString("")
// 	fmt.Fprintf(buffer, "Config\n")
// 	fmt.Fprintf(buffer, "  %s: %s\n", blue("KID"), red(c.KID))
// 	fmt.Fprintf(buffer, "  %s: %s\n", blue("RequestObjectSigningAlg"), red(c.RequestObjectSigningAlg))
// 	return strings.TrimSuffix(buffer.String(), "\n")
// }

type RequestObjectSigningAlgs []string

func (r RequestObjectSigningAlgs) ToInterface() []interface{} {
	clone := make([]interface{}, len(r))
	for i, v := range r {
		clone[i] = v
	}
	return clone
}

var (
	// RequestObjectSigningAlgValuesSupported see `oidcdiscovery/well-known_openid-configuration.go`.
	// nolint:gochecknoglobals
	RequestObjectSigningAlgValuesSupported = RequestObjectSigningAlgs{
		"PS384",
		"ES384",
		"RS384",
		"HS256",
		"HS512",
		"ES256",
		"RS256",
		"HS384",
		"ES512",
		"PS256",
		"PS512",
		"RS512",
	}
)

var (
	// FilePath validates if a string is an file path or not.
	filePath = validation.NewStringRule(func(name string) bool {
		if _, err := os.Stat(name); err != nil {
			if os.IsNotExist(err) {
				return false
			}
		}
		return true
	}, "must be a valid file path")
)

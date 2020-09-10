package lib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// OpenIDConfig ...
//
// https://as.aspsp.ob.forgerock.financial/oauth2/.well-known/openid-configuration
//
// {
//     "request_parameter_supported": true,
//     "claims_parameter_supported": true,
//     "introspection_endpoint": "https://as.aspsp.ob.forgerock.financial/oauth2/realms/root/realms/openbanking/introspect",
//     "check_session_iframe": "https://as.aspsp.ob.forgerock.financial/oauth2/realms/root/realms/openbanking/connect/checkSession",
//     "scopes_supported": [
//         "openid",
//         "payments",
//         "accounts"
//     ],
//     "issuer": "https://as.aspsp.ob.forgerock.financial/oauth2/openbanking",
//     "id_token_encryption_enc_values_supported": [
//         "A256GCM",
//         "A192GCM",
//         "A128GCM",
//         "A128CBC-HS256",
//         "A192CBC-HS384",
//         "A256CBC-HS512"
//     ],
//     "acr_values_supported": [
//         "urn:openbanking:psd2:sca",
//         "urn:openbanking:psd2:ca"
//     ],
//     "authorization_endpoint": "https://as.aspsp.ob.forgerock.financial/oauth2/realms/root/realms/openbanking/authorize",
//     "request_object_encryption_enc_values_supported": [
//         "A256GCM",
//         "A192GCM",
//         "A128GCM",
//         "A128CBC-HS256",
//         "A192CBC-HS384",
//         "A256CBC-HS512"
//     ],
//     "rcs_request_encryption_alg_values_supported": [
//         "RSA-OAEP",
//         "RSA-OAEP-256",
//         "A128KW",
//         "RSA1_5",
//         "A256KW",
//         "dir",
//         "A192KW"
//     ],
//     "claims_supported": [
//         "acr",
//         "zoneinfo",
//         "openbanking_intent_id",
//         "address",
//         "profile",
//         "name",
//         "phone_number",
//         "given_name",
//         "locale",
//         "family_name",
//         "email"
//     ],
//     "rcs_request_signing_alg_values_supported": [
//         "ES384",
//         "HS256",
//         "HS512",
//         "ES256",
//         "RS256",
//         "HS384",
//         "ES512"
//     ],
//     "token_endpoint_auth_methods_supported": [
//         "client_secret_post",
//         "private_key_jwt",
//         "client_secret_basic"
//     ],
//     "token_endpoint": "https://as.aspsp.ob.forgerock.financial/oauth2/realms/root/realms/openbanking/access_token",
//     "response_types_supported": [
//         "code",
//         "code id_token",
//         "id_token"
//     ],
//     "request_uri_parameter_supported": true,
//     "rcs_response_encryption_enc_values_supported": [
//         "A256GCM",
//         "A192GCM",
//         "A128GCM",
//         "A128CBC-HS256",
//         "A192CBC-HS384",
//         "A256CBC-HS512"
//     ],
//     "end_session_endpoint": "https://as.aspsp.ob.forgerock.financial/oauth2/realms/root/realms/openbanking/connect/endSession",
//     "rcs_request_encryption_enc_values_supported": [
//         "A256GCM",
//         "A192GCM",
//         "A128GCM",
//         "A128CBC-HS256",
//         "A192CBC-HS384",
//         "A256CBC-HS512"
//     ],
//     "version": "3.0",
//     "rcs_response_encryption_alg_values_supported": [
//         "RSA-OAEP",
//         "RSA-OAEP-256",
//         "A128KW",
//         "A256KW",
//         "RSA1_5",
//         "dir",
//         "A192KW"
//     ],
//     "userinfo_endpoint": "https://as.aspsp.ob.forgerock.financial/oauth2/realms/root/realms/openbanking/userinfo",
//     "id_token_encryption_alg_values_supported": [
//         "RSA-OAEP",
//         "RSA-OAEP-256",
//         "A128KW",
//         "A256KW",
//         "RSA1_5",
//         "dir",
//         "A192KW"
//     ],
//     "jwks_uri": "https://as.aspsp.ob.forgerock.financial/oauth2/realms/root/realms/openbanking/connect/jwk_uri",
//     "subject_types_supported": [
//         "public",
//         "pairwise"
//     ],
//     "id_token_signing_alg_values_supported": [
//         "ES384",
//         "HS256",
//         "HS512",
//         "ES256",
//         "RS256",
//         "HS384",
//         "ES512"
//     ],
//     "registration_endpoint": "https://as.aspsp.ob.forgerock.financial/oauth2/realms/root/realms/openbanking/register",
//     "request_object_signing_alg_values_supported": [
//         "ES384",
//         "HS256",
//         "HS512",
//         "ES256",
//         "RS256",
//         "HS384",
//         "ES512"
//     ],
//     "request_object_encryption_alg_values_supported": [
//         "RSA-OAEP",
//         "RSA-OAEP-256",
//         "A128KW",
//         "RSA1_5",
//         "A256KW",
//         "dir",
//         "A192KW"
//     ],
//     "rcs_response_signing_alg_values_supported": [
//         "ES384",
//         "HS256",
//         "HS512",
//         "ES256",
//         "RS256",
//         "HS384",
//         "ES512"
//     ]
// }
type OpenIDConfig struct {
	RequestParameterSupported                 bool     `json:"request_parameter_supported"`
	ClaimsParameterSupported                  bool     `json:"claims_parameter_supported"`
	RequestURIParameterSupported              bool     `json:"request_uri_parameter_supported"`
	IntrospectionEndpoint                     string   `json:"introspection_endpoint"`
	CheckSessionIframe                        string   `json:"check_session_iframe"`
	Issuer                                    string   `json:"issuer"`
	AuthorizationEndpoint                     string   `json:"authorization_endpoint"`
	TokenEndpoint                             string   `json:"token_endpoint"`
	EndSessionEndpoint                        string   `json:"end_session_endpoint"`
	Version                                   string   `json:"version"`
	UserinfoEndpoint                          string   `json:"userinfo_endpoint"`
	JwksURI                                   string   `json:"jwks_uri"`
	RegistrationEndpoint                      string   `json:"registration_endpoint"`
	ScopesSupported                           []string `json:"scopes_supported"`
	IDTokenEncryptionEncValuesSupported       []string `json:"id_token_encryption_enc_values_supported"`
	AcrValuesSupported                        []string `json:"acr_values_supported"`
	RequestObjectEncryptionEncValuesSupported []string `json:"request_object_encryption_enc_values_supported"`
	RcsRequestEncryptionAlgValuesSupported    []string `json:"rcs_request_encryption_alg_values_supported"`
	ClaimsSupported                           []string `json:"claims_supported"`
	RcsRequestSigningAlgValuesSupported       []string `json:"rcs_request_signing_alg_values_supported"`
	TokenEndpointAuthMethodsSupported         []string `json:"token_endpoint_auth_methods_supported"`
	ResponseTypesSupported                    []string `json:"response_types_supported"`
	RcsResponseEncryptionEncValuesSupported   []string `json:"rcs_response_encryption_enc_values_supported"`
	RcsRequestEncryptionEncValuesSupported    []string `json:"rcs_request_encryption_enc_values_supported"`
	RcsResponseEncryptionAlgValuesSupported   []string `json:"rcs_response_encryption_alg_values_supported"`
	IDTokenEncryptionAlgValuesSupported       []string `json:"id_token_encryption_alg_values_supported"`
	SubjectTypesSupported                     []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported          []string `json:"id_token_signing_alg_values_supported"`
	RequestObjectSigningAlgValuesSupported    []string `json:"request_object_signing_alg_values_supported"`
	RequestObjectEncryptionAlgValuesSupported []string `json:"request_object_encryption_alg_values_supported"`
	RcsResponseSigningAlgValuesSupported      []string `json:"rcs_response_signing_alg_values_supported"`
}

// GetOpenIDConfig ...
func GetOpenIDConfig() (*OpenIDConfig, error) {
	url := "https://as.aspsp.ob.forgerock.financial/oauth2/.well-known/openid-configuration"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	config := &OpenIDConfig{}
	if err := json.NewDecoder(res.Body).Decode(&config); err != nil {
		return nil, err

	}

	logrus.WithFields(logrus.Fields{
		"config": fmt.Sprintf("%+v", config),
	}).Debug("GetConfig")

	return config, nil
}

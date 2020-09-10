package oidcdynamicclientregistration

// Response ...
// https://openid.net/specs/openid-connect-registration-1_0-21.html#ClientMetadata
type Response struct {
	Scope                        string   `json:"scope"`
	ApplicationType              string   `json:"application_type"`
	JwksURI                      string   `json:"jwks_uri"`
	SubjectType                  string   `json:"subject_type"`
	IDTokenSignedResponseAlg     string   `json:"id_token_signed_response_alg"`
	IDTokenEncryptedResponseAlg  string   `json:"id_token_encrypted_response_alg"`
	IDTokenEncryptedResponseEnc  string   `json:"id_token_encrypted_response_enc"`
	UserinfoSignedResponseAlg    string   `json:"userinfo_signed_response_alg"`
	UserinfoEncryptedResponseAlg string   `json:"userinfo_encrypted_response_alg"`
	RequestObjectSigningAlg      string   `json:"request_object_signing_alg"`
	RequestObjectEncryptionAlg   string   `json:"request_object_encryption_alg"`
	TokenEndpointAuthMethod      string   `json:"token_endpoint_auth_method"`
	TokenEndpointAuthSigningAlg  string   `json:"token_endpoint_auth_signing_alg"`
	DefaultMaxAge                string   `json:"default_max_age"`
	SoftwareStatement            string   `json:"software_statement"`
	ClientID                     string   `json:"client_id"`
	ClientSecret                 string   `json:"client_secret"`
	RegistrationAccessToken      string   `json:"registration_access_token"`
	RegistrationClientURI        string   `json:"registration_client_uri"`
	ClientSecretExpiresAt        string   `json:"client_secret_expires_at"`
	Scopes                       []string `json:"scopes"`
	RedirectUris                 []string `json:"redirect_uris"`
	ResponseTypes                []string `json:"response_types"`
}

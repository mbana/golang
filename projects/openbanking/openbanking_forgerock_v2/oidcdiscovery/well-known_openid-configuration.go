package oidcdiscovery

// OpenIDConfiguration OpenID Connect Discovery 1.0 incorporating errata set 1
// https://openid.net/specs/openid-connect-discovery-1_0.html
// https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderConfigurationResponse
//
// TODO(mbana): Should probably use
// https://github.com/ory/hydra/blob/957a2d670a4be8c6e1a30b2df222fc566e13b8a1/oauth2/doc.go#L31
type OpenIDConfiguration struct {
	RequestParameterSupported                  bool     `json:"request_parameter_supported"`
	ClaimsParameterSupported                   bool     `json:"claims_parameter_supported"`
	RequestURIParameterSupported               bool     `json:"request_uri_parameter_supported"`
	IntrospectionEndpoint                      string   `json:"introspection_endpoint"`
	Issuer                                     string   `json:"issuer"`
	AuthorizationEndpoint                      string   `json:"authorization_endpoint"`
	TokenEndpoint                              string   `json:"token_endpoint"`
	Version                                    string   `json:"version"`
	UserinfoEndpoint                           string   `json:"userinfo_endpoint"`
	JwksURI                                    string   `json:"jwks_uri"`
	RegistrationEndpoint                       string   `json:"registration_endpoint"`
	GrantTypesSupported                        []string `json:"grant_types_supported"`
	ScopesSupported                            []string `json:"scopes_supported"`
	IDTokenEncryptionEncValuesSupported        []string `json:"id_token_encryption_enc_values_supported"`
	AcrValuesSupported                         []string `json:"acr_values_supported"`
	RequestObjectEncryptionEncValuesSupported  []string `json:"request_object_encryption_enc_values_supported"`
	ClaimsSupported                            []string `json:"claims_supported"`
	TokenEndpointAuthMethodsSupported          []string `json:"token_endpoint_auth_methods_supported"`
	ResponseTypesSupported                     []string `json:"response_types_supported"`
	IDTokenEncryptionAlgValuesSupported        []string `json:"id_token_encryption_alg_values_supported"`
	SubjectTypesSupported                      []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported           []string `json:"id_token_signing_alg_values_supported"`
	RequestObjectSigningAlgValuesSupported     []string `json:"request_object_signing_alg_values_supported"`
	RequestObjectEncryptionAlgValuesSupported  []string `json:"request_object_encryption_alg_values_supported"`
	UserInfoSigningAlgValuesSupported          []string `json:"userinfo_signing_alg_values_supported"`
	UserInfoEncryptionEncValuesSupported       []string `json:"userinfo_encryption_enc_values_supported"`
	UserInfoEncryptionAlgValuesSupported       []string `json:"userinfo_encryption_alg_values_supported"`
	TokenEndpointAuthSigningAlgValuesSupported []string `json:"token_endpoint_auth_signing_alg_values_supported"`
	// CheckSessionIframe                         string   `json:"check_session_iframe"`
	// EndSessionEndpoint                         string   `json:"end_session_endpoint"`
	// RcsRequestEncryptionEncValuesSupported     []string `json:"rcs_request_encryption_enc_values_supported"`
	// RcsRequestSigningAlgValuesSupported        []string `json:"rcs_request_signing_alg_values_supported"`
	// RcsResponseEncryptionAlgValuesSupported    []string `json:"rcs_response_encryption_alg_values_supported"`
	// RcsResponseEncryptionEncValuesSupported    []string `json:"rcs_response_encryption_enc_values_supported"`
	// RcsResponseSigningAlgValuesSupported       []string `json:"rcs_response_signing_alg_values_supported"`
	// RcsRequestEncryptionAlgValuesSupported     []string `json:"rcs_request_encryption_alg_values_supported"`
}

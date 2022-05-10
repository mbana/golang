package oidcdynamicclientregistration

// Config - https://openid.net/specs/openid-connect-registration-1_0-21.html#Terminology
type Config struct {
	// OAuth 2.0 Protected Resource through which a Client can be registered at an Authorization Server.
	ClientRegistrationEndpoint string
	// OAuth 2.0 Endpoint through which registration information for a registered Client can be managed. This URL for this endpoint is returned by the Authorization Server in the Client Information Response.
	ClientConfigurationEndpoint string
	// OAuth 2.0 Bearer Token issued by the Authorization Server through the Client Registration Endpoint that is used to authenticate the caller when accessing the Client's registration information at the Client Configuration Endpoint. This Access Token is associated with a particular registered Client.
	RegistrationAccessToken string
	// OAuth 2.0 Access Token optionally issued by an Authorization Server granting access to its Client Registration Endpoint. The contents of this token are service-specific and are out of scope for this specification. The means by which the Authorization Server issues this token and the means by which the Registration Endpoint validates it are also out of scope.
	InitialAccessToken string
}

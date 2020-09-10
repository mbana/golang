package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// RegisterTPPResponse example
//
// {
//     "scopes": [
//         "openid",
//         "payments",
//         "accounts"
//     ],
//     "scope": "openid payments accounts",
//     "redirect_uris": [
//         "https://bana.io/openbanking/forgerock"
//     ],
//     "response_types": [
//         "code"
//     ],
//     "application_type": "web",
//     "jwks_uri": "https://service.directory.ob.forgerock.financial/directory-services/api/software-statement/5b5a2008b093465496d238fc/application/jwk_uri",
//     "subject_type": "public",
//     "id_token_signed_response_alg": "HS256",
//     "id_token_encrypted_response_alg": "RSA1_5",
//     "id_token_encrypted_response_enc": "A128CBC_HS256",
//     "userinfo_signed_response_alg": "",
//     "userinfo_encrypted_response_alg": "",
//     "request_object_signing_alg": "RS256",
//     "request_object_encryption_alg": "",
//     "token_endpoint_auth_method": "private_key_jwt",
//     "token_endpoint_auth_signing_alg": "RS256",
//     "default_max_age": "1",
//     "software_statement": "...",
//     "client_id": "...",
//     "client_secret": "...",
//     "registration_access_token": "...",
//     "registration_client_uri": "...",
//     "client_secret_expires_at": "0"
// }
type RegisterTPPResponse struct {
	Scopes                       []string `json:"scopes"`
	Scope                        string   `json:"scope"`
	RedirectUris                 []string `json:"redirect_uris"`
	ResponseTypes                []string `json:"response_types"`
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
}

const (
	// RegisterResponseFile the file to save the response from register.
	RegisterResponseFile = ".ignore/register_response.json"
)

// Register ..
func (c *OpenBankingClient) Register() (*RegisterTPPResponse, error) {
	now := time.Now()
	iat := now.Unix()
	// exp := now.Add(24 * time.Hour).Unix()
	exp := time.Date(2019, 03, 29, 0, 0, 0, 0, time.UTC).Unix()
	jti := uuid.New().String()

	claims := jwt.MapClaims{
		"token_endpoint_auth_signing_alg": "RS256",
		"request_object_encryption_alg":   "RSA-OAEP-256",
		"grant_types": []string{
			"authorization_code",
			"refresh_token",
			"client_credentials",
		},
		"subject_type":     "public",
		"application_type": "web",
		"iss":              c.Envs.SSID,
		"redirect_uris": []string{
			// "https://bana.io/openbanking/forgerock",
			"http://localhost:8080/openbanking/banaio/forgerock",
		},
		"token_endpoint_auth_method": "private_key_jwt",
		"aud":                        c.OpenIDConfig.Issuer,
		"software_statement":         c.Envs.SSA,
		"scopes": []string{
			"openid",
			"accounts",
			"payments",
		},
		"request_object_signing_alg":    "RS256",
		"exp":                           exp,
		"request_object_encryption_enc": "A128CBC-HS256",
		"iat":                           iat,
		"jti":                           jti,
		"response_types": []string{
			"code",
			"code id_token",
			"id_token",
		},
		"id_token_signed_response_alg": "ES256",
	}
	jwt := Sign(claims, c.Envs.KID)
	url := c.OpenIDConfig.RegistrationEndpoint

	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Info("Register")

	resp, err := c.HTTPClient.Post(url, "application/jwt", bytes.NewBufferString(jwt))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
			"jwt": jwt,
		}).Warn("Register")
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 201 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"StatusCode": resp.StatusCode,
				"err":        err,
			}).Warn("Register:ReadAll")
			return nil, err
		}

		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"body":       string(body),
		}).Warn("Register:ReadAll")
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(c.RegisterResponse); err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"err":        err,
		}).Warn("Register:Decode")
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"StatusCode": resp.StatusCode,
		"client_id":  c.RegisterResponse.ClientID,
	}).Info("Register")

	bytes, err := json.Marshal(c.RegisterResponse)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"err":        err,
		}).Warn("Register:Marshal")
		return nil, err
	}

	if err := ioutil.WriteFile(RegisterResponseFile, bytes, 0644); err != nil {
		logrus.WithFields(logrus.Fields{
			"RegisterResponseFile": err,
			"bytes":                string(bytes),
			"err":                  err,
		}).Warn("Register:ioutil.WriteFile")
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"RegisterResponseFile": RegisterResponseFile,
		"length":               len(bytes),
	}).Info("Register:WriteFile")

	return c.RegisterResponse, nil
}

// UnRegister ...
func (c *OpenBankingClient) UnRegister() {
	claims := jwt.MapClaims{
		"grant_types": []string{
			"authorization_code",
			"refresh_token",
			"client_credentials",
		},
		// "request_object_encryption_alg": "RSA-OAEP-256",
		// "subject_type":                  "public",
		// "application_type":              "web",
		"redirect_uris": []string{
			"https://bana.io/openbanking/forgerock",
			"http://localhost:8080/openbanking/banaio",
			"http://localhost:8080/openbanking/banaio/forgerock",
		},
		"scopes": []string{
			"openid",
			"accounts",
			"payments",
		},
		"token_endpoint_auth_method": "private_key_jwt",
		"request_object_signing_alg": "RS256",
		"software_statement":         c.Envs.SSA,
		"iss":                        c.Envs.SSID,
		"aud":                        c.OpenIDConfig.Issuer,
	}
	jwt := Sign(claims, c.Envs.KID)
	url := "https://rs.aspsp.ob.forgerock.financial/open-banking/registerTPP"

	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Info("UnRegister")

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBufferString(jwt))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("UnRegister:NewRequest")
	}

	req.Header.Set("Content-Type", "application/jwt")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"err":        err,
		}).Error("UnRegister:Do")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"err":        err,
		}).Error("UnRegister:ReadAll")
	}

	if resp.StatusCode != 200 {
		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"Body":       string(body),
		}).Error("UnRegister")
	}

	logrus.WithFields(logrus.Fields{
		"StatusCode": resp.StatusCode,
		"Body":       string(body),
	}).Info("UnRegister")
}

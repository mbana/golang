package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// AccessTokenResponse example
//
// {
//     "access_token": "...",
//     "scope": "openid payments accounts",
//     "id_token": "...",
//     "token_type": "Bearer",
//     "expires_in": 86399
// }
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	IDToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

// GetAccessToken gets an `access_token` from the
// `https://as.aspsp.ob.forgerock.financial/oauth2/realms/root/realms/openbanking/access_token`
// endpoint and sets `OpenBankingClient.AccessToken` to the obtained `access_token`.
//
// see: https://openbanking.atlassian.net/wiki/spaces/DZ/pages/187793608/Integrating+a+TPP+with+ForgeRock+Model+Bank+on+Directory+Sandbox#IntegratingaTPPwithForgeRockModelBankonDirectorySandbox-3.1GetanaccesstokentorepresentyouasaTPPusingtheClientcredentialflow.
func (c *OpenBankingClient) GetAccessToken() error {
	now := time.Now()
	iat := now.Unix()
	// "The expiration time. After this time, this JWT won't be considered a valid credential. For security reasons, we recommend you set a short period of life, such as 1 or 2 minutes."
	// 30 minutes seems to work, anything larger fails.
	exp := now.Add(30 * time.Minute).Unix()
	// exp := now.Add(2 * time.Minute).Unix()
	// exp := time.Date(2019, 03, 29, 0, 0, 0, 0, time.UTC).Unix()
	jti := uuid.New().String()
	claims := jwt.MapClaims{
		"iss": c.RegisterResponse.ClientID,
		"sub": c.RegisterResponse.ClientID,
		"aud": c.OpenIDConfig.Issuer,
		"iat": iat,
		"exp": exp,
		"jti": jti,
	}

	clientAssertion := Sign(claims, c.Envs.KID)
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	data.Add("scope", "openid accounts payments")
	data.Add("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	data.Add("client_assertion", clientAssertion)

	req, err := http.NewRequest("POST", c.OpenIDConfig.TokenEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":  err,
			"req":  fmt.Sprintf("%+v", req),
			"data": data,
		}).Error("GetAccessToken:NewRequest")
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cache-Control", "no-cache")

	RequestToCurlCommand(req, "GetAccessToken")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode":       resp.StatusCode,
			"err":              err,
			"resp":             resp,
			"data":             data,
			"client_assertion": clientAssertion,
		}).Error("GetAccessToken:Do")
		return err
	}

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"StatusCode": resp.StatusCode,
				"err":        err,
			}).Error("GetAccessToken:ReadAll")
			return err
		}
		defer resp.Body.Close()

		return fmt.Errorf("token_endpoint failed, Status=%s, Body=%s", resp.Status, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(&c.AccessToken); err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"err":        err,
		}).Error("GetAccessToken:Decode")
		return err
	}

	return nil
}

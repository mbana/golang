package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"fmt"

	"github.com/banaio/openbankingforgerock/requests"
)

const (
	defaultGrantType             = "client_credentials"
	defaultAccessTokenExpiryTime = time.Minute * 30
)

// PostAccessToken gets an `access_token` from the
// `https://as.aspsp.ob.forgerock.financial/oauth2/realms/root/realms/openbanking/access_token`
// endpoint and sets `OpenBankingClient.AccessToken` to the obtained `access_token`.
//
// see: https://openbanking.atlassian.net/wiki/spaces/DZ/pages/187793608/Integrating+a+TPP+with+ForgeRock+Model+Bank+on+Directory+Sandbox#IntegratingaTPPwithForgeRockModelBankonDirectorySandbox-3.1GetanaccesstokentorepresentyouasaTPPusingtheClientcredentialflow.
func (c *OpenBankingClient) PostAccessToken() error {
	body, err := c.makeAccessTokenBody()
	if err != nil {
		return errors.Wrapf(err, "failed PostAccessToken: url=%q", c.OpenIDConfig.TokenEndpoint)
	}

	req, err := requests.NewRequest("POST", c.OpenIDConfig.TokenEndpoint, strings.NewReader(body))
	if err != nil {
		return errors.Wrapf(err, "failed PostAccessToken: url=%q", c.OpenIDConfig.TokenEndpoint)
	}
	req.Header.Set(requests.HeaderContentType, requests.MIMEApplicationForm)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return errors.Wrapf(err, "failed PostAccessToken: url=%q", c.OpenIDConfig.TokenEndpoint)
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrapf(err, "failed PostAccessToken: url=%q, statusCode=%d",
				c.OpenIDConfig.TokenEndpoint,
				resp.StatusCode)
		}
		defer resp.Body.Close()

		return fmt.Errorf("failed PostAccessToken: url=%q, statusCode=%d, body=%s",
			c.OpenIDConfig.TokenEndpoint,
			resp.StatusCode,
			string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(&c.AccessToken); err != nil {
		body, _ := ioutil.ReadAll(resp.Body) // nolint:errcheck
		defer resp.Body.Close()

		return errors.Wrapf(err, "failed PostAccessToken: url=%q, statusCode=%d, body=%s",
			c.OpenIDConfig.TokenEndpoint,
			resp.StatusCode,
			string(body))
	}

	logrus.WithFields(logrus.Fields{
		"expiry_time": time.Now().UTC().Add(time.Second * time.Duration(c.AccessToken.ExpiresIn)),
		"now":         time.Now().UTC(),
	}).Info("PostAccessToken")

	return nil
}

func (c *OpenBankingClient) makeAccessTokenBody() (string, error) {
	data := url.Values{}

	switch c.RegisterResponse.TokenEndpointAuthMethod {
	case TLSClientAuth:
		data.Add("grant_type", defaultGrantType)
		data.Add("scope", "openid payments fundsconfirmations accounts")
		data.Add("client_id", c.RegisterResponse.ClientID)
	case PrivateKeyJWT:
		now := time.Now()
		iat := now.Unix()
		// "The expiration time. After this time, this JWT won't be considered a valid credential. For security reasons, we recommend you set a short period of life, such as 1 or 2 minutes."
		// 30 minutes seems to work, anything larger fails.
		exp := now.Add(defaultAccessTokenExpiryTime).Unix()
		// exp := now.Add(2 * time.Minute).Unix()
		// exp := time.Date(2019, 03, 29, 0, 0, 0, 0, time.UTC).Unix()
		uuid, err := uuid.NewRandom()
		if err != nil {
			return "", nil
		}
		jti := uuid.String()
		claims := jwt.MapClaims{
			"iss": c.RegisterResponse.ClientID,
			"sub": c.RegisterResponse.ClientID,
			"aud": c.OpenIDConfig.Issuer,
			"iat": iat,
			"exp": exp,
			"jti": jti,
		}
		clientAssertion := c.Signer.Sign(claims)

		data.Add("grant_type", defaultGrantType)
		data.Add("scope", "openid payments fundsconfirmations accounts")
		data.Add("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
		data.Add("client_assertion", clientAssertion)
	}

	return data.Encode(), nil
}

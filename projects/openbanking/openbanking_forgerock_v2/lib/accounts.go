package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/banaio/openbankingforgerock/accounts"
	"github.com/banaio/openbankingforgerock/requests"
	"github.com/banaio/openbankingforgerock/signer"
	"github.com/banaio/openbankingforgerock/utils"
)

// POSTAccountRequests ...
// see: https://backstage.forgerock.com/knowledge/openbanking/book/b77473305#a77081077
func (c *OpenBankingClient) POSTAccountRequests() (*accounts.AccountRequestsResponse, error) {
	url := accounts.GetAccountURL("CreateAccountAccessConsent")

	filename := "accounts/post_account_requests.json"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "failed on ReadFile in POSTAccountRequests: filename=%+v", filename)
	}
	req, err := requests.NewRequest("POST", url, bytes.NewReader(file))
	if err != nil {
		return nil, errors.Wrapf(err, "failed on NewRequest in in POSTAccountRequests: url=%+v", url)
	}
	req.Header.Set(requests.HeaderAccept, requests.MIMEApplicationJSON)
	req.Header.Set(requests.HeaderContentType, requests.MIMEApplicationJSON)

	authorization := "Bearer " + c.AccessToken.AccessToken
	lastLoggedTime := time.Now().UTC().Format(time.RFC1123) // "Sun, 10 Sep 2017 19:43:31 UTC"
	interactionID := uuid.New().String()
	req.Header.Set(requests.HeaderAuthorization, authorization)
	req.Header.Set(requests.HeaderXFAPICustomerIPAddress, utils.GetCustomerIP())
	req.Header.Set(requests.HeaderXFAPICustomerLastLoggedTime, lastLoggedTime)
	req.Header.Set(requests.HeaderXFAPIFinancialID, "0015800001041REAAY")
	req.Header.Set(requests.HeaderXFAPIInteractionID, interactionID)
	req.Header.Set(requests.HeaderXIdempotencyKey, "FRESCO.21302.GFX.20")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed on HTTPClient.Do in in POSTAccountRequests: req=%+v", req)
	}

	if res.StatusCode != 201 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "failed on ReadAll in in POSTAccountRequests: res=%+v", res)
		}
		defer res.Body.Close()

		return nil, errors.Wrapf(err, "failed on StatusCode in in POSTAccountRequests: body=%+v", string(body))
	}

	response := &accounts.AccountRequestsResponse{}
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, errors.Wrapf(err, "failed on Decode in in POSTAccountRequests: res=%+v", res)
	}

	return response, nil
}

// GETAccountsRequestsHybridFlow -
func (c *OpenBankingClient) GETAccountsRequestsHybridFlow(accountRequests *accounts.AccountRequestsResponse) error {
	consentID := accountRequests.Data.ConsentID
	if consentID == "" {
		return errors.New(`failed GETAccountsRequestsHybridFlow on preconditions: accountRequests.Data.ConsentID != ""`)
	}

	if c.RegisterResponse.ClientID == "" {
		return errors.New(`failed GETAccountsRequestsHybridFlow on preconditions: c.RegisterResponse.ClientID != ""`)
	}

	now := time.Now()
	iat := now.Unix()
	// These don't work, an error is returned saying unreasonable expiry time
	// exp := now.Add(24 * time.Hour * 365).Unix()
	// exp := now.Add(24 * time.Hour).Unix()
	// exp := now.Add(time.Hour).Unix()
	// "The expiration time. After this time, this JWT won't be considered a valid credential. For security reasons, we recommend you set a short period of life, such as 1 or 2 minutes."
	// 30 minutes seems to work, anything larger fails.
	exp := now.Add(30 * time.Minute).Unix()
	// exp := now.Add(2 * time.Minute).Unix()
	// exp := time.Date(2019, 03, 29, 0, 0, 0, 0, time.UTC).Unix()
	uuid, err := uuid.NewRandom()
	if err != nil {
		return errors.Wrap(err, "failed GETAccountsRequestsHybridFlow on uuid.NewRandom")
	}
	jti := uuid.String()
	// claims := fmt.Sprintf(`{
	// 	"id_token": {
	// 		"acr": {
	// 			"value": "urn:openbanking:psd2:sca",
	// 			"essential": true
	// 		},
	// 		"openbanking_intent_id": {
	// 			"value": "%s",
	// 			"essential": true
	// 		}
	// 	},
	// 	"userinfo": {
	// 		"openbanking_intent_id": {
	// 			"value": "%s",
	// 			"essential": true
	// 		}
	// 	}
	// }`, consentID, consentID)
	claims := signer.NewClaims(consentID)

	// var claimsData map[string]interface{}
	// if err := json.Unmarshal([]byte(claims), &claimsData); err != nil {
	// 	logrus.WithFields(logrus.Fields{
	// 		"err":        err,
	// 		"claimsData": claimsData,
	// 	}).Fatal("GETAccountsRequests:Unmarshal")
	// }

	mapClaims := jwt.MapClaims{
		"aud":           c.OpenIDConfig.Issuer,
		"scope":         "openid payments fundsconfirmations accounts",
		"claims":        claims,
		"redirect_uri":  "http://localhost:8080/openbanking/banaio/forgerock",
		"state":         "state_accounts",
		"nonce":         "5a6b0d7832a9fb4f80f1170a",
		"iss":           c.RegisterResponse.ClientID,
		"iat":           iat,
		"exp":           exp,
		"jti":           jti,
		"client_id":     c.RegisterResponse.ClientID,
		"response_type": "code id_token",
		// "claims":        string(claimsData),
		// "sub":           c.RegisterResponse.ClientID,
	}
	request := c.Signer.Sign(mapClaims)

	req, err := requests.NewRequest("GET", c.OpenIDConfig.AuthorizationEndpoint, nil)
	if err != nil {
		return errors.Wrap(err, "failed GETAccountsRequestsHybridFlow on requests.NewRequest")
	}
	req.Header.Set(requests.HeaderContentType, requests.MIMEApplicationForm)

	query := req.URL.Query()
	query.Add("request", request)
	query.Add("response_type", "code id_token")
	query.Add("client_id", c.RegisterResponse.ClientID)
	query.Add("state", "state_accounts")
	query.Add("nonce", "5a6b0d7832a9fb4f80f1170a")
	query.Add("scope", "openid payments fundsconfirmations accounts")
	query.Add("redirect_uri", "http://localhost:8080/openbanking/banaio/forgerock")
	req.URL.RawQuery = query.Encode()

	url := req.URL.String()
	utils.PrintInBox(utils.Green, utils.Green("GETAccountsRequestsHybridFlow - opening link in browser to authorise:"), utils.Green(url))
	if err := browser.OpenURL(url); err != nil {
		return errors.Wrap(err, "failed GETAccountsRequestsHybridFlow on browser.OpenURL")
	}

	return nil
}

// POSTExchangeCodeForAccessToken -
// https://backstage.forgerock.com/knowledge/openbanking/book/b77473305#exchange
func (c *OpenBankingClient) POSTExchangeCodeForAccessToken(authoriseResponse AuthoriseResponse) (*ExchangeTokenResponse, error) {
	if c.RegisterResponse.ClientID == "" {
		return nil, errors.New("registerResponse.ClientID == '': POSTExchangeCodeForAccessToken")
	}
	if authoriseResponse.Code == "" {
		return nil, errors.New("authoriseResponse.Code == '': POSTExchangeCodeForAccessToken")
	}

	data := url.Values{}
	switch c.RegisterResponse.TokenEndpointAuthMethod {
	case TLSClientAuth:
		data.Add("grant_type", "authorization_code")
		data.Add("code", authoriseResponse.Code)
		data.Add("redirect_uri", "http://localhost:8080/openbanking/banaio/forgerock")
		data.Add("client_id", c.RegisterResponse.ClientID)

	case PrivateKeyJWT:
		now := time.Now()
		iat := now.Unix()
		// "The expiration time. After this time, this JWT won't be considered a valid credential. For security reasons, we recommend you set a short period of life, such as 1 or 2 minutes."
		// 30 minutes seems to work, anything larger fails.
		exp := now.Add(30 * time.Minute).Unix()
		// exp := now.Add(2 * time.Minute).Unix()
		// exp := now.Add(24 * time.Hour).Unix()
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

		clientAssertion := c.Signer.Sign(claims)

		data.Add("grant_type", "authorization_code")
		data.Add("code", authoriseResponse.Code)
		data.Add("redirect_uri", "http://localhost:8080/openbanking/banaio/forgerock")
		data.Add("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
		data.Add("client_assertion", clientAssertion)
	}

	req, err := requests.NewRequest("POST", c.OpenIDConfig.TokenEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":  err,
			"data": data,
		}).Error("POSTExchangeCodeForAccessToken:NewRequest")
		return nil, err
	}
	req.Header.Set(requests.HeaderContentType, requests.MIMEApplicationForm)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": res.StatusCode,
			"err":        err,
			"res":        res,
			"data":       data,
		}).Error("POSTExchangeCodeForAccessToken:Do")
		return nil, err
	}

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"StatusCode": res.StatusCode,
				"body":       string(body),
				"err":        err,
			}).Error("POSTExchangeCodeForAccessToken:ReadAll")
			return nil, err
		}
		defer res.Body.Close()

		logrus.WithFields(logrus.Fields{
			"StatusCode": res.StatusCode,
			"body":       string(body),
		}).Error("POSTExchangeCodeForAccessToken:ReadAll")
		return nil, errors.New(string(body))
	}

	exchangeAccessTokenResponse := &ExchangeTokenResponse{}
	if err := json.NewDecoder(res.Body).Decode(&exchangeAccessTokenResponse); err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": res.StatusCode,
			"err":        err,
		}).Error("POSTExchangeCodeForAccessToken:Decode")
		return nil, err
	}

	return exchangeAccessTokenResponse, nil
}

// GETAccountRequests -
func (c *OpenBankingClient) GETAccountRequests(exchangeAccessTokenResponse *ExchangeTokenResponse) (string, error) {
	// we need the `access_token`
	if exchangeAccessTokenResponse.AccessToken == "" {
		return "", errors.New(`failed in GETAccountRequests: precondition not met exchangeAccessTokenResponse.AccessToken != ""`)
	}

	url := accounts.GetAccountURL("GetAccounts")
	filename := "accounts/get_account_requests.json"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.Wrapf(err, `failed in GETAccountRequests: cannot read filename=%q`, filename)
	}

	req, err := requests.NewRequest("GET", url, bytes.NewReader(file))
	if err != nil {
		return "", errors.Wrap(err, `failed in GETAccountRequests: failed in requests.NewRequest`)
	}
	req.Header.Set(requests.HeaderAccept, requests.MIMEApplicationJSON)
	req.Header.Set(requests.HeaderContentType, requests.MIMEApplicationJSON)

	authorization := "Bearer " + exchangeAccessTokenResponse.AccessToken
	// now := time.Now().Format(time.RFC1123)
	// lastLoggedTime := "Sun, 10 Sep 2017 19:43:31 UTC"
	lastLoggedTime := time.Now().UTC().Format(time.RFC1123)
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, `failed in GETAccountRequests: failed in uuid.NewRandom`)
	}
	interactionID := uuid.String()

	req.Header.Set(requests.HeaderAuthorization, authorization)
	req.Header.Set(requests.HeaderXFAPICustomerIPAddress, utils.GetCustomerIP())
	req.Header.Set(requests.HeaderXFAPICustomerLastLoggedTime, lastLoggedTime)
	req.Header.Set(requests.HeaderXFAPIFinancialID, "0015800001041REAAY")
	req.Header.Set(requests.HeaderXFAPIInteractionID, interactionID)
	req.Header.Set(requests.HeaderXIdempotencyKey, "FRESCO.21302.GFX.20")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, `failed in GETAccountRequests: failed in c.HTTPClient.Do`)
	}

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.Wrap(err, `failed in GETAccountRequests: failed in ioutil.ReadAll`)
		}
		defer res.Body.Close()

		return "", errors.Wrapf(err, `failed in GETAccountRequests: body=%+v, StatusCode=%+v`, string(body), res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, `failed in GETAccountRequests: failed in ioutil.ReadAll`)
	}

	defer res.Body.Close()
	return string(body), nil
}

// GetStatements -
func (c *OpenBankingClient) GetStatements(exchangeAccessTokenResponse ExchangeTokenResponse) (string, error) {
	// we need the `access_token`
	if exchangeAccessTokenResponse.AccessToken == "" {
		err := errors.New(`exchangeAccessTokenResponse.AccessToken == ""`)
		logrus.WithFields(logrus.Fields{
			"err":                         err,
			"exchangeAccessTokenResponse": exchangeAccessTokenResponse,
		}).Error("GetStatements")
		return "", err
	}

	url := strings.Replace(accounts.GetAccountURL("GetAccountStatements"), "{AccountId}", "549908cd-e3a1-4e8b-932d-740d1549437d", 1)
	req, err := requests.NewRequest("GET", url, bytes.NewReader(nil))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
			"req": fmt.Sprintf("%+v", req),
		}).Error("GetStatements:NewRequest")
		return "", err
	}
	req.Header.Set(requests.HeaderAccept, requests.MIMEApplicationJSON)
	req.Header.Set(requests.HeaderContentType, requests.MIMEApplicationJSON)

	authorization := "Bearer " + exchangeAccessTokenResponse.AccessToken
	// now := time.Now().Format(time.RFC1123)
	// lastLoggedTime := "Sun, 10 Sep 2017 19:43:31 UTC"
	lastLoggedTime := time.Now().UTC().Format(time.RFC1123)
	interactionID := uuid.New().String()

	req.Header.Set(requests.HeaderAuthorization, authorization)
	req.Header.Set(requests.HeaderXFAPICustomerIPAddress, utils.GetCustomerIP())
	req.Header.Set(requests.HeaderXFAPICustomerLastLoggedTime, lastLoggedTime)
	req.Header.Set(requests.HeaderXFAPIFinancialID, "0015800001041REAAY")
	req.Header.Set(requests.HeaderXFAPIInteractionID, interactionID)
	req.Header.Set(requests.HeaderXIdempotencyKey, "FRESCO.21302.GFX.20")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": res.StatusCode,
			"Header":     req.Header,
			"err":        err,
			"res":        res,
		}).Error("GetStatements:Do")
		return "", err
	}

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"StatusCode": res.StatusCode,
				"err":        err,
			}).Error("GetStatements:ReadAll")
			return "", err
		}
		defer res.Body.Close()

		logrus.WithFields(logrus.Fields{
			"StatusCode": res.StatusCode,
			"body":       string(body),
		}).Error("GetStatements")
		return "", errors.New(string(body))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": res.StatusCode,
			"err":        err,
		}).Error("GetStatements:ReadAll")
		return "", err
	}

	defer res.Body.Close()
	return string(body), nil
}

// GetStatements -
func (c *OpenBankingClient) GetBalances(exchangeAccessTokenResponse ExchangeTokenResponse) (string, error) {
	// we need the `access_token`
	if exchangeAccessTokenResponse.AccessToken == "" {
		return "", errors.New(`failed in GetBalances: precondition not met exchangeAccessTokenResponse.AccessToken != ""`)
	}

	url := accounts.GetAccountURL("GetBalances")
	req, err := requests.NewRequest("GET", url, bytes.NewReader(nil))
	if err != nil {
		return "", errors.Wrap(err, `failed in GetBalances: failed in requests.NewRequest`)
	}
	req.Header.Set(requests.HeaderAccept, requests.MIMEApplicationJSON)
	req.Header.Set(requests.HeaderContentType, requests.MIMEApplicationJSON)

	authorization := "Bearer " + exchangeAccessTokenResponse.AccessToken
	// now := time.Now().Format(time.RFC1123)
	// lastLoggedTime := "Sun, 10 Sep 2017 19:43:31 UTC"
	lastLoggedTime := time.Now().UTC().Format(time.RFC1123)
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, `failed in GetBalances: failed in uuid.NewRandom`)
	}
	interactionID := uuid.String()

	req.Header.Set(requests.HeaderAuthorization, authorization)
	req.Header.Set(requests.HeaderXFAPICustomerIPAddress, utils.GetCustomerIP())
	req.Header.Set(requests.HeaderXFAPICustomerLastLoggedTime, lastLoggedTime)
	req.Header.Set(requests.HeaderXFAPIFinancialID, "0015800001041REAAY")
	req.Header.Set(requests.HeaderXFAPIInteractionID, interactionID)
	req.Header.Set(requests.HeaderXIdempotencyKey, "FRESCO.21302.GFX.20")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, `failed in GetBalances: failed in c.HTTPClient.Do`)
	}

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.Wrap(err, `failed in GetBalances: failed in ioutil.ReadAll`)
		}
		defer res.Body.Close()

		return "", errors.Wrapf(err, `failed in GetBalances: body=%+v, StatusCode=%+v`, string(body), res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, `failed in GetBalances: failed in ioutil.ReadAll`)
	}

	defer res.Body.Close()
	return string(body), nil
}

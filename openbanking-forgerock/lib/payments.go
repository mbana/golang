package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/browser"
	"github.com/pkg/errors"

	"net/http"

	"github.com/banaio/openbankingforgerock/payments"
	"github.com/banaio/openbankingforgerock/requests"
	"github.com/banaio/openbankingforgerock/signer"
	"github.com/banaio/openbankingforgerock/utils"
)

func (c *OpenBankingClient) CreateDomesticPaymentConsent() (*payments.CreateDomesticPaymentConsentResponse, error) {
	url := payments.GetPaymentsURL("CreateDomesticPaymentConsent")

	filename := "payments/testdata/CreateDomesticPaymentConsent.json"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed CreateDomesticPaymentConsent on ioutil.ReadFile")
	}
	req, err := requests.NewRequest("POST", url, bytes.NewReader(file))
	if err != nil {
		return nil, errors.Wrap(err, "failed CreateDomesticPaymentConsent on NewRequest")
	}
	req.Header.Set(requests.HeaderAccept, requests.MIMEApplicationJSON)
	req.Header.Set(requests.HeaderContentType, requests.MIMEApplicationJSON)

	jwsSignature, err := signer.MakeJWSSignature(c.Config, c.Keys, file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed CreateDomesticPaymentConsent in signer.MakeJWSSignature: url=%q", url)
	}

	authorization := "Bearer " + c.AccessToken.AccessToken
	lastLoggedTime := time.Now().UTC().Format(time.RFC1123) // "Sun, 10 Sep 2017 19:43:31 UTC"
	interactionID := uuid.New().String()
	req.Header.Set(requests.HeaderAuthorization, authorization)
	req.Header.Set(requests.HeaderXFAPICustomerIPAddress, utils.GetCustomerIP())
	req.Header.Set(requests.HeaderXFAPICustomerLastLoggedTime, lastLoggedTime)
	req.Header.Set(requests.HeaderXFAPIFinancialID, "0015800001041REAAY")
	req.Header.Set(requests.HeaderXFAPIInteractionID, interactionID)
	req.Header.Set(requests.HeaderXIdempotencyKey, "FRESCO.21302.GFX.20")
	req.Header.Set(requests.HeaderXJWSSignatureKey, jwsSignature)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed CreateDomesticPaymentConsent")
	}

	if res.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "failed CreateDomesticPaymentConsent in ioutil.ReadAll: res=%+v", res)
		}
		defer res.Body.Close()

		return nil, errors.Errorf("failed CreateDomesticPaymentConsent: body=%+v, StatuCode=%+v", string(body), res.StatusCode)
	}

	response := &payments.CreateDomesticPaymentConsentResponse{}
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "failed CreateDomesticPaymentConsent in ioutil.ReadAll: res=%+v", res)
		}
		defer res.Body.Close()

		return nil, errors.Wrapf(err, "failed CreateDomesticPaymentConsent: body=%+v", string(body))
	}

	return response, nil
}

func (c *OpenBankingClient) GetDomesticPaymentConsent(consent *payments.CreateDomesticPaymentConsentResponse) (string, error) {
	url := strings.Replace(
		payments.GetPaymentsURL("GetDomesticPaymentConsent"),
		"{ConsentId}",
		consent.Data.ConsentID,
		1,
	)

	req, err := requests.NewRequest("GET", url, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed GetDomesticPaymentConsent on NewRequest")
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
		return "", errors.Wrap(err, "failed GetDomesticPaymentConsent")
	}

	if res.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.Wrapf(err, "failed GetDomesticPaymentConsent in ioutil.ReadAll: res=%+v, StatuCode=%+v", res, res.StatusCode)
		}
		defer res.Body.Close()

		return "", errors.Errorf("failed GetDomesticPaymentConsent: body=%+v, StatuCode=%+v", string(body), res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrapf(err, "failed GetDomesticPaymentConsent in ioutil.ReadAll: res=%+v, StatuCode=%+v", res, res.StatusCode)
	}
	defer res.Body.Close()

	return string(body), nil
}

func (c *OpenBankingClient) CreateDomesticPaymentHybridFlow(consent *payments.CreateDomesticPaymentConsentResponse) error {
	consentID := consent.Data.ConsentID
	if consentID == "" {
		return errors.New(`failed CreateDomesticPaymentHybridFlow on preconditions: accountRequests.Data.ConsentID != ""`)
	}

	if c.RegisterResponse.ClientID == "" {
		return errors.New(`failed CreateDomesticPaymentHybridFlow on preconditions: c.RegisterResponse.ClientID != ""`)
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
		return errors.Wrap(err, "failed CreateDomesticPaymentHybridFlow on uuid.NewRandom")
	}
	jti := uuid.String()
	claims := signer.NewClaims(consentID)

	mapClaims := jwt.MapClaims{
		"aud":           c.OpenIDConfig.Issuer,
		"scope":         "openid payments fundsconfirmations accounts",
		"claims":        claims,
		"redirect_uri":  "http://localhost:8080/openbanking/banaio/forgerock",
		"state":         "state_payments",
		"nonce":         "5a6b0d7832a9fb4f80f1170a",
		"iss":           c.RegisterResponse.ClientID,
		"iat":           iat,
		"exp":           exp,
		"jti":           jti,
		"client_id":     c.RegisterResponse.ClientID,
		"response_type": "code id_token",
		// "sub":           c.RegisterResponse.ClientID,
	}
	request := c.Signer.Sign(mapClaims)

	req, err := requests.NewRequest("GET", c.OpenIDConfig.AuthorizationEndpoint, nil)
	if err != nil {
		return errors.Wrap(err, "failed CreateDomesticPaymentHybridFlow on requests.NewRequest")
	}
	req.Header.Set(requests.HeaderContentType, requests.MIMEApplicationForm)

	query := req.URL.Query()
	query.Add("request", request)
	query.Add("response_type", "code id_token")
	query.Add("client_id", c.RegisterResponse.ClientID)
	query.Add("state", "state_payments")
	query.Add("nonce", "5a6b0d7832a9fb4f80f1170a")
	query.Add("scope", "openid payments fundsconfirmations accounts")
	query.Add("redirect_uri", "http://localhost:8080/openbanking/banaio/forgerock")
	req.URL.RawQuery = query.Encode()

	url := req.URL.String()
	utils.PrintInBox(utils.Green, utils.Green("CreateDomesticPaymentHybridFlow - opening link in browser to authorise:"), utils.Green(url))
	if err := browser.OpenURL(url); err != nil {
		return errors.Wrap(err, "failed CreateDomesticPaymentHybridFlow on browser.OpenURL")
	}

	return nil
}

func (c *OpenBankingClient) CreateDomesticPayment(exchangeAccessTokenResponse *ExchangeTokenResponse, consent *payments.CreateDomesticPaymentConsentResponse) (string, error) {
	url := payments.GetPaymentsURL("CreateDomesticPayment")

	filename := "payments/testdata/CreateDomesticPayment.json"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.Wrap(err, "failed CreateDomesticPayment on ioutil.ReadFile")
	}
	reqBody := strings.Replace(
		string(file),
		"{ConsentId}",
		consent.Data.ConsentID,
		1,
	)
	req, err := requests.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		return "", errors.Wrap(err, "failed CreateDomesticPayment on NewRequest")
	}
	req.Header.Set(requests.HeaderAccept, requests.MIMEApplicationJSON)
	req.Header.Set(requests.HeaderContentType, requests.MIMEApplicationJSON)

	jwsSignature, err := signer.MakeJWSSignature(c.Config, c.Keys, []byte(reqBody))
	if err != nil {
		return "", errors.Wrapf(err, "failed CreateDomesticPayment in signer.MakeJWSSignature: url=%q", url)
	}

	authorization := "Bearer " + exchangeAccessTokenResponse.AccessToken
	lastLoggedTime := time.Now().UTC().Format(time.RFC1123) // "Sun, 10 Sep 2017 19:43:31 UTC"
	interactionID := uuid.New().String()
	req.Header.Set(requests.HeaderAuthorization, authorization)
	req.Header.Set(requests.HeaderXFAPICustomerIPAddress, utils.GetCustomerIP())
	req.Header.Set(requests.HeaderXFAPICustomerLastLoggedTime, lastLoggedTime)
	req.Header.Set(requests.HeaderXFAPIFinancialID, "0015800001041REAAY")
	req.Header.Set(requests.HeaderXFAPIInteractionID, interactionID)
	req.Header.Set(requests.HeaderXIdempotencyKey, "FRESCO.21302.GFX.20")
	req.Header.Set(requests.HeaderXJWSSignatureKey, jwsSignature)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "failed CreateDomesticPayment")
	}

	if res.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", errors.Wrapf(err, "failed CreateDomesticPayment in ioutil.ReadAll: res=%+v, StatuCode=%+v", res, res.StatusCode)
		}
		defer res.Body.Close()

		return "", errors.Errorf("failed CreateDomesticPayment: body=%+v, StatuCode=%+v", string(body), res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrapf(err, "failed CreateDomesticPayment in ioutil.ReadAll: res=%+v, StatuCode=%+v", res, res.StatusCode)
	}
	defer res.Body.Close()

	return string(body), nil
}

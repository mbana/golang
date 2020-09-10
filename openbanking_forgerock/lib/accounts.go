package lib

import (
	"bytes"
	"encoding/json"
	"errors"
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

// AuthoriseResponse ...
type AuthoriseResponse struct {
	Code    string `json:"code" form:"code" query:"code"`
	IDToken string `json:"id_token" form:"id_token" query:"id_token"`
	Scope   string `json:"scope" form:"scope" query:"scope"`
	State   string `json:"state" form:"state" query:"state"`
}

// ExchangeTokenResponse ...
type ExchangeTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Nonce        string `json:"nonce"`
}

// AccountRequestsResponse ...
// {
// 	"Data": {
// 		"AccountRequestId": "A02aff57e-80f9-4964-8548-4c9b17cfaa29",
// 		"Status": "AwaitingAuthorisation",
// 		"CreationDateTime": "2018-10-19T08:36:48+00:00",
// 		"Permissions": [
// 			"ReadAccountsDetail",
// 			"ReadBalances",
// 			"ReadBeneficiariesDetail",
// 			"ReadDirectDebits",
// 			"ReadProducts",
// 			"ReadStandingOrdersDetail",
// 			"ReadTransactionsCredits",
// 			"ReadTransactionsDebits",
// 			"ReadTransactionsDetail"
// 		],
// 		"ExpirationDateTime": "2018-05-02T00:00:00+00:00",
// 		"TransactionFromDateTime": "2017-05-03T00:00:00+00:00",
// 		"TransactionToDateTime": "2018-12-03T00:00:00+00:00"
// 	},
// 	"Risk": {}
// }
type AccountRequestsResponse struct {
	Data AccountRequestsResponseData `json:"Data"`
	Risk map[string]string           `json:"Risk"`
}

// AccountRequestsResponseData ...
// 	"Data": {
// 		"AccountRequestId": "A02aff57e-80f9-4964-8548-4c9b17cfaa29",
// 		"Status": "AwaitingAuthorisation",
// 		"CreationDateTime": "2018-10-19T08:36:48+00:00",
// 		"Permissions": [
// 			"ReadAccountsDetail",
// 			"ReadBalances",
// 			"ReadBeneficiariesDetail",
// 			"ReadDirectDebits",
// 			"ReadProducts",
// 			"ReadStandingOrdersDetail",
// 			"ReadTransactionsCredits",
// 			"ReadTransactionsDebits",
// 			"ReadTransactionsDetail"
// 		],
// 		"ExpirationDateTime": "2018-05-02T00:00:00+00:00",
// 		"TransactionFromDateTime": "2017-05-03T00:00:00+00:00",
// 		"TransactionToDateTime": "2018-12-03T00:00:00+00:00"
// 	},
type AccountRequestsResponseData struct {
	AccountRequestID        string   `json:"AccountRequestId"`
	Status                  string   `json:"Status"`
	CreationDateTime        string   `json:"CreationDateTime"`
	Permissions             []string `json:"Permissions"`
	ExpirationDateTime      string   `json:"ExpirationDateTime"`
	TransactionFromDateTime string   `json:"TransactionFromDateTime"`
	TransactionToDateTime   string   `json:"TransactionToDateTime"`
}

// GETAccountRequests the actual accounts request. This should be called once we have gone through
// the hybrid flow and exchanged the `code` token for the `access_token`.
// see: https://backstage.forgerock.com/knowledge/openbanking/book/b77473305#a77081077
//
// It returns a structure similar to the below:
// {
//     "Data": {
//         "Account": [
//             {
//                 "AccountId": "3b0576a9-038d-40ff-9fff-ca74871f9c2b",
//                 "Currency": "GBP",
//                 "Nickname": "Bills",
//                 "Account": {
//                     "SchemeName": "SortCodeAccountNumber",
//                     "Identification": "93163240337365",
//                     "Name": "mbana",
//                     "SecondaryIdentification": "69789331"
//                 }
//             },
//             {
//                 "AccountId": "e447a126-c7ed-4ac6-a88f-06f2a8ed4e3b",
//                 "Currency": "GBP",
//                 "Nickname": "Household",
//                 "Account": {
//                     "SchemeName": "SortCodeAccountNumber",
//                     "Identification": "93345435281245",
//                     "Name": "mbana"
//                 }
//             }
//         ]
//     },
//     "Links": {
//         "Self": "https://rs.aspsp.ob.forgerock.financial/open-banking/v1.1/accounts"
//     },
//     "Meta": {
//         "TotalPages": 1
//     }
// }
func (c *OpenBankingClient) GETAccountRequests(exchangeAccessTokenResponse ExchangeTokenResponse) (string, error) {
	// Example call from command line:
	// $ curl -X GET \
	//   https://rs.aspsp.dev-ob.forgerock.financial:443/open-banking/v1.1/accounts \
	//   -H 'Accept: application/json' \
	//   -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJ6aXAiOiJOT05FIiwia2lkIjoiRm9sN0lwZEtlTFptekt0Q0VnaTFMRGhTSXpNPSIsImFsZyI6IkVTMjU2In0.eyJzdWIiOiJkZW1vIiwiYXV0aF9sZXZlbCI6MCwiYXVkaXRUcmFja2luZ0lkIjoiOTllOTk0MDctZGViZi00MmJiLTgxNmMtZWMzYjY3OWQzNjA1IiwiaXNzIjoiaHR0cHM6Ly9hcy5hc3BzcC5kZXYtb2IuZm9yZ2Vyb2NrLmZpbmFuY2lhbDo4NDQzL29hdXRoMi9vcGVuYmFua2luZyIsInRva2VuTmFtZSI6ImFjY2Vzc190b2tlbiIsInRva2VuX3R5cGUiOiJCZWFyZXIiLCJhdXRoR3JhbnRJZCI6ImZkMWE1NmRjLWFjMWQtNDExYS1hYTUzLTU3MmNhNzEyM2M2ZCIsIm5vbmNlIjoiNWE2YjFhYzkzMmE5ZmI1MTY0MThkZDQ5IiwiYXVkIjoiODU5OTRjYTAtOWRiMy00NTA1LTkzNWItMDdkYzA3NWNiNzA3IiwibmJmIjoxNTE2OTY4NzA0LCJncmFudF90eXBlIjoiYXV0aG9yaXphdGlvbl9jb2RlIiwic2NvcGUiOlsib3BlbmlkIiwiYWNjb3VudHMiXSwiYXV0aF90aW1lIjoxNTE2OTY4NjcwMDAwLCJjbGFpbXMiOiJ7XCJpZF90b2tlblwiOntcImFjclwiOntcInZhbHVlXCI6XCJ1cm46b3BlbmJhbmtpbmc6cHNkMjpzY2FcIixcImVzc2VudGlhbFwiOnRydWV9LFwib3BlbmJhbmtpbmdfaW50ZW50X2lkXCI6e1widmFsdWVcIjpcIkEwOWIxYzE5Yi0wMDRlLTQzMjctODczMC1kNjU3MzkwYTBlYzlcIixcImVzc2VudGlhbFwiOnRydWV9fSxcInVzZXJpbmZvXCI6e1wib3BlbmJhbmtpbmdfaW50ZW50X2lkXCI6e1widmFsdWVcIjpcIkEwOWIxYzE5Yi0wMDRlLTQzMjctODczMC1kNjU3MzkwYTBlYzlcIixcImVzc2VudGlhbFwiOnRydWV9fX0iLCJyZWFsbSI6Ii9vcGVuYmFua2luZyIsImV4cCI6MTUxNjk3MjMwNCwiaWF0IjoxNTE2OTY4NzA0LCJleHBpcmVzX2luIjozNjAwLCJqdGkiOiIyNmZlNWFjOS01ZjNkLTQzNGUtOWY0OC0wOGVhMGQyZjBmODcifQ.y5Xc9LeUcit3C35lMSK9jX4XqYY7xtCd_V_Vvww6YW6ESbUbW2fEOpvVc1pFjrw1cumSrD3HgwUoOHDLUn_06w' \
	//   -H 'Cache-Control: no-cache' \
	//   -H 'Content-Type: application/json' \
	//   -H 'Postman-Token: 18ea0f4e-d1c5-95bc-808c-e5e95c7c2c99' \
	//   -H 'x-fapi-customer-ip-address: 104.25.212.99' \
	//   -H 'x-fapi-customer-last-logged-time: Sun, 10 Sep 2017 19:43:31 UTC' \
	//   -H 'x-fapi-financial-id: 0015800001041REAAY' \
	//   -H 'x-fapi-interaction-id: 93bac548-d2de-4546-b106-880a5018460d' \
	//   -H 'x-idempotency-key: FRESCO.21302.GFX.20' \
	//   -d '{
	//   "Data": {
	//     "Permissions": [
	//       "ReadAccountsDetail",
	//       "ReadBalances",
	//       "ReadBeneficiariesDetail",
	//       "ReadDirectDebits",
	//       "ReadProducts",
	//       "ReadStandingOrdersDetail",
	//       "ReadTransactionsCredits",
	//       "ReadTransactionsDebits",
	//       "ReadTransactionsDetail"
	//     ],
	//     "ExpirationDateTime": "2017-05-02T00:00:00+00:00",
	//     "TransactionFromDateTime": "2017-05-03T00:00:00+00:00",
	//     "TransactionToDateTime": "2017-12-03T00:00:00+00:00"
	//   },
	//   "Risk": {}
	// }'

	// we need the `access_token`
	if exchangeAccessTokenResponse.AccessToken == "" {
		err := errors.New(`exchangeAccessTokenResponse.AccessToken == ""`)
		logrus.WithFields(logrus.Fields{
			"err":                         err,
			"exchangeAccessTokenResponse": fmt.Sprintf("%+v", exchangeAccessTokenResponse),
		}).Error("GETAccountRequests")
		return "", err
	}

	url := GetAccountURL("GetAccounts")
	// call proxy to do validation
	// url := "http://localhost:8989/open-banking/v2.0/accounts"

	file, err := ioutil.ReadFile("lib/accounts/get_account_requests.json")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("GETAccountRequests:ReadFile")
		return "", err
	}

	req, err := http.NewRequest("GET", url, bytes.NewReader(file))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
			"req": fmt.Sprintf("%+v", req),
		}).Error("GETAccountRequests:NewRequest")
		return "", err
	}

	authorization := "Bearer " + exchangeAccessTokenResponse.AccessToken
	// ip := GetOutboundIP()
	// customerIP := ip.String()
	customerIP := "104.25.212.99"
	// now := time.Now().Format(time.RFC1123)
	// lastLoggedTime := "Sun, 10 Sep 2017 19:43:31 UTC"
	lastLoggedTime := time.Now().Format(time.RFC1123)
	interactionID := uuid.New().String()
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")

	req.Header.Set("Authorization", authorization)
	req.Header.Set("x-fapi-customer-ip-address", customerIP)
	// req.Header.Set("x-fapi-customer-last-logged-time", "Sun, 10 Sep 2017 19:43:31 UTC")
	req.Header.Set("x-fapi-customer-last-logged-time", lastLoggedTime)
	req.Header.Set("x-fapi-financial-id", "0015800001041REAAY")
	// req.Header.Set("x-fapi-financial-id", "5b507065b093465496d238a8")
	req.Header.Set("x-fapi-interaction-id", interactionID)
	req.Header.Set("x-idempotency-key", "FRESCO.21302.GFX.20")

	RequestToCurlCommand(req, "GETAccountRequests")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": res.StatusCode,
			"Header":     req.Header,
			"err":        err,
			"res":        res,
		}).Error("GETAccountRequests:Do")
		return "", err
	}

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"StatusCode": res.StatusCode,
				"err":        err,
			}).Error("GETAccountRequests:ReadAll")
			return "", err
		}

		logrus.WithFields(logrus.Fields{
			"StatusCode": res.StatusCode,
			"body":       string(body),
		}).Error("GETAccountRequests")
		return "", errors.New(string(body))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": res.StatusCode,
			"err":        err,
		}).Error("GETAccountRequests:ReadAll")
		return "", err
	}

	logrus.WithFields(logrus.Fields{
		"StatusCode": res.StatusCode,
		"body":       string(body),
	}).Info("GETAccountRequests")

	return string(body), nil
}

// POSTAccountRequests ...
// see: https://backstage.forgerock.com/knowledge/openbanking/book/b77473305#a77081077
func (c *OpenBankingClient) POSTAccountRequests() (*AccountRequestsResponse, error) {
	// $ curl -X POST \
	//   https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/account-requests \
	//   -H 'Accept: application/json' \
	//   -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJ6aXAiOiJOT05FIiwia2lkIjoiRm9sN0lwZEtlTFptekt0Q0VnaTFMRGhTSXpNPSIsImFsZyI6IkVTMjU2In0.eyJzdWIiOiI1YjhjY2E4NS0xODcwLTQ1MzYtYTAxNi00ZjI2YmJlYWU3NDIiLCJhdWRpdFRyYWNraW5nSWQiOiI5OTg3OTJjOS03NmNkLTRkZGUtYTQzMC1mOGVjZDZlOTdiZDEiLCJpc3MiOiJodHRwczovL2FzLmFzcHNwLmludGVnLW9iLmZvcmdlcm9jay5maW5hbmNpYWwvb2F1dGgyL29wZW5iYW5raW5nIiwidG9rZW5OYW1lIjoiYWNjZXNzX3Rva2VuIiwidG9rZW5fdHlwZSI6IkJlYXJlciIsImF1dGhHcmFudElkIjoiYzE0NjdkMzYtY2JhOS00NjY0LWFmOTgtMjEyZjU3NGJjYjQ0IiwiYXVkIjoiNWI4Y2NhODUtMTg3MC00NTM2LWEwMTYtNGYyNmJiZWFlNzQyIiwibmJmIjoxNTEzNzc1OTY1LCJncmFudF90eXBlIjoiY2xpZW50X2NyZWRlbnRpYWxzIiwic2NvcGUiOlsib3BlbmlkIiwiYWNjb3VudHMiXSwiYXV0aF90aW1lIjoxNTEzNzc1OTY1LCJyZWFsbSI6Ii9vcGVuYmFua2luZyIsImV4cCI6MTUxMzc3OTU2NSwiaWF0IjoxNTEzNzc1OTY1LCJleHBpcmVzX2luIjozNjAwLCJqdGkiOiIzYjhlNzc0OS00ZmM1LTQzMDctYTY0NC0zYmRjMzliZDE4NzEifQ.6qjz6oy9Qer9lFftPkummWaxrO1afPEypp8SxUPKYN2HVsC3vGV68WkDELYuBg01GOT73Ej3OAunlW5dbPPrlA' \
	//   -H 'Content-Type: application/json' \
	//   -H 'x-fapi-customer-ip-address: 104.25.212.99' \
	//   -H 'x-fapi-customer-last-logged-time: Sun, 10 Sep 2017 19:43:31 UTC' \
	//   -H 'x-fapi-financial-id: 0015800001041REAAY' \
	//   -H 'x-fapi-interaction-id: 93bac548-d2de-4546-b106-880a5018460d' \
	//   -H 'x-idempotency-key: FRESCO.21302.GFX.20' \
	//   -d '{
	//   "Data": {
	//     "Permissions": [
	//       "ReadAccountsDetail",
	//       "ReadBalances",
	//       "ReadBeneficiariesDetail",
	//       "ReadDirectDebits",
	//       "ReadProducts",
	//       "ReadStandingOrdersDetail",
	//       "ReadTransactionsCredits",
	//       "ReadTransactionsDebits",
	//       "ReadTransactionsDetail"
	//     ],
	//     "ExpirationDateTime": "2018-05-02T00:00:00+00:00",
	//     "TransactionFromDateTime": "2017-05-03T00:00:00+00:00",
	//     "TransactionToDateTime": "2018-12-03T00:00:00+00:00"
	//   },
	//   "Risk": {}
	// }'
	// url := "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/account-requests"
	url := GetAccountURL("CreateAccountRequest")

	// req, err := http.NewRequest("GET", url, strings.NewReader(data.Encode()))
	file, err := ioutil.ReadFile("lib/accounts/post_account_requests.json")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("POSTAccountRequests:ReadFile")
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(file))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
			"req": fmt.Sprintf("%+v", req),
		}).Fatal("POSTAccountRequests:NewRequest")
	}

	authorization := c.AccessToken.AccessToken
	// ip := GetOutboundIP()
	// customerIP := ip.String()
	customerIP := "104.25.212.99"
	// now := time.Now().Format(time.RFC1123)
	// lastLoggedTime := "Sun, 10 Sep 2017 19:43:31 UTC"
	lastLoggedTime := time.Now().Format(time.RFC1123)
	interactionID := uuid.New().String()
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")

	req.Header.Set("Authorization", "Bearer "+authorization)
	req.Header.Set("x-fapi-customer-ip-address", customerIP)
	// req.Header.Set("x-fapi-customer-last-logged-time", "Sun, 10 Sep 2017 19:43:31 UTC")
	req.Header.Set("x-fapi-customer-last-logged-time", lastLoggedTime)
	req.Header.Set("x-fapi-financial-id", "0015800001041REAAY")
	// req.Header.Set("x-fapi-financial-id", "5b507065b093465496d238a8")
	req.Header.Set("x-fapi-interaction-id", interactionID)
	req.Header.Set("x-idempotency-key", "FRESCO.21302.GFX.20")

	// RequestToCurlCommand(req, "GetAccounts")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"Header":     req.Header,
			"err":        err,
			"resp":       resp,
		}).Fatal("POSTAccountRequests:Do")
	}

	if resp.StatusCode != 201 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"StatusCode": resp.StatusCode,
				"err":        err,
			}).Fatal("POSTAccountRequests:ReadAll")
		}

		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"body":       string(body),
		}).Fatal("POSTAccountRequests")
	}

	response := &AccountRequestsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"err":        err,
		}).Error("POSTAccountRequests:Decode")
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"StatusCode":              resp.StatusCode,
		"AccountRequestsResponse": fmt.Sprintf("%+v", response),
	}).Info("POSTAccountRequests")

	return response, nil
}

// GETAccountsRequestsHybridFlow ...
func (c *OpenBankingClient) GETAccountsRequestsHybridFlow(accountRequests AccountRequestsResponse) {
	accountRequestID := accountRequests.Data.AccountRequestID
	if accountRequestID == "" {
		logrus.WithFields(logrus.Fields{
			"err":             errors.New("accountRequests.Data.AccountRequestID == ''"),
			"accountRequests": fmt.Sprintf("%+v", accountRequests),
		}).Fatal("GETAccountsRequests:NewRequest")
	}

	if c.RegisterResponse.ClientID == "" {
		logrus.WithFields(logrus.Fields{
			"err":                errors.New("c.RegisterResponse.ClientID == ''"),
			"c.RegisterResponse": fmt.Sprintf("%+v", c.RegisterResponse),
		}).Fatal("GETAccountsRequests:NewRequest")
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
	jti := uuid.New().String()
	claims := fmt.Sprintf(`{
    "id_token": {
      "acr": {
        "value": "urn:openbanking:psd2:sca",
        "essential": true
      },
      "openbanking_intent_id": {
        "value": "%s",
        "essential": true
      }
    },
    "userinfo": {
      "openbanking_intent_id": {
        "value": "%s",
        "essential": true
      }
    }
  }`, accountRequestID, accountRequestID)

	var claimsData map[string]interface{}
	if err := json.Unmarshal([]byte(claims), &claimsData); err != nil {
		logrus.WithFields(logrus.Fields{
			"err":        err,
			"claimsData": fmt.Sprintf("%+v", claimsData),
		}).Fatal("GETAccountsRequests:Unmarshal")
	}

	mapClaims := jwt.MapClaims{
		"aud":           c.OpenIDConfig.Issuer,
		"scope":         "accounts openid",
		"claims":        claimsData,
		"iss":           c.RegisterResponse.ClientID,
		"redirect_uri":  "http://localhost:8080/openbanking/banaio/forgerock",
		"state":         "5a6b0d7832a9fb4f80f1170a",
		"exp":           exp,
		"nonce":         "5a6b0d7832a9fb4f80f1170a",
		"iat":           iat,
		"client_id":     c.RegisterResponse.ClientID,
		"jti":           jti,
		"response_type": "code id_token",
		// "response_type": []string{"code", "id_token", "code id_token"},
		// "redirect_uri":  "https://bana.io/openbanking/forgerock",
	}
	request := Sign(mapClaims, c.Envs.KID)

	req, err := http.NewRequest("GET", c.OpenIDConfig.AuthorizationEndpoint, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
			"req": fmt.Sprintf("%+v", req),
			// "data": data,
		}).Fatal("GETAccountsRequests:NewRequest")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cache-Control", "no-cache")
	// req.Header.Set("Accept", "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")

	query := req.URL.Query()
	query.Add("request", request)
	query.Add("response_type", "code id_token")
	query.Add("client_id", c.RegisterResponse.ClientID)
	query.Add("state", "5a6b0d7832a9fb4f80f1170a")
	query.Add("nonce", "5a6b0d7832a9fb4f80f1170a")
	query.Add("scope", "accounts openid")
	// query.Add("redirect_uri", "https://bana.io/openbanking/forgerock")
	query.Add("redirect_uri", "http://localhost:8080/openbanking/banaio/forgerock")
	req.URL.RawQuery = query.Encode()

	fmt.Println("\033[92m ----------------------------------------- \033[0m")
	fmt.Println("\033[92m open this link in a browser to authorise: \033[0m")
	fmt.Println(req.URL.String())
	fmt.Println("\033[92m ----------------------------------------- \033[0m")
}

// POSTExchangeCodeForAccessToken ...
// https://backstage.forgerock.com/knowledge/openbanking/book/b77473305#exchange
func (c *OpenBankingClient) POSTExchangeCodeForAccessToken(authoriseResponse AuthoriseResponse) (*ExchangeTokenResponse, error) {
	if c.RegisterResponse.ClientID == "" {
		logrus.WithFields(logrus.Fields{
			"err":                errors.New("c.RegisterResponse.ClientID == ''"),
			"c.RegisterResponse": fmt.Sprintf("%+v", c.RegisterResponse),
		}).Error("POSTExchangeCodeForAccessToken:ClientID")
		return nil, errors.New("registerResponse.ClientID == ''")
	}
	if authoriseResponse.Code == "" {
		logrus.WithFields(logrus.Fields{
			"err":               errors.New("authoriseResponse.Code == ''"),
			"authoriseResponse": fmt.Sprintf("%+v", authoriseResponse),
		}).Error("POSTExchangeCodeForAccessToken:Code")
		return nil, errors.New("authoriseResponse.Code == ''")
	}

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

	clientAssertion := Sign(claims, c.Envs.KID)
	data := url.Values{}
	data.Add("grant_type", "authorization_code")
	data.Add("code", authoriseResponse.Code)
	data.Add("redirect_uri", "http://localhost:8080/openbanking/banaio/forgerock")
	data.Add("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	data.Add("client_assertion", clientAssertion)

	// grant_type=authorization_code&code=e03b6939-fc0e-469e-9f83-51d7b6bda140&redirect_uri=https%3A%2F%2Fgoogle.fr&client_assertion_type=urn%3Aietf%3Aparams%3Aoauth%3Aclient-assertion-type%3Ajwt-bearer&client_assertion=YOUR_CLIENT_AUTHENTICATION_JWT

	req, err := http.NewRequest("POST", c.OpenIDConfig.TokenEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":  err,
			"req":  fmt.Sprintf("%+v", req),
			"data": data,
		}).Error("POSTExchangeCodeForAccessToken:NewRequest")
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cache-Control", "no-cache")

	RequestToCurlCommand(req, "POSTExchangeCodeForAccessToken")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode":       resp.StatusCode,
			"err":              err,
			"resp":             resp,
			"data":             data,
			"client_assertion": clientAssertion,
		}).Error("POSTExchangeCodeForAccessToken:Do")
		return nil, err
	}

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"StatusCode": resp.StatusCode,
				"body":       string(body),
				"err":        err,
			}).Error("POSTExchangeCodeForAccessToken:ReadAll")
			return nil, err
		}

		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"body":       string(body),
		}).Error("POSTExchangeCodeForAccessToken:ReadAll")
		return nil, errors.New(string(body))
	}

	exchangeAccessTokenResponse := &ExchangeTokenResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&exchangeAccessTokenResponse); err != nil {
		logrus.WithFields(logrus.Fields{
			"StatusCode": resp.StatusCode,
			"err":        err,
		}).Error("POSTExchangeCodeForAccessToken:Decode")
		return nil, err
	}

	return exchangeAccessTokenResponse, nil
}

// GetAccountURL ...
func GetAccountURL(name string) string {
	discoveryResponse := GetDiscovery()

	// v1.1 or v2.0
	specVersion := 2
	specVersionIndex := specVersion - 1

	accountsURL := discoveryResponse.Data["AccountAndTransactionAPI"][specVersionIndex].Links[name]
	logrus.WithFields(logrus.Fields{
		"name":        name,
		"accountsURL": fmt.Sprintf("%+v", accountsURL),
	}).Info("GetAccountURL")

	return accountsURL
}

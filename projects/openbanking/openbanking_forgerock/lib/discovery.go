package lib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetDiscovery calls https://rs.aspsp.ob.forgerock.financial:443/open-banking/discovery and
// returns a `DiscoveryResponse` with values populated.
func GetDiscovery() *DiscoveryResponse {
	// curl -X GET \
	//   https://rs.aspsp.ob.forgerock.financial:443/open-banking/discovery \
	//   -H 'Cache-Control: no-cache' \
	//   -H 'Postman-Token: ddb09010-ff5f-4fce-9302-8d2aafab5afd'
	url := "https://rs.aspsp.ob.forgerock.financial:443/open-banking/discovery"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
			"req": fmt.Sprintf("%+v", req),
		}).Fatal("GetDiscovery:NewRequest")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("GetDiscovery")
	}

	discoveryResponse := &DiscoveryResponse{}
	if err := json.NewDecoder(res.Body).Decode(&discoveryResponse); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("GetDiscovery")
	}

	logrus.WithFields(logrus.Fields{
		"discoveryResponse": fmt.Sprintf("%+v", discoveryResponse),
	}).Debug("GetDiscovery")

	return discoveryResponse
}

// DiscoveryResponse is
//
// {
// 	"Data": {
// 		"PaymentInitiationAPI": [
// 			{
// 				"Version": "v1.1",
// 				"Links": {
// 					"CreateSingleImmediatePayment": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/payments",
// 					"GetSingleImmediatePayment": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/payments/{PaymentId}",
// 					"CreatePaymentSubmission": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/payment-submissions",
// 					"GetPaymentSubmission": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/payment-submissions/{PaymentSubmissionId}"
// 				}
// 			},
// 			{
// 				"Version": "v2.0",
// 				"Links": {
// 					"CreateSingleImmediatePayment": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/payments",
// 					"GetSingleImmediatePayment": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/payments/{PaymentId}",
// 					"CreatePaymentSubmission": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/payment-submissions",
// 					"GetPaymentSubmission": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/payment-submissions/{PaymentSubmissionId}"
// 				}
// 			}
// 		],
// 		"AccountAndTransactionAPI": [
// 			{
// 				"Version": "v1.1",
// 				"Links": {
// 					"CreateAccountRequest": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/account-requests",
// 					"GetAccountRequest": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/account-requests/{AccountRequestId}",
// 					"DeleteAccountRequest": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/account-requests/{AccountRequestId}",
// 					"GetAccounts": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/accounts",
// 					"GetAccount": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/accounts/{AccountId}",
// 					"GetAccountTransactions": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/accounts/{AccountId}/transactions",
// 					"GetAccountBeneficiaries": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/accounts/{AccountId}/beneficiaries",
// 					"GetAccountBalances": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/accounts/{AccountId}/balances",
// 					"GetAccountDirectDebits": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/accounts/{AccountId}/direct-debits",
// 					"GetAccountStandingOrders": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/accounts/{AccountId}/standing-orders",
// 					"GetAccountProduct": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/accounts/{AccountId}/product",
// 					"GetStandingOrders": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/standing-orders",
// 					"GetDirectDebits": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/direct-debits",
// 					"GetBeneficiaries": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/beneficiaries",
// 					"GetTransactions": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/transactions",
// 					"GetBalances": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/balances",
// 					"GetProducts": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/products"
// 				}
// 			},
// 			{
// 				"Version": "v2.0",
// 				"Links": {
// 					"CreateAccountRequest": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/account-requests",
// 					"GetAccountRequest": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/account-requests/{AccountRequestId}",
// 					"DeleteAccountRequest": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/account-requests/{AccountRequestId}",
// 					"GetAccounts": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts",
// 					"GetAccount": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}",
// 					"GetAccountTransactions": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/transactions",
// 					"GetAccountBeneficiaries": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/beneficiaries",
// 					"GetAccountBalances": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/balances",
// 					"GetAccountDirectDebits": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/direct-debits",
// 					"GetAccountStandingOrders": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/standing-orders",
// 					"GetAccountProduct": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/product",
// 					"GetStandingOrders": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/standing-orders",
// 					"GetDirectDebits": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/direct-debits",
// 					"GetBeneficiaries": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/beneficiaries",
// 					"GetTransactions": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/transactions",
// 					"GetBalances": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/balances",
// 					"GetProducts": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/products",
// 					"GetAccountOffers": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/offers",
// 					"GetAccountParty": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/party",
// 					"GetAccountScheduledPayments": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/scheduled-payments",
// 					"GetAccountStatements": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/statements",
// 					"GetAccountStatement": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/statements/{StatementId}",
// 					"GetAccountStatementFile": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/statements/{StatementId}/file",
// 					"GetAccountStatementTransactions": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/accounts/{AccountId}/statements/{StatementId}/transactions",
// 					"GetOffers": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/offers",
// 					"GetParty": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/party",
// 					"GetScheduledPayments": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/scheduled-payments",
// 					"GetStatement": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/statements"
// 				}
// 			}
// 		]
// 	}
// }
type DiscoveryResponse struct {
	Data map[string][]DiscoveryResponseData `json:"Data"`
}

// DiscoveryResponseData is the value one of keys of the maps take. In the example below
// `DiscoveryResponseData` is the right-hand side.
// "PaymentInitiationAPI": [
// 	{
// 		"Version": "v1.1",
// 		"Links": {
// 			"CreateSingleImmediatePayment": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/payments",
// 			"GetSingleImmediatePayment": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/payments/{PaymentId}",
// 			"CreatePaymentSubmission": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/payment-submissions",
// 			"GetPaymentSubmission": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/payment-submissions/{PaymentSubmissionId}"
// 		}
// 	},
// 	{
// 		"Version": "v2.0",
// 		"Links": {
// 			"CreateSingleImmediatePayment": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/payments",
// 			"GetSingleImmediatePayment": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/payments/{PaymentId}",
// 			"CreatePaymentSubmission": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/payment-submissions",
// 			"GetPaymentSubmission": "https://rs.aspsp.ob.forgerock.financial:443/open-banking/v2.0/payment-submissions/{PaymentSubmissionId}"
// 		}
// 	}
// ]
type DiscoveryResponseData struct {
	Version string            `json:"Version"`
	Links   map[string]string `json:"Links"`
}

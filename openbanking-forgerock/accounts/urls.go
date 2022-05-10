package accounts

import "fmt"

// GetAccountURL ...
func GetAccountURL(name string) string {
	if name == "GetAccounts" {
		return "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/v3.1/aisp/accounts"
	}
	if name == "CreateAccountAccessConsent" {
		return "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/v3.1/aisp/account-access-consents"
	}
	// "GetAccountStatements": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/statements",
	if name == "GetAccountStatements" {
		return "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/v3.1/aisp/accounts/{AccountId}/statements"
	}
	if name == "GetBalances" {
		return "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/v3.1/aisp/balances"
	}

	panic(fmt.Errorf("GetAccountURL: name=%q unrecognised", name))

	// discoveryResponse := GetDiscovery()

	// // v1.1 or v2.0
	// specVersion := 2
	// specVersionIndex := specVersion - 1

	// accountsURL := discoveryResponse.Data["AccountAndTransactionAPI"][specVersionIndex].Links[name]
	// logrus.WithFields(logrus.Fields{
	// 	"name":        name,
	// 	"accountsURL": fmt.Sprintf("%+v", accountsURL),
	// }).Info("GetAccountURL")

	// return accountsURL
}

// {
//     "Version": "v3.1",
//     "Links": {
//         "@type": "GenericOBDiscoveryAPILinks",
//         "links": {
//             "GetParty": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/party",
//             "GetAccountStatements": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/statements",
//             "GetScheduledPayments": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/scheduled-payments",
//             "CreateAccountAccessConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/account-access-consents",
//             "GetProducts": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/products",
//             "GetTransactions": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/transactions",
//             "GetAccountTransactions": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/transactions",
//             "GetAccountScheduledPayments": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/scheduled-payments",
//             "GetAccountStatementFile": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/statements/{StatementId}/file",
//             "GetAccountStandingOrders": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/standing-orders",
//             "GetBalances": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/balances",
//             "GetAccount": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}",
//             "GetAccountProduct": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/product",
//             "GetOffers": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/offers",
//             "GetBeneficiaries": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/beneficiaries",
//             "GetAccountStatementTransactions": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/statements/{StatementId}/transactions",
//             "GetAccountOffers": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/offers",
//             "GetAccountDirectDebits": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/direct-debits",
//             "GetAccountAccessConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/account-access-consents/{ConsentId}",
//             "GetStatements": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/statements",
//             "GetAccountParty": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/party",
//             "GetAccountBalances": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/balances",
//             "DeleteAccountAccessConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/account-access-consents/{ConsentId}",
//             "GetStandingOrders": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/standing-orders",
//             "GetAccountStatement": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/accounts/{AccountId}/statements/{StatementId}",
//             "GetDirectDebits": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/aisp/direct-debits"
//         }
//     }
// }

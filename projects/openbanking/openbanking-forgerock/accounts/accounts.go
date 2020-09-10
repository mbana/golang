package accounts

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
// {
//   "Data": {
//     "ConsentId": "urn-alphabank-intent-88379",
//     "Status": "AwaitingAuthorisation",,
//     "StatusUpdateDateTime": "2017-05-02T00:00:00+00:00"
//     "CreationDateTime": "2017-05-02T00:00:00+00:00",
//     "Permissions": [
//       "ReadAccountsDetail",
//       "ReadBalances",
//       "ReadBeneficiariesDetail",
//       "ReadDirectDebits",
//       "ReadProducts",
//       "ReadStandingOrdersDetail",
//       "ReadTransactionsCredits",
//       "ReadTransactionsDebits",
//       "ReadTransactionsDetail",
//       "ReadOffers",
//       "ReadPAN",
//       "ReadParty",
//       "ReadPartyPSU",
//       "ReadScheduledPaymentsDetail",
//       "ReadStatementsDetail"
//     ],
//     "ExpirationDateTime": "2017-08-02T00:00:00+00:00",
//     "TransactionFromDateTime": "2017-05-03T00:00:00+00:00",
//     "TransactionToDateTime": "2017-12-03T00:00:00+00:00"

//   },
//   "Risk": {},
//   "Links": {
//     "Self": "https://api.alphabank.com/open-banking/v3.1/aisp/account-access-consents/urn-alphabank-intent-88379"
//   },
//   "Meta": {
//     "TotalPages": 1
//   }
// }
type AccountRequestsResponseData struct {
	ConsentID               string   `json:"ConsentId"`
	Status                  string   `json:"Status"`
	StatusUpdateDateTime    string   `json:"StatusUpdateDateTime"`
	CreationDateTime        string   `json:"CreationDateTime"`
	Permissions             []string `json:"Permissions"`
	ExpirationDateTime      string   `json:"ExpirationDateTime"`
	TransactionFromDateTime string   `json:"TransactionFromDateTime"`
	TransactionToDateTime   string   `json:"TransactionToDateTime"`
}

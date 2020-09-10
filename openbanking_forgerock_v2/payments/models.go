package payments

type CreateDomesticPaymentConsentResponse struct {
	Data CreateDomesticPaymentConsentResponseData `json:"Data"`
	Risk map[string]string                        `json:"Risk"`
}

type CreateDomesticPaymentConsentResponseData struct {
	ConsentID string `json:"ConsentId"`
}

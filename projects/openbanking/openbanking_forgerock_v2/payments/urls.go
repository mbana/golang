package payments

import "fmt"

// {
//     "Version": "v3.1",
//     "Links": {
//         "@type": "GenericOBDiscoveryAPILinks",
//         "links": {
//             "GetDomesticPaymentConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-payment-consents/{ConsentId}",
//             "CreateFilePayment": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/file-payments",
//             "CreateDomesticPaymentConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-payment-consents",
//             "GetInternationalStandingOrder": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-standing-orders/{InternationalStandingOrderPaymentId}",
//             "GetFilePaymentFile": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/file-payment-consents/{ConsentId}/file",
//             "GetInternationalScheduledPaymentConsentsConsentIdFundsConfirmation": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-scheduled-payment-consents/{ConsentId}/funds-confirmation",
//             "CreateFilePaymentFile": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/file-payment-consents/{ConsentId}/file",
//             "CreateDomesticStandingOrder": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-standing-orders",
//             "GetDomesticPayment": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-payments/{DomesticPaymentId}",
//             "GetDomesticPaymentConsentsConsentIdFundsConfirmation": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-payment-consents/{ConsentId}/funds-confirmation",
//             "GetFilePaymentReport": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/file-payments/{FilePaymentId}/report-file",
//             "GetInternationalPaymentConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-payment-consents/{ConsentId}",
//             "CreateDomesticStandingOrderConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-standing-order-consents",
//             "CreateFilePaymentConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/file-payment-consents",
//             "CreateDomesticPayment": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-payments",
//             "GetInternationalPayment": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-payments/{InternationalPaymentId}",
//             "CreateInternationalStandingOrder": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-standing-orders",
//             "CreateInternationalPayment": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-payments",
//             "CreateInternationalStandingOrderConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-standing-order-consents",
//             "CreateDomesticScheduledPayment": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-scheduled-payments",
//             "GetInternationalScheduledPaymentConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-scheduled-payment-consents/{ConsentId}",
//             "GetInternationalPaymentConsentsConsentIdFundsConfirmation": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-payment-consents/{ConsentId}/funds-confirmation",
//             "GetInternationalStandingOrderConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-standing-order-consents/{ConsentId}",
//             "GetFilePaymentConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/file-payment-consents/{ConsentId}",
//             "CreateInternationalScheduledPaymentConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-scheduled-payment-consents",
//             "GetDomesticScheduledPayment": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-scheduled-payments/{DomesticScheduledPaymentId}",
//             "GetFilePayment": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/file-payments/{FilePaymentId}",
//             "CreateInternationalPaymentConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-payment-consents",
//             "GetDomesticScheduledPaymentConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-scheduled-payment-consents/{ConsentId}",
//             "CreateInternationalScheduledPayment": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-scheduled-payments",
//             "GetInternationalScheduledPayment": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/international-scheduled-payments/{InternationalScheduledPaymentId}",
//             "CreateDomesticScheduledPaymentConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-scheduled-payment-consents",
//             "GetDomesticStandingOrderConsent": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-standing-order-consents/{ConsentId}",
//             "GetDomesticStandingOrder": "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/open-banking/v3.1/pisp/domestic-standing-orders/{DomesticStandingOrderId}"
//         }
//     }
// }
func GetPaymentsURL(name string) string {
	if name == "CreateDomesticPayment" {
		return "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/v3.1/pisp/domestic-payments"
	}
	if name == "CreateDomesticPaymentConsent" {
		return "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/v3.1/pisp/domestic-payment-consents"
	}
	if name == "GetDomesticPayment" {
		return "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/v3.1/pisp/domestic-payments/{DomesticPaymentId}"
	}
	if name == "GetDomesticPaymentConsent" {
		return "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/v3.1/pisp/domestic-payment-consents/{ConsentId}"
	}
	if name == "GetDomesticPaymentConsentsConsentIdFundsConfirmation" {
		return "https://matls.rs.aspsp.ob.forgerock.financial/open-banking/v3.1/pisp/domestic-payment-consents/{ConsentId}/funds-confirmation"
	}

	panic(fmt.Errorf("GetAccountURL: name=%q unrecognised", name))
}

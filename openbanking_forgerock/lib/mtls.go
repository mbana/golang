package lib

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// MTLSTestResponseValid example
//
// {
//     "issuerId": "...",
//     "authorities": [
//         {
//             "authority": "AISP"
//         },
//         {
//             "authority": "PISP"
//         }
//     ]
// }
type MTLSTestResponseValid struct {
	IssuerID    string      `json:"issuerId"`
	Authorities []Authority `json:"authorities"`
}

// MTLSTestResponseInvalid example
//
// {
//     "issuerId": "...",
//     "authorities": [
//         {
//             "authority": "UNREGISTERED_TPP"
//         }
//     ]
// }
type MTLSTestResponseInvalid struct {
	IssuerID    string      `json:"issuerId"`
	Authorities []Authority `json:"authorities"`
}

// Authority example, see the `authorities` field.
//
// {
//     "issuerId": "...",
//     "authorities": [
//         {
//             "authority": "AISP"
//         },
//         {
//             "authority": "OB_CERTIFICATE"
//         },
//         {
//             "authority": "PISP"
//         }
//     ]
// }
type Authority struct {
	Authority string `json:"authority"`
}

// MTLSTest calls `https://rs.aspsp.ob.forgerock.financial/open-banking/mtlsTest` to conduct
// an MTLS test.
//
// The return should contain something like:
// {
//     "issuerId": "...",
//     "authorities": [
//         {
//             "authority": "AISP"
//         },
//         {
//             "authority": "OB_CERTIFICATE"
//         },
//         {
//             "authority": "PISP"
//         }
//     ]
// }
//
// or the below in the event of a failure:
//
// {
//     "issuerId": "...",
//     "authorities": [
//         {
//             "authority": "UNREGISTERED_TPP"
//         }
//     ]
// }
func (c *OpenBankingClient) MTLSTest() (string, error) {
	url := "https://rs.aspsp.ob.forgerock.financial/open-banking/mtlsTest"
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	logrus.WithFields(logrus.Fields{
		"StatusCode": resp.StatusCode,
		"Body":       string(body),
	}).Info("MTLSTest")

	return string(body), nil
}

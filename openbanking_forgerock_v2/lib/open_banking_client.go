package lib

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/banaio/openbankingforgerock/config"
	"github.com/banaio/openbankingforgerock/oidcdiscovery"
	"github.com/banaio/openbankingforgerock/oidcdynamicclientregistration"
	"github.com/banaio/openbankingforgerock/requests"
	"github.com/banaio/openbankingforgerock/signer"
	"github.com/banaio/openbankingforgerock/utils"
)

// OpenBankingClient is the interface to the Open Banking apis.
type OpenBankingClient struct {
	HTTPClient       *http.Client
	Config           *config.Config
	Keys             *signer.Keys
	Signer           *signer.Signer
	OpenIDConfig     *oidcdiscovery.OpenIDConfiguration
	RegisterResponse *oidcdynamicclientregistration.Response // result of calling https://rs.aspsp.ob.forgerock.financial/open-banking/registerTPP
	AccessToken      *AccessTokenResponse
}

// NewClient creates a new instance of OpenBankingClient.
func NewClient(config *config.Config) (*OpenBankingClient, error) {
	keys, err := signer.NewKeys(config)
	if err != nil {
		return nil, err
	}
	signer, err := signer.NewSigner(keys, config)
	if err != nil {
		return nil, err
	}
	httpClient, err := createSecureClient(keys)
	if err != nil {
		return nil, err
	}
	openIDConfig, err := oidcdiscovery.GetWellKnownOpenIDConfiguration(httpClient)
	if err != nil {
		return nil, err
	}

	return &OpenBankingClient{
		HTTPClient:       httpClient,
		Config:           config,
		Keys:             keys,
		Signer:           signer,
		OpenIDConfig:     openIDConfig,
		RegisterResponse: &oidcdynamicclientregistration.Response{},
		AccessToken:      &AccessTokenResponse{},
	}, nil
}

func createSecureClient(keys *signer.Keys) (*http.Client, error) {
	// the CertPool wants to add a root as a []byte so we read the file ourselves
	issuingCAFile := "./keys/certs/OB_SandBox_PP_Issuing CA.cer"
	issuingCA, err := ioutil.ReadFile(issuingCAFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file: issuingCAFile=%q", issuingCAFile)
	}

	rootCAFile := "./keys/certs/OB_SandBox_PP_Root CA.cer"
	rootCA, err := ioutil.ReadFile(rootCAFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file: rootCAFile=%q", rootCAFile)
	}

	forgerockCAFile := "./keys/certs/ForgeRock-directory-issuer.pem"
	forgerockCA, err := ioutil.ReadFile(forgerockCAFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file: forgerockCAFile=%q", forgerockCAFile)
	}

	// rootCAs := x509.NewCertPool()
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.Wrap(err, "failed in createSecureClient on x509.SystemCertPool")
	}
	if ok := rootCAs.AppendCertsFromPEM(issuingCA); !ok {
		logrus.WithFields(logrus.Fields{
			"issuingCA": string(issuingCA[:20]),
		}).Warn("createSecureClient:pool.AppendCertsFromPEM")
		// return nil, errors.Errorf("failed to pool.AppendCertsFromPEM: issuingCAFile=%q", rootCAFile)
	}
	if ok := rootCAs.AppendCertsFromPEM(rootCA); !ok {
		logrus.WithFields(logrus.Fields{
			"rootCA": string(rootCA[:20]),
		}).Warn("createSecureClient:pool.AppendCertsFromPEM")
		// return nil, errors.Errorf("failed to pool.AppendCertsFromPEM: rootCAFile=%q", rootCAFile)
	}
	if ok := rootCAs.AppendCertsFromPEM(forgerockCA); !ok {
		logrus.WithFields(logrus.Fields{
			"forgerockCA": string(forgerockCA[:20]),
		}).Warn("createSecureClient:pool.AppendCertsFromPEM")
		// return nil, errors.Errorf("failed to pool.AppendCertsFromPEM: rootCAFile=%q", rootCAFile)
	}
	// if ok := rootCAs.AppendCertsFromPEM(keys.Transport.Certificate.Certificate[0]); !ok {
	// 	logrus.WithFields(logrus.Fields{
	// 		"keys.Transport.Certificate.Certificate[0]": string(keys.Transport.Certificate.Certificate[0][:20]),
	// 	}).Warn("createSecureClient:pool.AppendCertsFromPEM")
	// 	// return nil, errors.Errorf("failed to pool.AppendCertsFromPEM: rootCAFile=%q", rootCAFile)
	// }

	transportCert := keys.Transport.Certificate
	tlsConfig := tls.Config{
		RootCAs: rootCAs,
		Certificates: []tls.Certificate{
			transportCert,
		},
		InsecureSkipVerify: false,
		Renegotiation:      tls.RenegotiateFreelyAsClient,
	}
	tlsConfig.BuildNameToCertificate()

	// Prints outs the requests and responses
	transport := &requests.DebugTransport{
		Transport: http.Transport{
			TLSClientConfig: &tlsConfig,
		},
	}

	// Does _not_ print outs the requests and responses
	// transport := &http.Transport{
	// 	TLSClientConfig: &tlsConfig,
	// }

	return &http.Client{
		Transport: transport,
		Timeout:   time.Second * 30,
	}, nil
}

// SetRegisterResponse get the registration response and set `c.RegisterResponse` to it.
func (c *OpenBankingClient) SetRegisterResponse() error {
	if _, err := os.Stat(RegisterResponseFile); os.IsNotExist(err) {
		logrus.WithFields(logrus.Fields{
			"RegisterResponseFile": RegisterResponseFile,
			"err":                  err,
		}).Warn("main:RegisterTPPResponse")

		c.UnRegister()
		logrus.Infoln("sleeping (1 seconds)")
		time.Sleep(1 * time.Second)
		if _, err := c.MTLSTest(); err != nil {
			return err
		}

		registerResponse, err := c.Register()
		if err != nil {
			return err
		}
		c.RegisterResponse = registerResponse

		// logrus.Warnln("sleeping (5 seconds)")
		// time.Sleep(5 * time.Second)

		if _, err := c.MTLSTest(); err != nil {
			return err
		}
	} else {
		utils.ReadJSONFromFile(RegisterResponseFile, c.RegisterResponse)
	}

	return nil
}

// MTLSTest calls `https://rs.aspsp.ob.forgerock.financial/open-banking/mtlsTest` to conduct
// an MTLS test.
//
// The return value should contain something like:
//
//     {
//         "issuerId": "...",
//         "authorities": [
//             {
//                 "authority": "AISP"
//             },
//             {
//                 "authority": "OB_CERTIFICATE"
//             },
//             {
//                 "authority": "PISP"
//             }
//         ]
//     }
//
// or the below in the event of a failure:
//
//     {
//         "issuerId": "...",
//         "authorities": [
//             {
//                 "authority": "UNREGISTERED_TPP"
//             }
//         ]
//     }
func (c *OpenBankingClient) MTLSTest() (string, error) {
	url := "https://rs.aspsp.ob.forgerock.financial/open-banking/mtlsTest"
	res, err := c.HTTPClient.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.Errorf("failed on StatusCode: %d (got) != %d (want), body=%+v", res.StatusCode, http.StatusOK, string(body))
	}

	return string(body), nil
}

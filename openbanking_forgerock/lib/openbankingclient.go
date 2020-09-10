package lib

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// OpenBankingClient is the interface to the Open Banking apis.
type OpenBankingClient struct {
	HTTPClient       *http.Client
	OpenIDConfig     *OpenIDConfig
	Envs             *Envs
	RegisterResponse *RegisterTPPResponse // result of calling https://rs.aspsp.ob.forgerock.financial/open-banking/registerTPP
	AccessToken      *AccessTokenResponse
}

// New creates a new instance of OpenBankingClient.
func New() (*OpenBankingClient, error) {
	// call GetEnvs to validate the envs
	envs, err := GetEnvs()
	if err != nil {
		return nil, err
	}
	openIDConfig, err := GetOpenIDConfig()
	if err != nil {
		return nil, err
	}
	httpClient := createSecureClient()

	return &OpenBankingClient{
		HTTPClient:       httpClient,
		OpenIDConfig:     openIDConfig,
		Envs:             envs,
		RegisterResponse: &RegisterTPPResponse{},
		AccessToken:      &AccessTokenResponse{},
	}, nil
}

// SetRegisterResponse get the registration response and
// set `c.RegisterResponse` to it.
func (c *OpenBankingClient) SetRegisterResponse() error {
	if _, err := os.Stat(RegisterResponseFile); os.IsNotExist(err) {
		logrus.WithFields(logrus.Fields{
			"RegisterResponseFile": RegisterResponseFile,
			"err":                  err,
		}).Warn("main:RegisterTPPResponse")

		c.UnRegister()
		logrus.Infoln("sleeping (5 seconds)")
		time.Sleep(5 * time.Second)
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
		ReadJSONFromFile(RegisterResponseFile, c.RegisterResponse)
	}

	return nil
}

// createSecureClient ...
func createSecureClient() *http.Client {
	// the CertPool wants to add a root as a []byte so we read the file ourselves
	issuingCA, err := ioutil.ReadFile(".ignore/keys/OB_SandBox_PP_Issuing CA.cer")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("createSecureClient")
	}

	rootCA, err := ioutil.ReadFile(".ignore/keys/OB_SandBox_PP_Root CA.cer")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("createSecureClient")
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(issuingCA)
	pool.AppendCertsFromPEM(rootCA)

	clientCert, err := tls.LoadX509KeyPair(
		".ignore/keys/transport/2a407fbb-7f37-4fd9-a49b-9994c46fb260.pem",
		".ignore/keys/transport/2a407fbb-7f37-4fd9-a49b-9994c46fb260.key",
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("createSecureClient")
	}

	tlsConfig := tls.Config{
		// RootCAs: pool,
		Certificates: []tls.Certificate{
			clientCert,
		},
		// TODO: need to remove this somehow...
		// InsecureSkipVerify: true,
	}

	transport := http.Transport{
		TLSClientConfig: &tlsConfig,
	}

	return &http.Client{
		Transport: &transport,
	}
}

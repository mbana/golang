package oidcdiscovery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"github.com/banaio/openbankingforgerock/requests"
)

const (
	wellKnownOpenIDConfiguration = "/.well-known/openid-configuration"
)

// GetWellKnownOpenIDConfiguration ...
func GetWellKnownOpenIDConfiguration(client *http.Client) (*OpenIDConfiguration, error) {
	// https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderConfigurationRequest
	url := "https://as.aspsp.ob.forgerock.financial/oauth2" + wellKnownOpenIDConfiguration
	req, err := requests.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed GetWellKnownOpenIDConfiguration: url=%q", url)
	}

	// https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderConfig
	req.Header.Set(requests.HeaderAccept, requests.MIMEApplicationJSON)

	res, err := client.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed GetWellKnownOpenIDConfiguration: url=%q", url)
	}

	// https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderConfigurationResponse
	if statusCode := res.StatusCode; http.StatusOK != statusCode {
		if res.Body != nil {
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, errors.Wrapf(err, "failed GetWellKnownOpenIDConfiguration on ioutil.ReadAll: url=%q, statusCode=%+v, err=%+v", url, statusCode, err)
			}
			defer res.Body.Close()

			return nil, errors.Errorf("failed GetWellKnownOpenIDConfiguration: url=%q, statusCode=%+v, body=%+v", url, statusCode, string(body))
		}

		err := fmt.Errorf("statusCode mismatch: %d (want) != %d (got)", http.StatusOK, statusCode)
		return nil, errors.Wrapf(err, "failed GetWellKnownOpenIDConfiguration: url=%q, err=%+v", url, err)
	}

	config := &OpenIDConfiguration{}
	if err := json.NewDecoder(res.Body).Decode(&config); err != nil {
		return nil, errors.Wrapf(err, "failed GetWellKnownOpenIDConfiguration: url=%q", url)
	}

	return config, nil
}

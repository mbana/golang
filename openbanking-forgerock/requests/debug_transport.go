package requests

import (
	"net/http"
	"net/http/httputil"

	"github.com/banaio/openbankingforgerock/utils"
)

type DebugTransport struct {
	http.Transport
}

func (d *DebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	SetDefaultHeaders(req)

	requestBytes, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		utils.PrintInBox(utils.Red, utils.Red("DebugTransport:httputil.DumpRequest:"), err.Error())
		return nil, err
	}
	utils.PrintInBox(utils.Yellow, utils.Yellow("DebugTransport:*http.Request"), utils.Yellow(string(requestBytes)))

	res, err := d.Transport.RoundTrip(req)
	if err != nil {
		utils.PrintInBox(utils.Red, utils.Red("DebugTransport:d.Transport.RoundTrip:"), err.Error())
		return nil, err
	}

	responseBytes, err := httputil.DumpResponse(res, true)
	if err != nil {
		utils.PrintInBox(utils.Red, utils.Red("DebugTransport:httputil.DumpResponse:"), err.Error())
		return nil, err
	}
	utils.PrintInBox(utils.Yellow, utils.Yellow("DebugTransport:*http.Response"), utils.Yellow(string(responseBytes)))

	return res, err
}

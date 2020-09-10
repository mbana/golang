package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/banaio/golang/openbanking_forgerock/lib"

	"github.com/labstack/echo"
)

func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		DisableTimestamp: true,
		ForceColors:      true,
		DisableColors:    false,
		TimestampFormat:  time.RFC3339,
	})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.StandardLogger().SetNoLock()
}

func main() {
	initLogger()

	done := make(chan error)
	client := startServer(done)

	res, err := client.MTLSTest()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("main:client.MTLSTest")
	}

	logrus.WithFields(logrus.Fields{
		"res": res,
	}).Info("main:client.MTLSTest")

	if err := client.SetRegisterResponse(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("main:client.SetRegisterResponse")
	}

	err = client.GetAccessToken()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("main:client.GetAccessToken")
	}

	// Start the Hybrid Flow, i.e., initiate intent to make an accounts request.
	// We need `Data.AccountRequestID` which is returned as part of this request.
	accountRequestsResponse, err := client.POSTAccountRequests()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("main:POSTAccountRequests")
	}

	// Get the URL we need to load in a browser so that:
	// 1. `code`, `id_token`, `scope` and `state` are returned.
	// 2. we can exchange `code` for an `access_token`.
	// 3. begin the actual accounts request.
	//
	// See the method/handler/function called when user has given
	// consent to access the accounts data and the redirect
	// url is called. The handler is `/openbanking/banaio/forgerock`.
	// See the `startServer` function.
	//
	// Example URL that user will need to load in the browser:
	// https://matls.as.aspsp.ob.forgerock.financial/oauth2/realms/root/realms/openbanking/authorize?client_id=34f64309-433d-4610-95d2-63d2f5253312&nonce=5a6b0d7832a9fb4f80f1170a&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fopenbanking%2Fbanaio%2Fforgerock&request=eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ2YzNmNDlkLTcxMTItNGM1Yy05YzlkLTg0OTI2ZTk5MmM3NCIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL21hdGxzLmFzLmFzcHNwLm9iLmZvcmdlcm9jay5maW5hbmNpYWwvb2F1dGgyL29wZW5iYW5raW5nIiwiY2xhaW1zIjp7ImlkX3Rva2VuIjp7ImFjciI6eyJlc3NlbnRpYWwiOnRydWUsInZhbHVlIjoidXJuOm9wZW5iYW5raW5nOnBzZDI6c2NhIn0sIm9wZW5iYW5raW5nX2ludGVudF9pZCI6eyJlc3NlbnRpYWwiOnRydWUsInZhbHVlIjoiQWJjM2UwOGJjLTcyYzUtNGUzMy1hYjYwLThiZDlhZjhhZGMxNiJ9fSwidXNlcmluZm8iOnsib3BlbmJhbmtpbmdfaW50ZW50X2lkIjp7ImVzc2VudGlhbCI6dHJ1ZSwidmFsdWUiOiJBYmMzZTA4YmMtNzJjNS00ZTMzLWFiNjAtOGJkOWFmOGFkYzE2In19fSwiY2xpZW50X2lkIjoiNTRmNjQzMDktNDMzZC00NjEwLTk1ZDItNjNkMmY1MjUzNDEyIiwiZXhwIjoxNTQwMTk3OTk5LCJpYXQiOjE1NDAxOTc4NzksImlzcyI6IjU0ZjY0MzA5LTQzM2QtNDYxMC05NWQyLTYzZDJmNTI1MzQxMiIsImp0aSI6IjJmODMyMzJjLTA0NmUtNGIyMC05NTc4LWRmMTljOTdhZTNmOSIsIm5vbmNlIjoiNWE2YjBkNzgzMmE5ZmI0ZjgwZjExNzBhIiwicmVkaXJlY3RfdXJpIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL29wZW5iYW5raW5nL2JhbmFpby9mb3JnZXJvY2siLCJyZXNwb25zZV90eXBlIjoiY29kZSBpZF90b2tlbiIsInNjb3BlIjoiYWNjb3VudHMgb3BlbmlkIiwic3RhdGUiOiI1YTZiMGQ3ODMyYTlmYjRmODBmMTE3MGEifQ.KuTvvOC2yz5idjUVH6I7aZlHj0jGuR06zJlNny8D01XoHvm0xA27YXyIwsQO-q0MlMDErBzzU8WMZ3-6wJxWth4thPmSu5zzVAYo7ZWEUDHhlq7YWZkATRintLv0PqUlx_h8r8N2tmtm0UWE2VtxKdRQN1jSD7_kjsw7w_vaP_OwvOA8lGEjU30JW4HxHLfxyeIjHxsTY_dlSiHvWwdmqlwEW9DQJtAYHGboJkX6GBXqV5zEHD4UdtjRYIkyPDAgHqt5smiEzMcuGwJoD2v4vSBEmpEdnmAANgPFxKhNsyNhm7HQXaL6vRLuasgrg7JW9F8iWvw-4BlASAcoBiwKCg&response_type=code+id_token&scope=accounts+openid&state=5a6b0d7832a9fb4f80f1170a
	client.GETAccountsRequestsHybridFlow(*accountRequestsResponse)

	// see the `/openbanking/banaio/forgerock` handler for how the
	// authorise response is handled and how the actual account
	// request is done.

	if err := <-done; err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("main:server")
	}
}

func startServer(done chan error) *lib.OpenBankingClient {
	client, err := lib.New()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("main:lib.New(")
	}

	go func() {
		e := echo.New()
		e.HideBanner = true

		// Handle callback from ForgeRock.
		// We convert the `code`, `id_token`, `scope` and `state` in the query fragment (the # symbol), i.e.,
		// ForgeRock do a call back in this format:
		//     http://localhost:8080/openbanking/banaio/forgerock#code=a052c795-742d-415a-843f-8a4939d740d1&scope=openid%20accounts&id_token=eyJ0eXAiOiJKV1QiLCJraWQiOiJGb2w3SXBkS2VMWm16S3RDRWdpMUxEaFNJek09IiwiYWxnIjoiRVMyNTYifQ.eyJzdWIiOiJtYmFuYSIsImF1ZGl0VHJhY2tpbmdJZCI6IjY5YzZkZmUzLWM4MDEtNGRkMi05Mjc1LTRjNWVhNzdjZWY1NS0xMDMzMDgyIiwiaXNzIjoiaHR0cHM6Ly9tYXRscy5hcy5hc3BzcC5vYi5mb3JnZXJvY2suZmluYW5jaWFsL29hdXRoMi9vcGVuYmFua2luZyIsInRva2VuTmFtZSI6ImlkX3Rva2VuIiwibm9uY2UiOiI1YTZiMGQ3ODMyYTlmYjRmODBmMTE3MGEiLCJhY3IiOiJ1cm46b3BlbmJhbmtpbmc6cHNkMjpzY2EiLCJhdWQiOiI1NGY2NDMwOS00MzNkLTQ2MTAtOTVkMi02M2QyZjUyNTM0MTIiLCJjX2hhc2giOiIxbGt1SEFuaVJDZlZNS2xEc0pxTTNBIiwib3BlbmJhbmtpbmdfaW50ZW50X2lkIjoiQTY5MDA3Nzc1LTcwZGQtNGIyMi1iZmM1LTlkNTI0YTkxZjk4MCIsInNfaGFzaCI6ImZ0OWRrQTdTWXdlb2hlZXpjOGFHeEEiLCJhenAiOiI1NGY2NDMwOS00MzNkLTQ2MTAtOTVkMi02M2QyZjUyNTM0MTIiLCJhdXRoX3RpbWUiOjE1Mzk5NDM3NzUsInJlYWxtIjoiL29wZW5iYW5raW5nIiwiZXhwIjoxNTQwMDMwMTgxLCJ0b2tlblR5cGUiOiJKV1RUb2tlbiIsImlhdCI6MTUzOTk0Mzc4MX0.8bm69KPVQIuvcTlC-p0FGcplTV1LnmtacHybV2PTb2uEgMgrL3JNA0jpT2OYO73r3zPC41mNQlMDvVOUn78osQ&state=5a6b0d7832a9fb4f80f1170a
		//
		// So we convert (see assets/index.html) this query fragment to a form the backend can receive,
		// i.e., convert the query fragment to a form with the body content set to the values of the query fragment.
		e.Static("/openbanking/banaio/forgerock", "assets")

		// This is the callback from Forgerock containing
		// `code`, `id_token`, `scope` and `state` in a form instead of part of the the query fragment.
		e.POST("/openbanking/banaio/forgerock", func(c echo.Context) error {
			// `code`, `id_token`, `scope` and `state` in the
			// from the body instead of in the query fragment.
			authoriseResponse := &lib.AuthoriseResponse{}
			if err := c.Bind(authoriseResponse); err != nil {
				logrus.WithFields(logrus.Fields{
					"err":               err,
					"authoriseResponse": fmt.Sprintf("%+v", authoriseResponse),
					"echo.Context":      fmt.Sprintf("%+v", c),
				}).Error("POST:/openbanking/banaio/forgerock:Bind")
				return err
			}

			logrus.WithFields(logrus.Fields{
				"c":                 fmt.Sprintf("%+v", c),
				"authoriseResponse": fmt.Sprintf("%+v", authoriseResponse),
			}).Info("POST:/openbanking/banaio/forgerock")

			// Exchange `code` for `access_token`
			exchangeAccessTokenResponse, err := client.POSTExchangeCodeForAccessToken(*authoriseResponse)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err":               err,
					"authoriseResponse": fmt.Sprintf("%+v", authoriseResponse),
				}).Error("POST:/openbanking/banaio/forgerock:POSTExchangeCodeForAccessToken")
				return err
			}

			logrus.WithFields(logrus.Fields{
				"exchangeAccessTokenResponse": fmt.Sprintf("%+v", exchangeAccessTokenResponse),
			}).Info("POST:/openbanking/banaio/forgerock")

			// logrus.Warnln("sleeping (5 seconds)")
			// time.Sleep(5 * time.Second)

			// Make the call to the endpoint
			//     https://rs.aspsp.ob.forgerock.financial:443/open-banking/v1.1/accounts.
			// The `access_token` acquired using the Hybrid Flow is used in this step.
			//
			// Example response:
			// {"Data":{"Account":[{"AccountId":"3b0576a9-038d-40ff-9fff-ca74871f9c2b","Currency":"GBP","Nickname":"Bills","Account":{"SchemeName":"SortCodeAccountNumber","Identification":"93163240337365","Name":"mbana","SecondaryIdentification":"69789331"}},{"AccountId":"e447a126-c7ed-4ac6-a88f-06f2a8ed4e3b","Currency":"GBP","Nickname":"Household","Account":{"SchemeName":"SortCodeAccountNumber","Identification":"93345435281245","Name":"mbana"}}]},"Links":{"Self":"https://rs.aspsp.ob.forgerock.financial/open-banking/v1.1/accounts"},"Meta":{"TotalPages":1}}
			resp, err := client.GETAccountRequests(*exchangeAccessTokenResponse)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err,
				}).Error("POST:/openbanking/banaio/forgerock:GETAccountRequests")
				return err
			}

			return c.String(http.StatusOK, resp)
		})

		done <- e.Start(":8080")
	}()

	time.Sleep(1 * time.Millisecond)

	return client
}

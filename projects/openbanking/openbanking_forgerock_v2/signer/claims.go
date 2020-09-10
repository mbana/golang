package signer

// Claims - marshals to below:
//
//    {
//        "id_token": {
//            "acr": {
//                "value": "urn:openbanking:psd2:sca",
//                "essential": true
//            },
//            "openbanking_intent_id": {
//                "value": "<consent_id>",
//                "essential": true
//            }
//        },
//        "userinfo": {
//            "openbanking_intent_id": {
//                "value": "<consent_id>",
//                "essential": true
//            }
//        }
//    }
type Claims struct {
	IDToken  IDToken  `json:"id_token,omitempty"`
	UserInfo UserInfo `json:"user_info,omitempty"`
}

func NewClaims(consentID string) *Claims {
	return &Claims{
		IDToken: IDToken{
			OpenBankingIntentID: OpenBankingIntentID{
				Value:     consentID,
				Essential: true,
			},
			ACR: OpenBankingIntentID{
				Value:     "urn:openbanking:psd2:sca",
				Essential: true,
			},
		},
		UserInfo: UserInfo{
			OpenBankingIntentID: OpenBankingIntentID{
				Value:     consentID,
				Essential: true,
			},
		},
	}
}

type IDToken struct {
	OpenBankingIntentID OpenBankingIntentID `json:"openbanking_intent_id,omitempty"`
	ACR                 OpenBankingIntentID `json:"acr,omitempty"`
}

type UserInfo struct {
	OpenBankingIntentID OpenBankingIntentID `json:"openbanking_intent_id,omitempty"`
}

type OpenBankingIntentID struct {
	Essential bool   `json:"essential"`
	Value     string `json:"value"`
}

var (
	// DefaultGrantTypes ...
	// nolint:gochecknoglobals
	DefaultGrantTypes = []string{
		"authorization_code",
		"refresh_token",
		"client_credentials",
	}
	// DefaultResponseTypes ...
	// nolint:gochecknoglobals
	DefaultResponseTypes = []string{
		"code token id_token",
		"code",
		"code id_token",
		"device_code",
		"id_token",
		"code token",
		"token",
		"token id_token",
	}
	// DefaultRedirectURIS ...
	// nolint:gochecknoglobals
	DefaultRedirectURIS = []string{
		"http://127.0.0.1:8080/openbanking/banaio",
		"http://127.0.0.1:8080/openbanking/banaio/forgerock",
		"http://127.0.0.1:8080/openbanking/banaio/callback",
		"http://127.0.0.1:8080/openbanking/banaio/forgerock/callback",
		"https://127.0.0.1:8080/openbanking/banaio",
		"https://127.0.0.1:8080/openbanking/banaio/forgerock",
		"https://127.0.0.1:8080/openbanking/banaio/callback",
		"https://127.0.0.1:8080/openbanking/banaio/forgerock/callback",

		"http://localhost:8080/openbanking/banaio",
		"http://localhost:8080/openbanking/banaio/forgerock",
		"http://localhost:8080/openbanking/banaio/callback",
		"http://localhost:8080/openbanking/banaio/forgerock/callback",
		"https://localhost:8080/openbanking/banaio",
		"https://localhost:8080/openbanking/banaio/forgerock",
		"https://localhost:8080/openbanking/banaio/callback",
		"https://localhost:8080/openbanking/banaio/forgerock/callback",

		"http://bana.io/openbanking/forgerock",
		"http://bana.io/openbanking/forgerock/callback",
		"https://bana.io/openbanking/forgerock",
		"https://bana.io/openbanking/forgerock/callback",
	}
	// DefaultScopes ...
	// nolint:gochecknoglobals
	DefaultScopes = []string{
		"openid",
		"payments",
		"fundsconfirmations",
		"accounts",
	}
	// DefaultScope ...
	// nolint:gochecknoglobals
	DefaultScope = "openid payments fundsconfirmations accounts"
	// DefaultTokenEndpointAuthMethod ...
	// nolint:gosec, nolint:gochecknoglobals
	DefaultTokenEndpointAuthMethod = "private_key_jwt"
)

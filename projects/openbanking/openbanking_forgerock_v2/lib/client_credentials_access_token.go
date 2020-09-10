package lib

// AccessTokenResponse example
//
// {
//     "access_token": "...",
//     "scope": "openid payments accounts",
//     "id_token": "...",
//     "token_type": "Bearer",
//     "expires_in": 86399
// }
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	IDToken     string `json:"id_token"`
	ExpiresIn   int64  `json:"expires_in"` // relative seconds from now
	Scope       string `json:"scope"`
}

// AuthoriseResponse ...
type AuthoriseResponse struct {
	Code    string `json:"code" form:"code" query:"code"`
	IDToken string `json:"id_token" form:"id_token" query:"id_token"`
	Scope   string `json:"scope" form:"scope" query:"scope"`
	State   string `json:"state" form:"state" query:"state"`
}

// ExchangeTokenResponse ...
type ExchangeTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"` // relative seconds from now
	Nonce        string `json:"nonce"`
}

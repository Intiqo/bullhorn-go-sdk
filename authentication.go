package bullhorn

var ApiToken string
var RestApiUrl string

type AuthParams struct {
	ClientId          string
	ClientSecret      string
	Username          string
	Password          string
	AuthenticationUrl string
	ApiUrl            string
	AccessToken       string
	RefreshToken      string
}

type AccessTokenResponse struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	TokenType    string  `json:"token_type"`
	ExpiresIn    float64 `json:"expires_in"`
}

type RestApiResponse struct {
	RestToken string `json:"BhRestToken"`
	RestUrl   string `json:"restUrl"`
}

type BullhornAuthentication struct {
	AuthorizationCode string
	AccessTokenResponse
	RestApiResponse
}

package bullhorn

var ApiToken = "85dd8576-e87a-4670-892d-208428a54ef6"
var RestApiUrl = "https://rest91.bullhornstaffing.com/rest-services/6wg974/"

type AuthParams struct {
	ClientId          string
	ClientSecret      string
	Username          string
	Password          string
	AuthenticationUrl string
	ApiUrl            string
	AuthorizationCode string
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

type TokenResponse struct {
	AuthorizationCode string
	AccessTokenResponse
	RestApiResponse
}

type PingResponse struct {
	SessionExpires int64 `json:"sessionExpires"`
}

package bullhorn

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/gofrs/uuid"
)

type AuthParams struct {
	ClientId     string
	ClientSecret string
	Username     string
	Password     string

	AuthenticationUrl string
	LoginUrl          string
}

type AccessTokenResponse struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	TokenType    string  `json:"token_type"`
	ExpiresIn    float64 `json:"expires_in"`
}

type RestApiResponse struct {
	ApiToken string `json:"BhRestToken"`
	ApiUrl   string `json:"restUrl"`
}

type TokenResponse struct {
	AuthorizationCode string
	AccessTokenResponse
	RestApiResponse
}

// authenticate Authenticates the client
func authenticate(B Backend, params *AuthParams) (*TokenResponse, error) {
	var authCode string
	var err error
	authCode, err = getAuthorizationCode(B, params)
	if err != nil {
		return nil, err
	}
	authTokenData, err := getAccessAndRefreshTokenFromAuthCode(B, params, authCode)
	if err != nil {
		return nil, err
	}
	restApiData, err := getRestToken(B, params, authTokenData.AccessToken)
	if err != nil {
		return nil, err
	}
	return &TokenResponse{
		AuthorizationCode:   authCode,
		AccessTokenResponse: *authTokenData,
		RestApiResponse:     *restApiData,
	}, nil
}

// getAccessAndRefreshToken Gets authorization code
//
// Extracts the code from the redirected URL sent by Bullhorn
func getAuthorizationCode(B Backend, params *AuthParams) (string, error) {
	query := make(map[string]string)
	query["response_type"] = "code"
	query["action"] = "Login"
	query["username"] = params.Username
	query["password"] = params.Password
	query["client_id"] = params.ClientId
	query["client_secret"] = params.ClientSecret
	id, _ := uuid.NewV4()
	query["state"] = id.String()

	requestUrl := fmt.Sprintf("%s/oauth/authorize", params.AuthenticationUrl)
	rr, _, err := B.Call(requestUrl, "post", nil, query, nil)
	if err != nil && !strings.Contains(err.Error(), "stopped after 0 redirects") {
		return "", err
	}
	if rr == nil {
		return "", errors.New("failed to parse raw response")
	}
	location := rr.RawResponse.Header.Get("location")
	if location == "" {
		return "", errors.New("failed to get authorization code")
	}

	parsedUrl, err := url.Parse(location)
	if err != nil {
		return "", errors.New("failed to get authorization code")
	}

	values, err := url.ParseQuery(parsedUrl.RawQuery)
	if err != nil {
		return "", errors.New("failed to get authorization code")
	}
	authCode := values.Get("code")
	return authCode, nil
}

// getAccessAndRefreshTokenFromAuthCode Gets the access and refresh tokens from authorization code
func getAccessAndRefreshTokenFromAuthCode(B Backend, params *AuthParams, authCode string) (*AccessTokenResponse, error) {
	query := make(map[string]string)
	query["grant_type"] = "authorization_code"
	query["code"] = authCode
	query["client_id"] = params.ClientId
	query["client_secret"] = params.ClientSecret

	requestUrl := fmt.Sprintf("%s/oauth/token", params.AuthenticationUrl)
	_, cr, err := B.Call(requestUrl, "post", nil, query, nil)
	if err != nil {
		return nil, err
	}
	var resp AccessTokenResponse
	err = B.ParseResponse(cr.Data, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// getNewAccessAndRefreshToken Gets new access and refresh tokens from existing refresh token
func getNewAccessAndRefreshToken(B Backend, params *AuthParams, refreshToken string) (*AccessTokenResponse, error) {
	query := make(map[string]string)
	query["grant_type"] = "refresh_token"
	query["refresh_token"] = refreshToken
	query["client_id"] = params.ClientId
	query["client_secret"] = params.ClientSecret

	requestUrl := fmt.Sprintf("%s/oauth/token", params.AuthenticationUrl)
	_, cr, err := B.Call(requestUrl, "post", nil, query, nil)
	if err != nil {
		return nil, err
	}
	var resp AccessTokenResponse
	err = B.ParseResponse(cr.Data, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// getRestToken Gets the rest api token
func getRestToken(B Backend, params *AuthParams, accessToken string) (*RestApiResponse, error) {
	query := make(map[string]string)
	query["access_token"] = accessToken
	query["version"] = "2.0"

	requestUrl := fmt.Sprintf("%s/rest-services/login", params.LoginUrl)
	_, cr, err := B.Call(requestUrl, "post", nil, query, nil)
	if err != nil {
		return nil, err
	}
	var resp RestApiResponse
	err = B.ParseResponse(cr.Data, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

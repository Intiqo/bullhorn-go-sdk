package authentication

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/adwitiyaio/bullhorn"
	"github.com/adwitiyaio/bullhorn/backend"
)

type Client struct {
	B backend.Service
}

// getClient Gets the authentication client
func getClient() Client {
	return Client{backend.GetBackendService()}
}

// NewAuthenticationData Gets all the authentication data required to create API requests on Bullhorn
//
// If you pass the refresh token as part of params, we'll use that to get a new api token,
// if otherwise, we'll obtain the authorization code, get access token and then get a new api token
//
// Ensure that you store the response data somewhere so you can reuse it when the tokens expire
func NewAuthenticationData(params *bullhorn.AuthParams) (*bullhorn.TokenResponse, error) {
	return getClient().NewAuthenticationData(params)
}

// NewAuthenticationData Gets all the authentication data required to create API requests on Bullhorn
//
// If you pass the refresh token as part of params, we'll use that to get a new api token,
// if otherwise, we'll obtain the authorization code, get access token and then get a new api token
//
// Ensure that you store the response data somewhere so you can reuse it when the tokens expire
func (c Client) NewAuthenticationData(params *bullhorn.AuthParams) (*bullhorn.TokenResponse, error) {
	var authCode string
	var err error
	if params.RefreshToken == "" {
		authCode, err = c.getAuthorizationCode(params)
		if err != nil {
			return nil, err
		}
	}
	authTokenData, err := c.getAccessAndRefreshToken(params, authCode)
	if err != nil {
		return nil, err
	}
	restApiData, err := c.getRestToken(params, authTokenData.AccessToken)
	if err != nil {
		return nil, err
	}
	return &bullhorn.TokenResponse{
		AuthorizationCode:   authCode,
		AccessTokenResponse: *authTokenData,
		RestApiResponse:     *restApiData,
	}, nil
}

// getAccessAndRefreshToken Gets authorization code
//
// Extracts the code from the redirected URL sent by Bullhorn
func (c Client) getAuthorizationCode(params *bullhorn.AuthParams) (string, error) {
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
	rr, _, err := c.B.Call(requestUrl, "post", nil, query, nil)
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

// getAccessAndRefreshToken Gets the access and refresh tokens
func (c Client) getAccessAndRefreshToken(params *bullhorn.AuthParams, authCode string) (*bullhorn.AccessTokenResponse, error) {
	query := make(map[string]string)
	if params.RefreshToken == "" {
		query["grant_type"] = "authorization_code"
		query["code"] = authCode
	} else {
		query["grant_type"] = "refresh_token"
		query["refresh_token"] = params.RefreshToken
	}
	query["client_id"] = params.ClientId
	query["client_secret"] = params.ClientSecret

	requestUrl := fmt.Sprintf("%s/oauth/token", params.AuthenticationUrl)
	_, cr, err := c.B.Call(requestUrl, "post", nil, query, nil)
	if err != nil {
		return nil, err
	}
	var resp bullhorn.AccessTokenResponse
	err = c.B.ParseResponse(cr.Data, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// getRestToken Gets the rest api token
func (c Client) getRestToken(params *bullhorn.AuthParams, accessToken string) (*bullhorn.RestApiResponse, error) {
	query := make(map[string]string)
	query["access_token"] = accessToken
	query["version"] = "2.0"

	requestUrl := fmt.Sprintf("%s/rest-services/login", params.ApiUrl)
	_, cr, err := c.B.Call(requestUrl, "post", nil, query, nil)
	if err != nil {
		return nil, err
	}
	var resp bullhorn.RestApiResponse
	err = c.B.ParseResponse(cr.Data, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func Ping(apiUrl string, bhRestToken string) (int64, error) {
	return getClient().Ping(apiUrl, bhRestToken)
}

// Ping Pings the bullhorn backend to check the status of a rest token
//
// Use this service to check the expiration time of your rest token
func (c Client) Ping(apiUrl string, bhRestToken string) (int64, error) {
	headers := make(map[string]string)
	headers["ApiToken"] = bhRestToken

	requestUrl := fmt.Sprintf("%s/ping", apiUrl)
	_, cr, err := c.B.Call(requestUrl, "get", headers, nil, nil)
	if err != nil {
		return 0, err
	}
	if cr.Code >= 400 {
		return 0, err
	}
	var pingResponse bullhorn.PingResponse
	err = c.B.ParseResponse(cr.Data, &pingResponse)
	if err != nil {
		return 0, err
	}
	return pingResponse.SessionExpires, nil
}

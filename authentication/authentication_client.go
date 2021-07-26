package authentication

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/adwitiyaio/bullhorn"
	"github.com/adwitiyaio/bullhorn/backend"
	"github.com/gofrs/uuid"
)

var AccessToken string

type Client struct {
	B backend.Client
}

// New Gets all the authentication data required to create API requests on Bullhorn
func New(params *bullhorn.AuthParams) (*bullhorn.BullhornAuthentication, error) {
	return getClient().NewAuthenticationData(params)
}

// NewAuthenticationData Gets all the authentication data required to create API requests on Bullhorn
//
// If you pass the refresh token as part of params, we'll use that to get a new api token,
// if otherwise, we'll obtain the authorization code, get access token and then get a new api token
//
// Ensure that you store the response data somewhere so you can reuse it when the tokens expire
func (c Client) NewAuthenticationData(params *bullhorn.AuthParams) (*bullhorn.BullhornAuthentication, error) {
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
	return &bullhorn.BullhornAuthentication{
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
	c.B.ParseResponse(cr.Data, &resp)
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
	c.B.ParseResponse(cr.Data, &resp)
	return &resp, nil
}

// getClient Gets the authentication client
func getClient() Client {
	return Client{backend.Getbackend()}
}

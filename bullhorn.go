package bullhorn

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type QueryOptions struct {
	Fields       []string
	Associations []string
	Layout       string
	Meta         string
	ShowEditable string
	Start        int
	Count        int
	OrderBy      string
	Sort         string
}

type Client interface {
	// Ping Pings bullhorn and gets the expiration time of the current api token
	//
	// If the rest api token is invalid, you'll need to obtain a new rest token by using the GetNewRestToken method
	Ping() error

	// GetEntity Gets an entity
	//
	// Name should be one of Entity
	GetEntity(name string, id int, options QueryOptions) (*resty.Response, interface{}, error)
	// CreateEntity Creates a new entity
	//
	// Name should be one of Entity
	//
	// Data should be passed as a map of fields. As for the field names of each entity,
	// see https://bullhorn.github.io/rest-api-docs/entityref.html
	CreateEntity(name string, data map[string]interface{}) (*resty.Response, *CreateResponse, error)
	// UpdateEntity Updates an existing entity
	//
	// Name should be one of Entity
	//
	// Data should be passed as a map of fields. As for the field names of each entity,
	// see https://bullhorn.github.io/rest-api-docs/entityref.html
	UpdateEntity(name string, id int, data map[string]interface{}) (*resty.Response, *UpdateResponse, error)
	// DeleteEntity Deletes an existing entity
	DeleteEntity(name string, id int) (*resty.Response, error)
}

func NewClient(params *AuthParams) (Client, error) {
	b := NewBackend()
	rr, tokenData, err := authenticate(b, params)
	if err != nil {
		fmt.Printf("authentication raw response: %v", rr)
		return nil, err
	}
	c := &bullhornClient{
		B:            b,
		ClientId:     params.ClientId,
		ClientSecret: params.ClientSecret,
		Username:     params.Username,
		Password:     params.Password,

		AuthenticationUrl: params.AuthenticationUrl,
		LoginUrl:          params.LoginUrl,
		ApiUrl:            tokenData.ApiUrl,

		AuthorizationCode: tokenData.AuthorizationCode,
		AccessToken:       tokenData.AccessToken,
		RefreshToken:      tokenData.RefreshToken,
		ApiToken:          tokenData.ApiToken,
	}
	err = c.Ping()
	if err != nil {
		return nil, err
	}
	return c, nil
}

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

// Client ... A client to interact with Bullhorn APIs
//
// See https://bullhorn.github.io/rest-api-docs/entityref.html for entity reference
// See https://bullhorn.github.io/rest-api-docs/index.html for API reference
type Client interface {
	// Ping ... Ping bullhorn and gets the expiration time of the current api token
	//
	// If the rest api token is invalid, you'll need to obtain a new rest token by using the GetNewRestToken method
	Ping() error

	// GetEntity ... Get an entity
	//
	// Name should be a valid Bullhorn Entity name
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#get-entity for more information
	GetEntity(name string, id int, options QueryOptions) (*resty.Response, interface{}, error)
	// QueryEntity ... Query an entity with criteria
	//
	// Name should be a valid Bullhorn Entity name
	// See https://bullhorn.github.io/rest-api-docs/index.html#query for more information
	QueryEntity(name string, query string, options QueryOptions) (*resty.Response, interface{}, error)
	// SearchEntity ... Search an entity with criteria
	//
	// Name should be a valid Bullhorn Entity name
	//
	// Query should be in the following format (lucene) for example:
	//
	// (name:Mel*) AND status:"Active Account"
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#post-search for more information
	SearchEntity(name string, query string, options QueryOptions) (*resty.Response, interface{}, error)
	// CreateEntity Create a new entity
	//
	// Name should be a valid Bullhorn Entity name
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#put-entity for more information
	//
	// Data should be passed as a map of fields. As for the field names of each entity,
	//
	// See https://bullhorn.github.io/rest-api-docs/entityref.html
	CreateEntity(name string, data map[string]interface{}) (*resty.Response, *CreateResponse, error)
	// UpdateEntity ... Update an existing entity
	//
	// Name should be a valid Bullhorn Entity name
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#post-entity for more information
	//
	UpdateEntity(name string, id int, data map[string]interface{}) (*resty.Response, *UpdateResponse, error)
	// DeleteEntity ... Delete an existing entity
	//
	// Name should be a valid Bullhorn Entity name
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#delete-entity for more information
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
		RestToken:         tokenData.RestToken,
	}
	err = c.Ping()
	if err != nil {
		return nil, err
	}
	return c, nil
}

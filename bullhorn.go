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

type GetAssociationInput struct {
	IDs              []int `json:"ids"`
	Count            *int  `json:"count"`
	Start            *int  `json:"start"`
	ShowTotalMatched *bool `json:"showTotalMatched"`
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
	// GetMultipleEntities ... Get multiple entities by their IDs
	//
	// Name should be a valid Bullhorn Entity name
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#multiple-entities for more information
	GetMultipleEntities(name string, ids []int, options QueryOptions) (*resty.Response, interface{}, error)
	// GetAssociations ... Get associated records for an entity
	//
	// Entity should be a valid Bullhorn Entity name
	//
	// Association should be a valid Bullhorn Association name for the entity
	//
	// ShowTotalMatched will be set to true by default, pass false if you don't want to show the total matched records
	//
	// Use Count & Start for pagination
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#association for more information
	GetAssociations(entity string, association string, in GetAssociationInput) (*resty.Response, interface{}, error)
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
	// AssociateEntities Associate entities
	//
	// Name should be a valid Bullhorn Entity name
	//
	// Association should be a valid Bullhorn Association name
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#put-entity for more information
	//
	AssociateEntities(name string, id int, association string, associationIds []string) (
		*resty.Response, *CreateResponse, error,
	)
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
	// SubscribeToEvents ... Subscribe to events for entities and operations
	//
	// Subscription ID must be a unique ID specific to a subscription
	//
	// Event Types can be any / all of "INSERTED", "UPDATED" or "DELETED"
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#put-event-subscription for more information
	SubscribeToEvents(subscriptionId string, entities []string, eventTypes []string) (
		*resty.Response, *SubscribeEventResponse, error,
	)
	// UnsubscribeFromEvents ... Unsubscribe from events
	//
	// Subscription ID must be a unique ID specific to a subscription
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#delete-event-subscription for more information
	UnsubscribeFromEvents(subscriptionId string, eventTypes []string) (
		*resty.Response, *UnsubscribeEventResponse, error,
	)
	// FetchEvents ... Fetch all the latest events for a specific subscription
	//
	// Size indicates the number of events you'd like to receive in one go
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#get-event-subscription for more information
	FetchEvents(subscriptionId string, size uint64) (*resty.Response, *FetchEventResponse, error)
	// GetAttachmentsForEntity ... Get attachments for an entity
	//
	// Name should be a valid Bullhorn Entity name
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#get-entity-entitytype-entityid-fileattachments for more information
	GetAttachmentsForEntity(name string, entityId int, options QueryOptions) (*resty.Response, interface{}, error)
	// GetFileForEntity ... Get file details for an entity
	//
	// Name should be a valid Bullhorn Entity name
	//
	// See https://bullhorn.github.io/rest-api-docs/index.html#get-my-entity-s for more information
	GetFileForEntity(name string, entityId int, fileId int) (*resty.Response, interface{}, error)
	// ParseResponseForEntity ... Parse response data into an entity struct
	//
	// Name should be a valid Bullhorn Entity name
	//
	// In case you have "associations" with a request, pass that here
	// so that the response can be parsed into the associated entity
	//
	// In case the expected response is an array, pass the isArray as true to get the
	// response data in a slice of entities
	ParseResponseForEntity(name string, data interface{}, associations []string, isArray bool) (interface{}, error)
}

func NewClient(params *AuthParams) (Client, error) {
	b := NewBackend()
	// If the caller is already authenticated, we ping the server to validate the token, if not valid,
	// we create a new token for the caller
	if params.AuthorizationCode != "" && params.AccessToken != "" && params.RefreshToken != "" && params.
		RestToken != "" && params.ApiUrl != "" {
		c := &bullhornClient{
			B:            b,
			ClientId:     params.ClientId,
			ClientSecret: params.ClientSecret,
			Username:     params.Username,
			Password:     params.Password,

			AuthenticationUrl: params.AuthenticationUrl,
			LoginUrl:          params.LoginUrl,
			ApiUrl:            params.ApiUrl,

			AuthorizationCode: params.AuthorizationCode,
			AccessToken:       params.AccessToken,
			RefreshToken:      params.RefreshToken,
			RestToken:         params.RestToken,
		}
		err := c.Ping()
		// If the ping fails, we establish the entire authentication process again
		if err != nil {
			c, err := authenticateAndInitialize(b, params)
			if err != nil {
				return nil, err
			}
			return c, nil
		}
		return c, nil
	}
	// If the caller doesn't have any existing tokens, we create a new token for the caller
	c, err := authenticateAndInitialize(b, params)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func authenticateAndInitialize(b Backend, params *AuthParams) (Client, error) {
	rr, tokenData, err := authenticate(b, params)
	if err != nil {
		fmt.Printf("authentication raw response: %v", rr)
		return nil, err
	}
	params.ApiUrl = tokenData.ApiUrl
	params.AccessToken = tokenData.AccessToken
	params.AuthorizationCode = tokenData.AuthorizationCode
	params.RefreshToken = tokenData.RefreshToken
	params.RestToken = tokenData.RestToken
	c := &bullhornClient{
		B:            b,
		ClientId:     params.ClientId,
		ClientSecret: params.ClientSecret,
		Username:     params.Username,
		Password:     params.Password,

		AuthenticationUrl: params.AuthenticationUrl,
		LoginUrl:          params.LoginUrl,
		ApiUrl:            tokenData.ApiUrl,

		AuthorizationCode: params.AuthorizationCode,
		AccessToken:       params.AccessToken,
		RefreshToken:      params.RefreshToken,
		RestToken:         params.RestToken,
		RestTokenTtl:      params.RestTokenTTL,
	}
	err = c.Ping()
	if err != nil {
		return nil, err
	}
	return c, nil
}

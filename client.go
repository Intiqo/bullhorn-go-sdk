package bullhorn

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type bullhornClient struct {
	B Backend

	ClientId     string
	ClientSecret string
	Username     string
	Password     string

	AuthenticationUrl string
	LoginUrl          string
	ApiUrl            string

	AuthorizationCode   string
	AccessToken         string
	RefreshToken        string
	RestToken           string
	RestTokenTtl        string
	RestTokenExpiryTime int64
}

func (b *bullhornClient) getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["BhRestToken"] = b.RestToken
	return headers
}

func (b *bullhornClient) validateEntity(name string) error {
	switch name {
	case CountryEntity:
		return nil
	case StateEntity:
		return nil
	case CandidateEntity:
		return nil
	case ClientCorporationEntity:
		return nil
	case ClientContactEntity:
		return nil
	case LocationEntity:
		return nil
	case JobOrderEntity:
		return nil
	case JobSubmissionEntity:
		return nil
	case PlacementEntity:
		return nil
	}
	return fmt.Errorf("unsupported entity %s", name)
}

func (b *bullhornClient) ParseResponseForEntity(
	name string, data interface{}, associations []string,
	isArray bool,
) (interface{}, error) {
	if len(associations) > 0 {
		var resp interface{}
		err := b.B.ParseResponse(data, &resp)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	switch name {
	case CountryEntity:
		var country Country
		var countries []Country
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &countries)
			if err != nil {
				return nil, err
			}
			return countries, nil
		} else {
			err = b.B.ParseResponse(data, &country)
			if err != nil {
				return nil, err
			}
			return country, nil
		}
	case StateEntity:
		var state State
		var states []State
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &states)
			if err != nil {
				return nil, err
			}
			return states, nil
		} else {
			err = b.B.ParseResponse(data, &state)
			if err != nil {
				return nil, err
			}
			return state, nil
		}
	case CandidateEntity:
		var candidate Candidate
		var candidates []Candidate
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &candidates)
			if err != nil {
				return nil, err
			}
			return candidates, nil
		} else {
			err = b.B.ParseResponse(data, &candidate)
			if err != nil {
				return nil, err
			}
			return candidate, nil
		}
	case ClientCorporationEntity:
		var clientCorporation ClientCorporation
		var clientCorporations []ClientCorporation
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &clientCorporations)
			if err != nil {
				return nil, err
			}
			return clientCorporations, nil
		} else {
			err = b.B.ParseResponse(data, &clientCorporation)
			if err != nil {
				return nil, err
			}
			return clientCorporation, nil
		}
	case ClientContactEntity:
		var clientContact ClientContact
		var clientContacts []ClientContact
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &clientContacts)
			if err != nil {
				return nil, err
			}
			return clientContacts, nil
		} else {
			err = b.B.ParseResponse(data, &clientContact)
			if err != nil {
				return nil, err
			}
			return clientContact, nil
		}
	case LocationEntity:
		var location Location
		var locations []Location
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &locations)
			if err != nil {
				return nil, err
			}
			return locations, nil
		} else {
			err = b.B.ParseResponse(data, &location)
			if err != nil {
				return nil, err
			}
			return location, nil
		}
	case JobOrderEntity:
		var jobOrder JobOrder
		var jobOrders []JobOrder
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &jobOrders)
			if err != nil {
				return nil, err
			}
			return jobOrders, nil
		} else {
			err = b.B.ParseResponse(data, &jobOrder)
			if err != nil {
				return nil, err
			}
			return jobOrder, nil
		}
	case JobSubmissionEntity:
		var jobSubmission JobSubmission
		var jobSubmissions []JobSubmission
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &jobSubmissions)
			if err != nil {
				return nil, err
			}
			return jobSubmissions, nil
		} else {
			err = b.B.ParseResponse(data, &jobSubmission)
			if err != nil {
				return nil, err
			}
			return jobSubmission, nil
		}
	case PlacementEntity:
		var placement Placement
		var placements []Placement
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &placements)
			if err != nil {
				return nil, err
			}
			return placements, nil
		} else {
			err = b.B.ParseResponse(data, &placement)
			if err != nil {
				return nil, err
			}
			return placement, nil
		}
	}
	return nil, fmt.Errorf("unsupported entity %s", name)
}

func (b *bullhornClient) Ping() error {
	requestUrl := fmt.Sprintf("%s/ping", b.ApiUrl)
	rr, cr, err := b.B.Call(requestUrl, "get", b.getHeaders(), nil, nil)
	if err != nil {
		fmt.Printf("failed to ping bullhorn with error %s\n", err)
		fmt.Printf("raw response -> %v", rr)
		return err
	}
	if cr.Code >= 400 {
		fmt.Printf("failed to ping bullhorn with error %s\n", err)
		fmt.Printf("raw response -> %v", rr)
		return err
	}
	var pingResponse Ping
	err = b.B.ParseResponse(cr.Data, &pingResponse)
	if err != nil {
		fmt.Printf("failed to ping bullhorn with error %s\n", err)
		fmt.Printf("raw response -> %v", rr)
		return err
	}
	b.RestTokenExpiryTime = pingResponse.SessionExpires
	return nil
}

func (b *bullhornClient) GetEntity(name string, id int, options QueryOptions) (*resty.Response, interface{}, error) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, nil, err
	}
	err = b.validateEntity(name)
	if err != nil {
		return nil, nil, err
	}
	params := make(map[string]string)
	params["fields"] = strings.Join(options.Fields[:], ",")

	requestUrl := fmt.Sprintf("%s/entity/%s/%d", b.ApiUrl, name, id)
	if len(options.Associations) > 0 {
		requestUrl = fmt.Sprintf("%s/entity/%s/%d/%s", b.ApiUrl, name, id, strings.Join(options.Associations[:], ","))
	}
	rr, cr, err := b.B.Call(requestUrl, "get", b.getHeaders(), params, nil)
	if err != nil {
		var bhErr Error
		err := json.Unmarshal(rr.Body(), &bhErr)
		if err != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, errors.New(bhErr.Message)
	}
	dataMap := cr.Data.(map[string]interface{})
	return rr, dataMap["data"], nil
}

func (b *bullhornClient) QueryEntity(name string, query string, options QueryOptions) (
	*resty.Response, interface{},
	error,
) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, nil, err
	}
	err = b.validateEntity(name)
	if err != nil {
		return nil, nil, err
	}

	count := 10
	if options.Count != 0 {
		count = options.Count
	}
	start := 0
	if options.Start != 0 {
		start = options.Start
	}
	body := Query{
		Where: query,
		Options: Options{
			Fields: strings.Join(options.Fields[:], ","),
			Count:  count,
			Start:  start,
		},
	}
	if options.OrderBy != "" {
		body.Options.OrderBy = options.OrderBy
	}
	requestUrl := fmt.Sprintf("%s/query/%s", b.ApiUrl, name)
	rr, cr, err := b.B.Call(requestUrl, "post", b.getHeaders(), nil, body)
	if err != nil {
		var bhErr Error
		err := json.Unmarshal(rr.Body(), &bhErr)
		if err != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, errors.New(bhErr.Message)
	}
	dataMap := cr.Data.(map[string]interface{})
	return rr, dataMap["data"], nil
}

func (b *bullhornClient) SearchEntity(name string, query string, options QueryOptions) (
	*resty.Response, interface{}, error,
) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, nil, err
	}
	err = b.validateEntity(name)
	if err != nil {
		return nil, nil, err
	}

	count := 100
	if options.Count != 0 {
		count = options.Count
	}
	start := 0
	if options.Start != 0 {
		start = options.Start
	}
	body := Search{
		Query: query,
		Options: Options{
			Fields: strings.Join(options.Fields[:], ","),
			Count:  count,
			Start:  start,
		},
	}
	if options.OrderBy != "" {
		body.Options.OrderBy = options.OrderBy
	}
	requestUrl := fmt.Sprintf("%s/search/%s", b.ApiUrl, name)
	rr, cr, err := b.B.Call(requestUrl, "post", b.getHeaders(), nil, body)
	if err != nil {
		var bhErr Error
		err := json.Unmarshal(rr.Body(), &bhErr)
		if err != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, errors.New(bhErr.Message)
	}
	dataMap := cr.Data.(map[string]interface{})
	return rr, dataMap["data"], nil
}

func (b *bullhornClient) CreateEntity(name string, data map[string]interface{}) (
	*resty.Response, *CreateResponse, error,
) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, nil, err
	}
	err = b.validateEntity(name)
	if err != nil {
		return nil, nil, err
	}
	requestUrl := fmt.Sprintf("%s/entity/%s", b.ApiUrl, name)
	rr, cr, err := b.B.Call(requestUrl, "put", b.getHeaders(), nil, data)
	if err != nil {
		var bhErr Error
		err := json.Unmarshal(rr.Body(), &bhErr)
		if err != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, errors.New(bhErr.Message)
	}
	var createResponse CreateResponse
	err = b.B.ParseResponse(cr.Data, &createResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &createResponse, nil
}

func (b *bullhornClient) AssociateEntities(
	name string, id int, association string, associationIds []string,
) (*resty.Response, *CreateResponse, error) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, nil, err
	}
	err = b.validateEntity(name)
	if err != nil {
		return nil, nil, err
	}
	requestUrl := fmt.Sprintf(
		"%s/entity/%s/%d/%s/%s", b.ApiUrl, name, id, association, strings.Join(associationIds[:], ","),
	)
	rr, cr, err := b.B.Call(requestUrl, "put", b.getHeaders(), nil, nil)
	if err != nil {
		var bhErr Error
		err := json.Unmarshal(rr.Body(), &bhErr)
		if err != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, errors.New(bhErr.Message)
	}
	var createResponse CreateResponse
	err = b.B.ParseResponse(cr.Data, &createResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &createResponse, nil
}

func (b *bullhornClient) UpdateEntity(name string, id int, data map[string]interface{}) (
	*resty.Response, *UpdateResponse, error,
) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, nil, err
	}
	err = b.validateEntity(name)
	if err != nil {
		return nil, nil, err
	}
	requestUrl := fmt.Sprintf("%s/entity/%s/%d", b.ApiUrl, name, id)
	rr, cr, err := b.B.Call(requestUrl, "post", b.getHeaders(), nil, data)
	if err != nil {
		var bhErr Error
		err := json.Unmarshal(rr.Body(), &bhErr)
		if err != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, errors.New(bhErr.Message)
	}
	var updateResponse UpdateResponse
	err = b.B.ParseResponse(cr.Data, &updateResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &updateResponse, nil
}

func (b *bullhornClient) DeleteEntity(name string, id int) (*resty.Response, error) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, err
	}
	err = b.validateEntity(name)
	if err != nil {
		return nil, err
	}
	requestUrl := fmt.Sprintf("%s/entity/%s/%d", b.ApiUrl, name, id)
	rr, _, err := b.B.Call(requestUrl, "delete", b.getHeaders(), nil, nil)
	if err != nil {
		var bhErr Error
		err := json.Unmarshal(rr.Body(), &bhErr)
		if err != nil {
			return rr, errors.New(string(rr.Body()))
		}
		return rr, errors.New(bhErr.Message)
	}
	return rr, nil
}

func (b *bullhornClient) SubscribeToEvent(subscriptionId string, entities []string, eventTypes []string) (*resty.Response, *SubscribeEventResponse, error) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, nil, err
	}
	for _, ent := range entities {
		err = b.validateEntity(ent)
		if err != nil {
			return nil, nil, err
		}
	}
	for _, et := range eventTypes {
		switch et {
		case "INSERTED":
			fallthrough
		case "UPDATED":
			fallthrough
		case "DELETED":
			break
		default:
			return nil, nil, fmt.Errorf("unsupported event type %s", et)
		}
	}
	params := make(map[string]string)
	params["type"] = "entity"
	params["names"] = strings.Join(entities, ",")
	params["eventTypes"] = strings.Join(eventTypes, ",")
	requestUrl := fmt.Sprintf("%s/event/subscription/%s", b.ApiUrl, subscriptionId)
	rr, cr, err := b.B.Call(requestUrl, "put", b.getHeaders(), params, nil)
	if err != nil {
		var bhErr Error
		err := json.Unmarshal(rr.Body(), &bhErr)
		if err != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, errors.New(bhErr.Message)
	}
	var subscribeResponse SubscribeEventResponse
	err = b.B.ParseResponse(cr.Data, &subscribeResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &subscribeResponse, nil
}

func (b *bullhornClient) FetchEvents(subscriptionId string, size uint64) (*resty.Response, *FetchEventResponse, error) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, nil, err
	}
	params := make(map[string]string)
	params["maxEvents"] = strconv.FormatUint(size, 10)

	requestUrl := fmt.Sprintf("%s/event/subscription/%s", b.ApiUrl, subscriptionId)
	rr, cr, err := b.B.Call(requestUrl, "get", b.getHeaders(), params, nil)
	if err != nil {
		var bhErr Error
		err := json.Unmarshal(rr.Body(), &bhErr)
		if err != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, errors.New(bhErr.Message)
	}
	var fetchEventResponse FetchEventResponse
	err = b.B.ParseResponse(cr.Data, &fetchEventResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &fetchEventResponse, nil
}

func (b *bullhornClient) updateTokensForClient() error {
	params := &AuthParams{
		ClientId:          b.ClientId,
		ClientSecret:      b.ClientSecret,
		Username:          b.Username,
		Password:          b.Password,
		AuthenticationUrl: b.AuthenticationUrl,
		LoginUrl:          b.LoginUrl,
		RestTokenTTL:      b.RestTokenTtl,
	}
	rr, accessTokenResponse, err := getNewAccessAndRefreshToken(b.B, params, b.RefreshToken)
	if err != nil {
		rr, tokenResponse, err := authenticate(b.B, params)
		if err != nil {
			fmt.Printf("failed to authenticate user while updating tokens with error %s\n", err)
			fmt.Printf("raw response -> %v", rr)
		}
		b.AuthorizationCode = tokenResponse.AuthorizationCode
		b.AccessToken = tokenResponse.AccessToken
		b.RefreshToken = tokenResponse.RefreshToken
		b.RestToken = tokenResponse.RestToken
		b.ApiUrl = tokenResponse.ApiUrl
		return nil
	}
	rr, restApiResponse, err := getRestToken(b.B, params, accessTokenResponse.AccessToken)
	if err != nil {
		fmt.Printf("failed to get rest token while updating tokens with error %s\n", err)
		fmt.Printf("raw response -> %v", rr)
		return err
	}
	b.RestToken = restApiResponse.RestToken
	b.ApiUrl = restApiResponse.ApiUrl
	return nil
}

func (b *bullhornClient) checkAndUpdateTokens() error {
	now := time.Now().UnixMilli()
	if b.RestTokenExpiryTime <= now {
		return b.updateTokensForClient()
	}
	return nil
}

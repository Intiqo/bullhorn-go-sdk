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
	case CandidateCertificationEntity:
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
	case FileAttachmentEntity:
		return nil
	case FileEntity:
		return nil
	case CertificationEntity:
		return nil
	}
	return fmt.Errorf("unsupported entity %s", name)
}

func (b *bullhornClient) validateEventTypes(et string) error {
	switch et {
	case "INSERTED":
		fallthrough
	case "UPDATED":
		fallthrough
	case "DELETED":
		return nil
	default:
		return fmt.Errorf("unsupported event type %s", et)
	}
}

func (b *bullhornClient) ParseResponseForEntity(name string, data interface{}, associations []string, isArray bool) (interface{}, error) {
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
	case CandidateCertificationEntity:
		var candidateCertification CandidateCertification
		var candidateCertifications []CandidateCertification
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &candidateCertifications)
			if err != nil {
				return nil, err
			}
			return candidateCertifications, nil
		} else {
			err = b.B.ParseResponse(data, &candidateCertification)
			if err != nil {
				return nil, err
			}
			return candidateCertification, nil
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
	case FileAttachmentEntity:
		var fileAttachment FileAttachment
		var fileAttachments []FileAttachment
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &fileAttachments)
			if err != nil {
				return nil, err
			}
			return fileAttachments, nil
		} else {
			err = b.B.ParseResponse(data, &fileAttachment)
			if err != nil {
				return nil, err
			}
			return fileAttachment, nil
		}
	case FileEntity:
		var file File
		var files []File
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &files)
			if err != nil {
				return nil, err
			}
			return files, nil
		} else {
			err = b.B.ParseResponse(data, &file)
			if err != nil {
				return nil, err
			}
			return file, nil
		}
	case CertificationEntity:
		var certification Certification
		var certifications []Certification
		var err error
		if isArray {
			err = b.B.ParseResponse(data, &certifications)
			if err != nil {
				return nil, err
			}
			return certifications, nil
		} else {
			err = b.B.ParseResponse(data, &certification)
			if err != nil {
				return nil, err
			}
			return certification, nil
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
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "get", b.getHeaders(), params, nil)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	dataMap := cr.Data.(map[string]interface{})
	return rr, dataMap["data"], nil
}

func (b *bullhornClient) GetMultipleEntities(name string, ids []int, options QueryOptions) (*resty.Response, interface{}, error) {
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

	idsString := make([]string, len(ids))
	for i, id := range ids {
		idsString[i] = strconv.Itoa(id)
	}

	requestUrl := fmt.Sprintf("%s/entity/%s/%s", b.ApiUrl, name, strings.Join(idsString, ","))
	if len(options.Associations) > 0 {
		requestUrl = fmt.Sprintf("%s/entity/%s/%s/%s", b.ApiUrl, name, strings.Join(idsString, ","), strings.Join(options.Associations[:], ","))
	}
	rr, cr, err := b.B.Call(requestUrl, "get", b.getHeaders(), params, nil)
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "get", b.getHeaders(), params, nil)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	dataMap := cr.Data.(map[string]interface{})
	return rr, dataMap["data"], nil
}

func (b *bullhornClient) GetAllEntities(name string, options QueryOptions) (*resty.Response, interface{}, error) {
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

	params := make(map[string]string)
	params["fields"] = strings.Join(options.Fields[:], ",")
	params["count"] = strconv.Itoa(count)
	params["start"] = strconv.Itoa(start)

	requestUrl := fmt.Sprintf("%s/my%ss", b.ApiUrl, name)
	rr, cr, err := b.B.Call(requestUrl, "get", b.getHeaders(), params, nil)
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "get", b.getHeaders(), params, nil)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	dataMap := cr.Data.(map[string]interface{})
	return rr, dataMap["data"], nil
}

func (b *bullhornClient) GetAssociations(entity string, association string, in GetAssociationInput) (*resty.Response, interface{}, error) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, nil, err
	}
	err = b.validateEntity(entity)
	if err != nil {
		return nil, nil, err
	}

	if in.ShowTotalMatched == nil {
		in.ShowTotalMatched = new(bool)
		*in.ShowTotalMatched = true
	}

	requestUrl := fmt.Sprintf("%s/association/%s/%s", b.ApiUrl, entity, association)
	rr, cr, err := b.B.Call(requestUrl, "post", b.getHeaders(), nil, in)
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "post", b.getHeaders(), nil, in)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	dataMap := cr.Data.(map[string]interface{})
	return rr, dataMap["data"], nil
}

func (b *bullhornClient) QueryEntity(name string, query string, options QueryOptions) (*resty.Response, interface{}, error) {
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
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "post", b.getHeaders(), nil, body)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	dataMap := cr.Data.(map[string]interface{})
	return rr, dataMap["data"], nil
}

func (b *bullhornClient) SearchEntity(name string, query string, options QueryOptions) (*resty.Response, interface{}, error) {
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
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "post", b.getHeaders(), nil, body)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	dataMap := cr.Data.(map[string]interface{})
	return rr, dataMap["data"], nil
}

func (b *bullhornClient) CreateEntity(name string, data map[string]interface{}) (*resty.Response, *CreateResponse, error) {
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
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "put", b.getHeaders(), nil, data)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	var createResponse CreateResponse
	err = b.B.ParseResponse(cr.Data, &createResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &createResponse, nil
}

func (b *bullhornClient) AssociateEntities(name string, id int, association string, associationIds []string) (*resty.Response, *CreateResponse, error) {
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
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "put", b.getHeaders(), nil, nil)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	var createResponse CreateResponse
	err = b.B.ParseResponse(cr.Data, &createResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &createResponse, nil
}

func (b *bullhornClient) UpdateEntity(name string, id int, data map[string]interface{}) (*resty.Response, *UpdateResponse, error) {
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
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "post", b.getHeaders(), nil, data)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
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
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, err
		}
		rr, _, err = b.B.Call(requestUrl, "delete", b.getHeaders(), nil, nil)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, errors.New(string(rr.Body()))
		}
		return rr, err
	}
	return rr, nil
}

func (b *bullhornClient) SubscribeToEvents(subscriptionId string, entities []string, eventTypes []string) (*resty.Response, *SubscribeEventResponse, error) {
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
		err = b.validateEventTypes(et)
		if err != nil {
			return nil, nil, err
		}
	}
	params := make(map[string]string)
	params["type"] = "entity"
	params["names"] = strings.Join(entities, ",")
	params["eventTypes"] = strings.Join(eventTypes, ",")
	requestUrl := fmt.Sprintf("%s/event/subscription/%s", b.ApiUrl, subscriptionId)
	rr, cr, err := b.B.Call(requestUrl, "put", b.getHeaders(), params, nil)
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "put", b.getHeaders(), params, nil)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	var subscribeResponse SubscribeEventResponse
	err = b.B.ParseResponse(cr.Data, &subscribeResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &subscribeResponse, nil
}

func (b *bullhornClient) UnsubscribeFromEvents(subscriptionId string, eventTypes []string) (*resty.Response, *UnsubscribeEventResponse, error) {
	var err error
	for _, et := range eventTypes {
		err = b.validateEventTypes(et)
		if err != nil {
			return nil, nil, err
		}
	}
	params := make(map[string]string)
	params["eventTypes"] = strings.Join(eventTypes, ",")
	requestUrl := fmt.Sprintf("%s/event/subscription/%s", b.ApiUrl, subscriptionId)
	rr, cr, err := b.B.Call(requestUrl, "delete", b.getHeaders(), params, nil)
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "delete", b.getHeaders(), params, nil)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	var unsubscribeEventResponse UnsubscribeEventResponse
	err = b.B.ParseResponse(cr.Data, &unsubscribeEventResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &unsubscribeEventResponse, nil
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
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "get", b.getHeaders(), params, nil)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	var fetchEventResponse FetchEventResponse
	err = b.B.ParseResponse(cr.Data, &fetchEventResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &fetchEventResponse, nil
}

func (b *bullhornClient) GetAttachmentsForEntity(entity string, entityId int, options QueryOptions) (*resty.Response, interface{}, error) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, nil, err
	}
	err = b.validateEntity(entity)
	if err != nil {
		return nil, nil, err
	}
	params := make(map[string]string)
	params["fields"] = strings.Join(options.Fields[:], ",")

	requestUrl := fmt.Sprintf("%s/entity/%s/%d/fileAttachments", b.ApiUrl, entity, entityId)
	rr, cr, err := b.B.Call(requestUrl, "get", b.getHeaders(), params, nil)
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "get", b.getHeaders(), params, nil)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	dataMap := cr.Data.(map[string]interface{})
	return rr, dataMap["data"], nil
}

func (b *bullhornClient) GetFileForEntity(entity string, entityId int, fileId int) (*resty.Response, interface{}, error) {
	err := b.checkAndUpdateTokens()
	if err != nil {
		return nil, nil, err
	}

	requestUrl := fmt.Sprintf("%s/file/%s/%d/%d", b.ApiUrl, entity, entityId, fileId)
	rr, cr, err := b.B.Call(requestUrl, "get", b.getHeaders(), nil, nil)
	if rr.RawResponse.StatusCode == 401 {
		err = b.updateTokensForClient()
		if err != nil {
			return nil, nil, err
		}
		rr, cr, err = b.B.Call(requestUrl, "get", b.getHeaders(), nil, nil)
	}
	if err != nil {
		var bhErr Error
		jErr := json.Unmarshal(rr.Body(), &bhErr)
		if jErr != nil {
			return rr, nil, errors.New(string(rr.Body()))
		}
		return rr, nil, err
	}
	data := cr.Data.(map[string]interface{})
	return rr, data, nil
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
	_, accessTokenResponse, err := getNewAccessAndRefreshToken(b.B, params, b.RefreshToken)
	if err != nil {
		rr, tokenResponse, err := authenticate(b.B, params)
		if err != nil {
			fmt.Printf("failed to authenticate user while updating tokens with error %s\n", err)
			fmt.Printf("raw response -> %v", rr)
			return err
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

package bullhorn

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type bullhornClient struct {
	B Backend

	AuthenticationUrl string
	LoginUrl          string
	ApiUrl            string

	AuthorizationCode string
	AccessToken       string
	RefreshToken      string
	ApiToken          string
}

func (b bullhornClient) getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["BhRestToken"] = b.ApiToken
	return headers
}

func (b bullhornClient) validateEntity(name string) error {
	switch name {
	case "Candidate":
		return nil
	case "ClientCorporation":
		return nil
	case "ClientContact":
		return nil
	case "Location":
		return nil
	case "JobOrder":
		return nil
	case "JobSubmission":
		return nil
	case "Placement":
		return nil
	}
	return fmt.Errorf("unsupported entity %s", name)
}

func (b bullhornClient) parseResponseForEntity(name string, data interface{}) (interface{}, error) {
	switch name {
	case "Candidate":
		var candidate Candidate
		err := b.B.ParseResponse(data, &candidate)
		if err != nil {
			return nil, err
		}
		return candidate, nil
	case "ClientCorporation":
		var clientCorporation ClientCorporation
		err := b.B.ParseResponse(data, &clientCorporation)
		if err != nil {
			return nil, err
		}
		return clientCorporation, nil
	case "ClientContact":
		var clientContact ClientContact
		err := b.B.ParseResponse(data, &clientContact)
		if err != nil {
			return nil, err
		}
		return clientContact, nil
	case "Location":
		var location Location
		err := b.B.ParseResponse(data, &location)
		if err != nil {
			return nil, err
		}
		return location, nil
	case "JobOrder":
		var jobOrder JobOrder
		err := b.B.ParseResponse(data, &jobOrder)
		if err != nil {
			return nil, err
		}
		return jobOrder, nil
	case "JobSubmission":
		var jobSubmission JobSubmission
		err := b.B.ParseResponse(data, &jobSubmission)
		if err != nil {
			return nil, err
		}
		return jobSubmission, nil
	case "Placement":
		var placement Placement
		err := b.B.ParseResponse(data, &placement)
		if err != nil {
			return nil, err
		}
		return placement, nil
	}
	return nil, fmt.Errorf("unsupported entity %s", name)
}

func (b bullhornClient) Ping() (int64, error) {
	requestUrl := fmt.Sprintf("%s/ping", b.ApiUrl)
	_, cr, err := b.B.Call(requestUrl, "get", b.getHeaders(), nil, nil)
	if err != nil {
		return 0, err
	}
	if cr.Code >= 400 {
		return 0, err
	}
	var pingResponse Ping
	err = b.B.ParseResponse(cr.Data, &pingResponse)
	if err != nil {
		return 0, err
	}
	return pingResponse.SessionExpires, nil
}

func (b bullhornClient) GetEntity(name string, id int, options QueryOptions) (*resty.Response, interface{}, error) {
	err := b.validateEntity(name)
	if err != nil {
		return nil, nil, err
	}
	params := make(map[string]string)
	params["fields"] = strings.Join(options.Fields[:], ",")

	requestUrl := fmt.Sprintf("%s/entity/%s/%d", b.ApiUrl, name, id)
	rr, cr, err := b.B.Call(requestUrl, "get", b.getHeaders(), params, nil)
	if err != nil {
		return rr, nil, err
	}
	dataMap := cr.Data.(map[string]interface{})
	responseData, err := b.parseResponseForEntity(name, dataMap["data"])
	if err != nil {
		return rr, nil, err
	}
	return rr, responseData, nil
}

func (b bullhornClient) CreateEntity(name string, data map[string]interface{}) (*resty.Response, *CreateResponse, error) {
	err := b.validateEntity(name)
	if err != nil {
		return nil, nil, err
	}
	requestUrl := fmt.Sprintf("%s/entity/%s", b.ApiUrl, name)
	rr, cr, err := b.B.Call(requestUrl, "put", b.getHeaders(), nil, data)
	if err != nil {
		return rr, nil, err
	}
	var createResponse CreateResponse
	err = b.B.ParseResponse(cr.Data, &createResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &createResponse, nil
}

func (b bullhornClient) UpdateEntity(name string, id int, data map[string]interface{}) (*resty.Response, *UpdateResponse, error) {
	err := b.validateEntity(name)
	if err != nil {
		return nil, nil, err
	}
	requestUrl := fmt.Sprintf("%s/entity/%s/%d", b.ApiUrl, name, id)
	rr, cr, err := b.B.Call(requestUrl, "post", b.getHeaders(), nil, data)
	if err != nil {
		return rr, nil, err
	}
	var updateResponse UpdateResponse
	err = b.B.ParseResponse(cr.Data, &updateResponse)
	if err != nil {
		return rr, nil, err
	}
	return rr, &updateResponse, nil
}

func (b bullhornClient) DeleteEntity(name string, id int) (*resty.Response, error) {
	err := b.validateEntity(name)
	if err != nil {
		return nil, err
	}
	requestUrl := fmt.Sprintf("%s/entity/%s/%d", b.ApiUrl, name, id)
	rr, _, err := b.B.Call(requestUrl, "post", b.getHeaders(), nil, nil)
	if err != nil {
		return rr, err
	}
	return rr, nil
}

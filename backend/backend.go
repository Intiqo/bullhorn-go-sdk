package backend

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

var backend Service

func init() {
	r := resty.New()
	r.SetRedirectPolicy(resty.FlexibleRedirectPolicy(0))
	backend = &ServiceImplementation{
		HTTPClient: r,
	}
}

func GetBackendService() Service {
	return backend
}

type ClientResponse struct {
	Code int
	Data interface{}
}

type Service interface {
	Call(url string, method string, headers map[string]string, query map[string]string, body interface{}) (*resty.Response, *ClientResponse, error)
	ParseResponse(data interface{}, response interface{}) error
}

type ServiceImplementation struct {
	HTTPClient *resty.Client
}

func (b ServiceImplementation) Call(url string, method string, headers map[string]string, query map[string]string, body interface{}) (*resty.Response, *ClientResponse, error) {
	var response interface{}
	req := b.HTTPClient.R().
		SetHeaders(headers).
		SetQueryParams(query).
		SetBody(body).
		SetResult(&response)

	var resp *resty.Response
	var err error

	switch method {
	case "get":
		resp, err = get(url, req)
	case "post":
		resp, err = post(url, req)
	case "put":
		resp, err = put(url, req)
	case "patch":
		resp, err = patch(url, req)
	case "delete":
		resp, err = rdelete(url, req)
	}

	if err != nil {
		return resp, nil, err
	}

	if resp.RawResponse.StatusCode >= 400 {
		return resp, nil, fmt.Errorf("request failed with code %s. see raw response for details", resp.Status())
	}

	return resp, &ClientResponse{
		Code: resp.RawResponse.StatusCode,
		Data: response,
	}, err
}

func (b ServiceImplementation) ParseResponse(data interface{}, response interface{}) error {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonString, response)
	if err != nil {
		return nil
	}
	return nil
}

func get(url string, r *resty.Request) (*resty.Response, error) {
	return r.Get(url)
}

func post(url string, r *resty.Request) (*resty.Response, error) {
	return r.Post(url)
}

func put(url string, r *resty.Request) (*resty.Response, error) {
	return r.Put(url)
}

func rdelete(url string, r *resty.Request) (*resty.Response, error) {
	return r.Delete(url)
}

func patch(url string, r *resty.Request) (*resty.Response, error) {
	return r.Patch(url)
}

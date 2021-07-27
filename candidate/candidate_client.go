package candidate

import (
	"fmt"
	"strings"

	"github.com/adwitiyaio/bullhorn"
	"github.com/adwitiyaio/bullhorn/backend"
)

type Options struct {
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

type Entity interface {
	Get(name string, entityIdList []int, options Options) (interface{}, error)
	Create(name string, entityId int, data interface{}) (interface{}, error)
	Update(name string, entityId int, data interface{}) (interface{}, error)
}

type Client struct {
	B        backend.Service
	ApiToken string
	ApiUrl   string
}

func getClient() Client {
	return Client{B: backend.GetBackendService(), ApiToken: bullhorn.ApiToken, ApiUrl: bullhorn.RestApiUrl}
}

// Get Gets a candidate record
func Get(entityId int, options Options) (*bullhorn.Candidate, error) {
	return getClient().Get(entityId, options)
}

// Get Gets a candidate record
func (c Client) Get(entityId int, options Options) (*bullhorn.Candidate, error) {
	headers := make(map[string]string)
	headers["BhRestToken"] = c.ApiToken

	params := make(map[string]string)
	params["fields"] = strings.Join(options.Fields[:], ",")

	requestUrl := fmt.Sprintf("%s/entity/Candidate/%d", c.ApiUrl, entityId)
	_, cr, err := c.B.Call(requestUrl, "get", headers, params, nil)
	if err != nil {
		return nil, err
	}
	if cr.Code >= 400 {
		return nil, err
	}
	dataMap := cr.Data.(map[string]interface{})
	var candidate bullhorn.Candidate
	err = c.B.ParseResponse(dataMap["data"], &candidate)
	if err != nil {
		return nil, err
	}
	return &candidate, nil
}

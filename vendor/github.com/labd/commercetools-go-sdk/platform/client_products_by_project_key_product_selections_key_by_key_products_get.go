package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput
}

func (r *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput struct {
	Expand    []string
	Limit     *int
	Offset    *int
	WithTotal *bool
}

func (input *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Expand(v []string) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Limit(v int) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Offset(v int) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) WithTotal(v bool) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) WithQueryParams(input ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Execute(ctx context.Context) (result *ProductSelectionProductPagedQueryResponse, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj

	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
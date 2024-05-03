package kagi

import (
	"fmt"
	"github.com/httpjamesm/kagigo/types"
)

const (
	EndpointTypeWeb  string = "web"
	EndpointTypeNews string = "news"
)

type EnrichmentParams struct {
	Q string `json:"q"`
}

type EnrichmentResponse struct {
	Meta struct {
		ID   string `json:"id"`
		Node string `json:"node"`
		Ms   int    `json:"ms"`
		API  string `json:"api_balance"`
	} `json:"meta"`
	Data []struct {
		T         int      `json:"t"`
		Rank      int      `json:"rank"`
		URL       string   `json:"url"`
		Title     string   `json:"title"`
		Snippet   string   `json:"snippet"`
		Published string   `json:"published"`
		List      []string `json:"list"`
	} `json:"data"`
	Errors []types.Error `json:"error"`
}

func (c *Client) EnrichmentCompletion(endpointType string, params EnrichmentParams) (res EnrichmentResponse, err error) {
	if params.Q == "" {
		err = fmt.Errorf("query is required")
		return
	}

	if endpointType == "" {
		err = fmt.Errorf("endpoint_type is required")
		return
	}

	if endpointType != EndpointTypeWeb && endpointType != EndpointTypeNews {
		err = fmt.Errorf("endpoint_type must be web or news")
		return
	}

	err = c.SendRequest("GET", "/enrich/"+endpointType, params, &res)
	if err != nil {
		return
	}

	// will this ever happen here?
	if len(res.Errors) != 0 {
		errObj := res.Errors[0]
		err = fmt.Errorf("api returned error: %v", fmt.Sprintf("[code: %d, msg: %s, ref: %v]", errObj.Code, errObj.Msg, errObj.Ref))
		return
	}

	return
}

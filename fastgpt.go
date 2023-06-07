package kagi

import (
	"fmt"
)

type FastGPTCompletionParams struct {
	Query     string `json:"query"`
	WebSearch bool   `json:"web_search"`
	Cache     bool   `json:"cache"`
}

type FastGPTCompletionResponse struct {
	Meta struct {
		ID   string `json:"id"`
		Node string `json:"node"`
		Ms   int    `json:"ms"`
	} `json:"meta"`
	Data struct {
		Output string `json:"output"`
		Tokens int    `json:"tokens"`
	} `json:"data"`
}

func (c *Client) FastGPTCompletion(params FastGPTCompletionParams) (res FastGPTCompletionResponse, err error) {
	if params.Query == "" {
		err = fmt.Errorf("query is required")
		return
	}

	err = c.SendRequest("POST", "/fastgpt", params, &res)
	if err != nil {
		return
	}

	return
}

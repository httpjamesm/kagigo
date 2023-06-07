package kagi

import "fmt"

type SummaryType string

const (
	SummaryTypeSummary   SummaryType = "summary"
	SummaryTypeTakeaways SummaryType = "takeaway"
)

type SummaryEngine string

const (
	SummaryEngineT5   SummaryEngine = "cecil"
	SummaryEngineGPT3 SummaryEngine = "agnes"
	SummaryEngineGPT2 SummaryEngine = "daphne"
	SummaryEngineGPT1 SummaryEngine = "muriel"
)

type UniversalSummarizerParams struct {
	URL         string        `json:"url"`
	SummaryType SummaryType   `json:"summary_type"`
	Engine      SummaryEngine `json:"engine"`
}

type UniversalSummarizerResponse struct {
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

func (c *Client) UniversalSummarizerCompletion(params UniversalSummarizerParams) (res UniversalSummarizerResponse, err error) {
	if params.URL == "" {
		err = fmt.Errorf("url is required")
		return
	}

	err = c.SendRequest("POST", "/summarize", params, &res)
	if err != nil {
		return
	}

	return
}

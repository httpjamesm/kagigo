# kagigo

An unofficial [Kagi API](https://help.kagi.com/kagi/api/overview.html) client for Go.

## Installation

```bash
go get -u github.com/httpjamesm/kagigo
```

## Quick Start

### Client

```go
client := kagi.NewClient(&kagi.ClientConfig{
    APIKey:     os.Getenv("KAGI_API_KEY"),
    APIVersion: "v0",
})
```
### FastGPT

```go
response, err := client.FastGPTCompletion(kagi.FastGPTCompletionParams{
    Query:     "query",
    WebSearch: true,
    Cache:     true,
})
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(response.Data.Output)
```

### Universal Summarizer

```go
response, err := client.UniversalSummarizerCompletion(kagi.UniversalSummarizerParams{
    URL:         "https://blog.kagi.com/security-audit",
    SummaryType: kagi.SummaryTypeSummary,
    Engine:      kagi.SummaryEngineCecil,
})
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(response.Data.Output)
```

### Enrichment

```go
response, err := client.EnrichmentCompletion(kagi.EndpointTypeWeb, kagi.EnrichmentParams{
    Q: "kagi search",
})
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(response.Data)
```

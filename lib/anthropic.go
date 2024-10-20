package lib

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "cillers-cli/config"
)

const (
    AnthropicAPIURL     = "https://api.anthropic.com/v1/messages"
    AnthropicAPIVersion = "2023-06-01"
    AnthropicModel      = "claude-3-5-sonnet-20240620"
)

type AnthropicRequest struct {
    Model     string    `json:"model"`
    MaxTokens int       `json:"max_tokens"`
    Messages  []Message `json:"messages"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type AnthropicResponse struct {
    Content []struct {
        Text string `json:"text"`
    } `json:"content"`
}

func SendPromptToAnthropic(prompt string) (string, error) {
    var conf config.Config = config.LoadConfig()
    apiKey := conf.AnthropicAPIKey
    if apiKey == "" {
        return "", fmt.Errorf("ANTHROPIC_API_KEY environment variable is not set")
    }
    request := AnthropicRequest{
        Model:     AnthropicModel,
        MaxTokens: 8192,
        Messages: []Message{
            {Role: "user", Content: prompt},
        },
    }

    jsonData, err := json.Marshal(request)
    if err != nil {
        return "", fmt.Errorf("error marshaling request: %w", err)
    }

    req, err := http.NewRequest("POST", AnthropicAPIURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", fmt.Errorf("error creating request: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("x-api-key", apiKey)
    req.Header.Set("anthropic-version", AnthropicAPIVersion)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("error sending request: %w", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error reading response body: %w", err)
    }

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
    }

    var anthropicResp AnthropicResponse
    err = json.Unmarshal(body, &anthropicResp)
    if err != nil {
        return "", fmt.Errorf("error unmarshaling response: %w", err)
    }

    if len(anthropicResp.Content) == 0 {
        return "", fmt.Errorf("empty response from Anthropic API")
    }

    return anthropicResp.Content[0].Text, nil
}

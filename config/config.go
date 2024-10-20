package config

import (
    "gopkg.in/yaml.v3"
    "io/ioutil"
)

type Config struct {
    Version string
    TemplateRepoURL string
    DocumentationURL string
    AnthropicAPIKey string
    AnthropicAPIURL string
    AnthropicModel string
}

type Secrets struct {
    Anthropic struct {
        APIKey string `yaml:"api_key"`
    } `yaml:"anthropic"`
}

func readAnthropicAPIKey() string {
    data, err := ioutil.ReadFile(".cillers/secrets_and_local_config/secrets.yml")
    if err != nil {
        return ""
    }

    var secrets Secrets
    err = yaml.Unmarshal(data, &secrets)
    if err != nil {
        return ""
    }

    return secrets.Anthropic.APIKey
}

func LoadConfig() Config {
    return Config{
        Version: "v0.1.2",
        TemplateRepoURL: "https://github.com/Cillers-com/create-cillers-system",
        DocumentationURL: "https://github.com/Cillers-com/cillers-cli",
        AnthropicAPIKey: readAnthropicAPIKey(),
        AnthropicAPIURL: "https://api.anthropic.com/v1/messages",
        AnthropicModel: "claude-3-5-sonnet-20240620",
    }
}

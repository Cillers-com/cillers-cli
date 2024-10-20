package config
import "os"

type Config struct {
    Version string
    TemplateRepoURL string
    DocumentationURL string
    AnthropicAPIKey string
    AnthropicAPIURL string
    AnthropicModel string
}

func LoadConfig() Config {
    return Config{
        Version: "v0.1.2",
        TemplateRepoURL: "https://github.com/Cillers-com/create-cillers-system",
        DocumentationURL: "https://github.com/Cillers-com/cillers-cli",
        AnthropicAPIKey: os.Getenv("ANTHROPIC_API_KEY"),
        AnthropicAPIURL: "https://api.anthropic.com/v1/messages",
        AnthropicModel: "claude-3-5-sonnet-20240620",
    }
}

package config

type Config struct {
    Version         string
    TemplateRepoURL string
    DocumentationURL string
}

func LoadConfig() Config {
    return Config{
        Version:         "v0.1.1",
        TemplateRepoURL: "https://github.com/Cillers-com/create-cillers-system",
        DocumentationURL: "https://github.com/Cillers-com/cillers-cli",
    }
}

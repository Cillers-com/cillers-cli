package config

type Config struct {
    Version         string
    TemplateRepoURL string
    DocumentationURL string
}

func Get() Config {
    return Config{
        Version:         "v0.1.0",
        TemplateRepoURL: "https://github.com/Cillers-com/create-cillers-system",
        DocumentationURL: "https://github.com/Cillers-com/cillers-cli",
    }
}

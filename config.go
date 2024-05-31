package emailist

import "os"

type Config struct {
	APIKey string
}

func NewConfig() *Config {
	key := os.Getenv("MAIL_API_KEY")
	return &Config{APIKey: key}
}

func NewConfigWithAPIKey(key string) *Config {
	return &Config{APIKey: key}
}

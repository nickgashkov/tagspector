package config

type Config struct {
	pattern string
}

func New(args []string, environ []string) *Config {
	return &Config{}
}

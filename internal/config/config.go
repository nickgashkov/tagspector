package config

import (
	"strings"
)

type Config struct {
	pattern string
}

type Option struct {
	key string
	help string
}

func newOption(key, help string) *Option {
	return &Option{key, help}
}

func (o Option) getValue(args, environ map[string]string) (string, bool) {
	argsKey := o.key
	environKey := getEnvironKeyFromArgsKey(o.key)

	if value, ok := args[argsKey]; ok {
		return value, true
	}

	if value, ok := environ[environKey]; ok {
		return value, true
	}

	return "", false
}

func getEnvironKeyFromArgsKey(argsKey string) string {
	environKey := strings.TrimLeft(argsKey, "-")
	environKey = strings.ToUpper(environKey)
	environKey = "TAGSPECTOR" + "_" + environKey

	return environKey
}

func New(args []string, environ []string) *Config {
	argsMap := parseArgs(args)
	environMap := parseEnviron(environ)

	patternOption := newOption("--pattern", "A pattern to search the codetags from.")
	pattern, ok := patternOption.getValue(argsMap, environMap)

	if ok {
		return &Config{pattern}
	}

	return &Config{}
}

func parseArgs(args []string) map[string]string {
	var lastArg string
	parsed := make(map[string]string)

	for _, arg := range args {
		if strings.HasPrefix (arg, "--") {
			lastArg = arg
		} else {
			parsed[lastArg] = arg
		}
	}

	return parsed
}

func parseEnviron(environ []string) map[string]string {
	parsed := make(map[string]string)

	for _, item := range environ {
		key, value := parseEnvironItem(item)
		parsed[key] = value
	}

	return parsed
}

func parseEnvironItem(item string) (string, string){
	s := strings.Split(item, "=")
	return s[0], s[1]
}

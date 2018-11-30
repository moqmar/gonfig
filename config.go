package gonfig

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Open loads a configuration file into the memory and returns a Config object
func Open(file ...string) *Config {
	// TODO: multiple files and YAML strings
	f, err := ioutil.ReadFile(file[0])
	if err != nil {
		panic(err)
	}

	data := map[interface{}]interface{}{} // TODO: anything?!
	yaml.Unmarshal(f, data)
	fmt.Printf("%+v\n", data)
	cfg := &Config{}
	apply(data, cfg)
	fmt.Printf("%+v\n", cfg)
	return cfg
}

// OpenDynamic loads a configuration file into the memory and returns a Config object that changes when the file changes and supports OnChange handlers.
func OpenDynamic(file ...string) *Config {
	// TODO:
	return nil
}

// Env uses environment variables as the default configuration.
// You can specify a blacklist with paths that won't be read from the environment, or use ...config.SystemBlacklist to ignore common system variables.
// Usage: config.Open(..., config.Env())
// Example: EXAMPLE_TEST=1 => cfg.Get("example.test").Int() == 1
func Env(blacklist ...string) string {
	// TODO:
	return "---\n{}"
}

// Arg uses command line arguments as the default configuration.
// You can specify a blacklist with paths that won't be loaded from the command line arguments.
// Usage: config.Open(..., config.Arg())
// Example: --example-test 1 => cfg.Get("example.test").Int() == 1
func Arg(blacklist ...string) string {
	// TODO:s
	return "---\n{}"
}

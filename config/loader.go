package config

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Github   Github
	Output   Output
	Strategy Strategy
}

func LoadFile(file string) (Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return Config{}, err
	}
	return Load(b)
}

func Load(body []byte) (Config, error) {
	cfg := Config{}
	err := yaml.Unmarshal(body, &cfg)
	return cfg, err
}

package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Ip string       `yaml:"ip"`
	Port string     `yaml:"port"`
	User string     `yaml:"user"`
	Password string `yaml:"password"`
	Dbname string   `yaml:"dbname"`
}

func Load(filename string) ([]byte, error) {
	ymlData, err := ioutil.ReadFile(filename)
	return []byte(ymlData), err
}

func Parse(data []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(data, &config)
	return config, err
}

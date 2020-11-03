package config

import (
	"lottery_back/pkg/util"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Server struct {
		Host []string `yaml:"host"`
	} `yaml:"server"`
	Api struct {
		Key string `yaml:"key"`
	} `yaml:"api"`
	ResourcePath struct {
		Applicant string `yaml:"applicant"`
		Prize     string `yaml:"prize"`
	} `yaml:"resource_path"`
}

func LoadConfig() Config {
	f, err := os.Open("conf/config.yaml")
	util.CheckFatalError(err)
	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	util.CheckFatalError(err)

	return cfg
}

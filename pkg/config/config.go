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

var ConfigData Config

func LoadConfig() {
	f, err := os.Open("conf/config.yaml")
	util.CheckFatalError(err)
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&ConfigData)
	util.CheckFatalError(err)
}

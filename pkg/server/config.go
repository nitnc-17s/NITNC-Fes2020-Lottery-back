package server

import (
	"lottery_back/pkg/util"
	"os"

	"github.com/go-yaml/yaml"
)

type config struct {
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

func loadConfig() *config {
	f, err := os.Open("conf/Config.yaml")
	util.CheckFatalError(err)
	defer f.Close()

	var cfg config
	err = yaml.NewDecoder(f).Decode(&cfg)
	util.CheckFatalError(err)

	return &cfg
}

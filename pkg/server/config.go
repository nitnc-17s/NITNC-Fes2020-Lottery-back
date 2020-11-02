package server

import (
	"log"
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
}

func loadConfig() *config {
	f, err := os.Open("conf/config.yaml")
	if err != nil {
		log.Fatal("Config Open error:", err)
		return nil
	}
	defer f.Close()

	var cfg config
	err = yaml.NewDecoder(f).Decode(&cfg)

	if err != nil {
		log.Fatal("Config Parse error:", err)
		return nil
	}

	return &cfg
}

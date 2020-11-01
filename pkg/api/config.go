package api

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type config struct {
	Server server `yaml:"server"`
}

type server struct {
	Host string `yaml:"host"`
}

func loadConfig() *config {
	f, err := os.Open("conf/config.yml")
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

package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	Local struct {
		LocalServer   string
		Dbms          string
		Db            string
		User          string
		Password      string
		Host          string
		RedisHost     string
		RedisPassword string
		JwtSecretKey  string
	}
}

func InitToml(path string) Config {
	config := new(Config)

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed toml setting: %v", err)
	}
	defer file.Close()

	if _, err := toml.NewDecoder(file).Decode(config); err != nil {
		log.Fatalf("failed toml setting: %v", err)
	}

	return *config
}

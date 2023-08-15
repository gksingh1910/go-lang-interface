package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	store  storeConfig
	Server HttpServerConfig
}

type storeConfig struct {
	storeUrl    string
	storePort   string
	storeSchema string
}

type HttpServerConfig struct {
	Host string
	Port string
}

var Conf *Config

func Init() {
	godotenv.Load()

	serverConfig := HttpServerConfig{
		Host: os.Getenv("HTTP_SERVER"),
		Port: os.Getenv("HTTP_PORT"),
	}

	Conf = &Config{
		Server: serverConfig,
	}

}

func Get() Config {
	fmt.Printf("\n http server host is %v", *Conf)
	return *Conf
}

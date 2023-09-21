package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerConfig HTTPServerConfig
	StoreConfig  DBConfig
}

type HTTPServerConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	DBUrl string
}

var cfg *Config

func Init() {
	godotenv.Load()

	HTTPServerConf := HTTPServerConfig{
		Host: os.Getenv("HTTPServerHost"),
		Port: os.Getenv("HTTPServerPort"),
	}

	cfg = &Config{
		ServerConfig: HTTPServerConf,
	}
}

func Get() *Config {
	return cfg
}

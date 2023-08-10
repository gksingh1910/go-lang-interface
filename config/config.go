package config

import (
	"fmt"
	"os"
)

type config struct {
	store  storeConfig
	server httpServerConfig
}

type storeConfig struct {
	storeUrl    string
	storePort   string
	storeSchema string
}

type httpServerConfig struct {
	host string
	port string
}

func init() {
	serverConfig := httpServerConfig{
		host: os.Getenv("HTTP_SERVER"),
		port: os.Getenv("HTTP_PORT"),
	}

	conf := config{
		server: serverConfig,
	}

	fmt.Printf(conf.server.host)

}

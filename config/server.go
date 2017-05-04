package config

import (
	"os"
)

var DEFAULT_PORT = 8080

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	return port
}

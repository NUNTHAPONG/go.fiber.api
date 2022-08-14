package env

import (
	"os"

	"github.com/joho/godotenv"
)

type NetworkEnv struct {
	HOST string
	PORT string
}

func Network() NetworkEnv {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env")
	}
	return NetworkEnv{
		HOST: os.Getenv("SERVER_HOST"),
		PORT: os.Getenv("SERVER_PORT"),
	}
}
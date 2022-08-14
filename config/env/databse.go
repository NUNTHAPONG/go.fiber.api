package env

import (
	"os"

	"github.com/joho/godotenv"
)

type DbEnv struct {
	USERNAME string
	PASSWORD string
	HOST string
	PORT string
	DATABASE string
	SCHEMA string
}

func Db() DbEnv {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env")
	}
	return DbEnv{
		USERNAME: os.Getenv("DB_USER"),
		PASSWORD: os.Getenv("DB_PASS"),
		HOST: os.Getenv("DB_HOST"),
		PORT: os.Getenv("DB_PORT"),
		DATABASE: os.Getenv("DB_NAME"),
		SCHEMA: os.Getenv("DB_PATH"),
	}
}
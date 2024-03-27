package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	APP_NAME    string
	SERVER_PORT string
	SERVER_MODE string
}

var ENV = Env{}

func ParseEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error parsing .env file")
		return err
	}

	ENV.APP_NAME = os.Getenv("APP_NAME")
	ENV.SERVER_PORT = os.Getenv("SERVER_PORT")
	ENV.SERVER_MODE = os.Getenv("SERVER_MODE")

	return nil
}

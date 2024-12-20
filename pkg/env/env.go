package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var POSTGRES_URL string
var POSTGRES_DB string
var POSTGRES_USER string
var POSTGRES_PASSWORD string
var POSTGRES_HOST string
var POSTGRES_PORT uint16

func LoadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	POSTGRES_URL = os.Getenv("POSTGRES_URL")
	POSTGRES_DB = os.Getenv("POSTGRES_DB")
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")

	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatal("Error parsing POSTGRES_PORT")
		panic(err)
	}
	POSTGRES_PORT = uint16(port)
}

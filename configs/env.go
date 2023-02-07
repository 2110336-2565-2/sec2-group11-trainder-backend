package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGO_URI")
}

func EnvMongoDBName() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGO_DATABASE_NAME")
}

func EnvTokenLifeSpan() (token_lifespan int, err error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	token_lifespan, err = strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	return token_lifespan, err
}

func EnvApiSecret() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("API_SECRET")
}

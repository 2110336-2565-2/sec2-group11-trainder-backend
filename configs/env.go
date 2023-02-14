package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func loadDotEnv() {
	env := os.Getenv("TRAINDER_DO_NOT_USE_DOTENV")
	// if flag not set load .env file
	if env == "" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func EnvMongoURI() string {
	loadDotEnv()
	return os.Getenv("MONGO_URI")
}

func EnvMongoDBName() string {
	loadDotEnv()
	return os.Getenv("MONGO_DATABASE_NAME")
}

func EnvTokenLifeSpan() (tokenLifespan int, err error) {
	loadDotEnv()
	tokenLifespan, err = strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	return tokenLifespan, err
}

func EnvApiSecret() string {
	loadDotEnv()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("API_SECRET")
}

func EnvInitRootPassword() string {
	loadDotEnv()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv("INIT_ROOT_PASSWORD")
	if env == "" {
		log.Fatal("Root password cannot be empty")
	}
	return env
}

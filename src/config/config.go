package config

import (
	"flag"
	"golang-backend-microservice/container/log"
	"golang-backend-microservice/container/utils"
	"os"

	"github.com/joho/godotenv"
)

const RETRY_TIMER = 30

func Init() {
	setupEnvironmentVariables()
	initLogger()
}

func initLogger() {
	log.CreateTransports(log.Console, log.File, log.Rollbar)
}

func setupEnvironmentVariables() {
	envList := flag.String("env", ".env.development", "Env file")
	if err := godotenv.Load(*envList); err != nil {
		_, exists := os.LookupEnv("ENVIRONMENT")
		if exists {
			return
		}
		var envName string = ".env.development"
		newEnv := map[string]string{
			"ENVIRONMENT": utils.ENV_DEVELOPMENT,

			// NATS Server
			"NATS_HOST":    "localhost:4222",
			"NATS_USER":    "local",
			"NATS_PASS":    "password",
			"NATS_TIMEOUT": "1000",

			// MySQL Server
			"MYSQL_HOST": "db",
			"MYSQL_USER": "root",
			"MYSQL_PASS": "password",

			// MongoDB Server
			"MONGO_HOST": "mongodb://localhost:27017",
			"MONGO_USER": "root",
			"MONGO_PASS": "password",

			// Others
			"ROLLBAR_ACCESS_TOKEN": "",
		}
		godotenv.Write(newEnv, envName)
		godotenv.Load(envName)
	}
}

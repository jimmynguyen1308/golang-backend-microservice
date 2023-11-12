package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	setupEnvironmentVariables()
}

func setupEnvironmentVariables() {
	envList := flag.String("env", ".env.development", "Env file")
	if err := godotenv.Load(*envList); err != nil {
		env, exists := os.LookupEnv("ENVIRONMENT")
		if exists {
			log.Printf("Global ENVIRONMENT: %s\n", env)
			return
		}
		var envName string = ".env.development"
		log.Printf("Global ENVIRONMENT: %s\n", envName)
		newEnv := map[string]string{
			"ENVIRONMENT": "development",

			// NATS Server
			"NATS_HOST":    "localhost:4222",
			"NATS_USER":    "local",
			"NATS_PASS":    "password",
			"NATS_TIMEOUT": "1000",

			// MySQL Server
			"MYSQL_HOST": "localhost",
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

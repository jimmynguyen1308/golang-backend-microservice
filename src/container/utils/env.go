package utils

import (
	"os"
)

type Env = string

const (
	ENV_DEVELOPMENT Env = "development"
	ENV_TESTING     Env = "testing"
	ENV_STAGING     Env = "staging"
	ENV_PRODUCTION  Env = "production"
)

func IsEnv(envs ...Env) bool {
	for _, env := range envs {
		if os.Getenv("ENVIRONMENT") == env {
			return true
		}
	}
	return false
}

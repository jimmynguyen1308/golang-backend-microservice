package utils

import "os"

const (
	Development = "development"
	Testing     = "testing"
	Staging     = "staging"
	Production  = "production"
)

func IsEnv(envs ...string) bool {
	for _, env := range envs {
		if os.Getenv(env) != "" {
			return true
		}
	}
	return false
}

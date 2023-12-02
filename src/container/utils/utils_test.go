package utils_test

import (
	"golang-backend-microservice/container/utils"
	"testing"
)

func TestIsEnv(t *testing.T) {
	t.Setenv("ENVIRONMENT", utils.ENV_DEVELOPMENT)
	if !utils.IsEnv(utils.ENV_DEVELOPMENT) {
		t.Errorf("Expected IsEnv(%s) to be true, but got false", utils.ENV_DEVELOPMENT)
	}
	if utils.IsEnv(utils.ENV_PRODUCTION) {
		t.Errorf("Expected IsEnv(%s) to be false, but got true", utils.ENV_PRODUCTION)
	}
}

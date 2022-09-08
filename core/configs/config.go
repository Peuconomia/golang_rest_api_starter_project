package configs

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"golang_rest_api_starter_project/core/types"
	"os"
)

type AppConfig struct {
	GinMode string
}

func NewAppConfig() *AppConfig {
	// Configure and setup AppConfig
	return setupAppConfigurations()
}

var Module = fx.Provide(
	NewAppConfig,
)

func setupAppConfigurations() *AppConfig {
	var environment *string = flag.String("environment", types.None.String(),
		`environment flag is used to specify which environment files to use for configuration. 
If not present, none is default and no files are loaded. Otherwise, specific flags are used for specific files.

default = .env
staging = .env.staging
development = .env.development
production = .env.production
none = no files are loaded
`)
	// Parse all the flags
	flag.Parse()

	// Load .env files according to the environment_mode
	loadEnvironmentVariables(types.EnvModeFromString(environment))

	// create placeholder appConfig
	var appConfig *AppConfig = &AppConfig{}

	// region - Env variables setup
	setupEnvironmentVariables(appConfig)
	// endregion

	// return the AppConfig
	return appConfig
}

func loadEnvironmentVariables(envMode types.EnvMode) {
	switch envMode {
	case types.Default:
		if environmentFileExists(".env") {
			godotenv.Load(".env")
		} else {
			panic(".env file does not exist.")
		}
		godotenv.Load()

	case types.Staging:
		if environmentFileExists(".env.staging") {
			godotenv.Load(".env.staging")
		} else {
			panic(".env.staging file does not exist.")
		}

	case types.Development:
		if environmentFileExists(".env.development") {
			godotenv.Load(".env.development")
		} else {
			panic(".env.development file does not exist.")
		}

	case types.Production:
		if environmentFileExists(".env.production") {
			godotenv.Load(".env.production")
		} else {
			panic(".env.production file does not exist.")
		}

	default:
	}
}

func environmentFileExists(fileName string) bool {
	if _, err := os.Stat(fmt.Sprintf("./%s", fileName)); err == nil {
		// file exist
		return true
	}

	return false
}

func setupEnvironmentVariables(appConfig *AppConfig) {
	// Setup GIN_MODE
	if ginMode := os.Getenv(types.GinMode); len(ginMode) != 0 {
		appConfig.GinMode = ginMode
	} else {
		panic(fmt.Sprintf("%s cannot be empty or missing.", types.GinMode))
	}
}

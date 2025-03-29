package environment

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	"github.com/AJackTi/go-ecommerce-microservices-example/internal/pkg/constants"
)

type Environment string

var (
	Development = Environment(constants.Dev)
	Test        = Environment(constants.Test)
	Production  = Environment(constants.Production)
)

func ConfigAppEnv(environments ...Environment) Environment {
	environment := Environment("")
	if len(environments) > 0 {
		environment = environments[0]
	} else {
		environment = Development
	}

	// setup viper to read from os environment with `viper.Get`
	viper.AutomaticEnv()

	// https://articles.wesionary.team/environment-variable-configuration-in-your-golang-project-using-viper-4e8289ef664d
	// load environment variables from .env file to system environment variables, it just finds `.env` file in our current `executing working directory` in our app for example `catalogs/service`
	err := loadEnvFilesRecursive()
	if err != nil {
		log.Printf(".env file cannot be found, err: %v", err)
	}

	setRootWorkingDirectoryEnvironment()

	FixProjectRootWorkingDirectoryPath()

	manualEnv := os.Getenv(constants.AppEnv)
	if manualEnv != "" {
		environment = Environment(manualEnv)
	}

	return environment
}

func loadEnvFilesRecursive() error {
	// Start from the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Keep traversing up the directory hierarchy until you find an ".env" file
	for {
		envFilePath := filepath.Join(dir, ".env")
		err := godotenv.Load(envFilePath)
		if err == nil {
			// .env file found and loaded
			return nil
		}

		// Move up on directory level
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// Reached the root directory, stop searching
			break
		}

		dir = parentDir
	}

	return errors.New(".env file not found in the project directory or its parent hierarchy")
}

func setRootWorkingDirectoryEnvironment() {
	absoluteRootWorkingDirectory := GetProjectRootWorkingDirectory()

	// When we `Set` a viper with string value, we should get it from viper with `viper.GetString`, elsewhere we get empty string
	viper.Set(constants.AppRootPath, absoluteRootWorkingDirectory)
}

func (env Environment) IsDevelopment() bool {
	return env == Development
}

func (env Environment) IsProduction() bool {
	return env == Production
}

func (env Environment) IsTest() bool {
	return env == Test
}

func (env Environment) GetEnvironmentName() string {
	return string(env)
}

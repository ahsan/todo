package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetTodosDir() string {
	return GetEnvVar("TODO_DIR")
}

func GetLogLevel() string {
	return GetEnvVar("TODO_LOG_LEVEL")
}

func GetEnvVar(envVar string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file is not present.")
	}

	var value string = os.Getenv(envVar)
	if value == "" {
		fmt.Println(envVar + " environment variable is unset.")
	}
	return value
}

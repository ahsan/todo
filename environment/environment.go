package environment

import (
	"fmt"
	"os"
)

func GetNotesDir() string {
	return getEnvVar("TODO_DIR")
}

func GetLogLevel() string {
	return getEnvVar("TODO_LOG_LEVEL")
}

func getEnvVar(envVar string) string {
	var value string = os.Getenv(envVar)
	if value == "" {
		fmt.Println(envVar + " environment variable is unset.")
	}
	return value
}

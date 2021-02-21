package environment

import (
	"os"
)

func GetTodosDir() string {
	return GetEnvVar("TODO_DIR")
}

func GetLogLevel() string {
	return GetEnvVar("TODO_LOG_LEVEL")
}

func GetEnvVar(envVar string) string {
	var value string = os.Getenv(envVar)
	// if value == "" {
	// 	fmt.Println(envVar + " environment variable is unset.")
	// }
	return value
}

package filesystem

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ahsan/todo/logger"
)

func loadJsonFile(listName string) TodoJson {
	listFileFullPath := getListFullPath(listName)
	marshalledByteArr, err := ioutil.ReadFile(listFileFullPath)
	if err != nil {
		logger.Error("Could not read JSON file")
	}
	var todoJson TodoJson
	json.Unmarshal(marshalledByteArr, &todoJson)
	return todoJson
}

func listExists(listName string) bool {
	if _, err := os.Stat(getListFullPath(listName)); err == nil {
		return true
	}

	return false
}

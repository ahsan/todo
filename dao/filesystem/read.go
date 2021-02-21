package filesystem

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/types"
)

func getTodoJson(listName string) types.TodoJson {
	listFileFullPath := getListFullPath(listName)

	if !listExists(listName) {
		logger.Info("Todo list does not exist")
		return types.TodoJson{}
	}

	marshalledByteArr, err := ioutil.ReadFile(listFileFullPath)
	if err != nil {
		logger.Error("Could not read JSON file: " + err.Error())
	}
	var todoJson types.TodoJson
	json.Unmarshal(marshalledByteArr, &todoJson)
	return todoJson
}

func listExists(listName string) bool {
	if _, err := os.Stat(getListFullPath(listName)); err == nil {
		return true
	}

	return false
}

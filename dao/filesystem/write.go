package filesystem

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/types"
)

func addTodoToList(listName string, todo string) bool {
	if exists := listExists(listName); exists == false {
		createListFile(listName)
	}

	currentTodoJson := getTodoJson(listName)
	updatedTodoJson, err := addTodoToJson(todo, currentTodoJson)
	if err != nil {
		return logger.Error("Could not update Json object")
	}
	return writeJsonFile(listName, updatedTodoJson)
}

func addTodoToJson(todo string, todoJson types.TodoJson) (types.TodoJson, error) {
	// todo: think about deep copying
	var biggestId int = -1
	for _, todo := range todoJson.Todos {
		if todo.Id > biggestId {
			biggestId = todo.Id
		}
	}

	var newTodoId = biggestId + 1
	todoJson.Todos = append(todoJson.Todos, types.Todo{
		Id:          newTodoId,
		Description: todo,
	})
	return todoJson, nil
}

func createListFile(listName string) bool {
	listFileFullPath := getListFullPath(listName)
	_, err := os.Create(listFileFullPath)
	if err != nil {
		logger.Error(err.Error())
	}
	writeJsonFile(listName, emptyTodoJson())
	logger.Info("Created file: " + listFileFullPath)
	return false
}

func writeJsonFile(listName string, todoJson types.TodoJson) bool {
	marshalled, err := json.Marshal(todoJson)
	if err != nil {
		return logger.Error("Could not convert TodoJson object to JSON: " + err.Error())
	}
	err = ioutil.WriteFile(getListFullPath(listName), marshalled, DAEFAULT_FILE_PERM)
	if err != nil {
		return logger.Error("Could not write todoJson to file: " + err.Error())
	}

	return true
}

func emptyTodoJson() types.TodoJson {
	return types.TodoJson{
		Todos: []types.Todo{},
	}
}

package filesystem

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ahsan/todo/logger"
)

func AddTodo(listName string, todo string) bool {
	if exists := listExists(listName); exists == false {
		createListFile(listName)
	}

	currentTodoJson := loadJsonFile(listName)
	updatedTodoJson, err := addTodoToJson(todo, currentTodoJson)
	if err != nil {
		return logger.Error("Could not update Json object")
	}
	fmt.Println("Adding this todo to file: ")
	fmt.Println(updatedTodoJson)
	return writeJsonFile(listName, updatedTodoJson)
}

func addTodoToJson(todo string, todoJson TodoJson) (TodoJson, error) {
	// todo: think about deep copying
	var biggestId int = -1
	for _, todo := range todoJson.Todos {
		if todo.Id > biggestId {
			biggestId = todo.Id
		}
	}

	var newTodoId = biggestId + 1
	todoJson.Todos = append(todoJson.Todos, Todo{
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

func writeJsonFile(listName string, todoJson TodoJson) bool {
	marshalled, err := json.Marshal(todoJson)
	if err != nil {
		return logger.Error("Could not convert TodoJson object to JSON: " + err.Error())
	}
	fmt.Println("Marshalled: ", marshalled)
	err = ioutil.WriteFile(getListFullPath(listName), marshalled, DAEFAULT_FILE_PERM)
	if err != nil {
		return logger.Error("Could not write todoJson to file: " + err.Error())
	}

	return true
}

func emptyTodoJson() TodoJson {
	return TodoJson{
		Todos: []Todo{},
	}
}

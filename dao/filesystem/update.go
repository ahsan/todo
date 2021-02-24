package filesystem

import (
	"fmt"
	"strconv"

	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/types"
)

func updateTodo(listName string, todoId string, statusUpdater func(types.Todo) types.Todo) bool {
	todoJson := getTodoJson(listName)

	for i, todo := range todoJson.Todos {
		if todo.Id == strToInt(todoId) {
			todoJson.Todos[i] = statusUpdater(todo)
		}
	}

	return writeJsonFile(listName, todoJson)
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		logger.Error(fmt.Sprintf("Could not convert todoId %s to int", s))
	}
	return i
}

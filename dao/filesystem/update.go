package filesystem

import (
	"fmt"
	"strconv"

	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/types"
)

func changeStatusToInProgress(todo types.Todo) types.Todo {
	todo.Status.InProgress = true
	return todo
}

func changeStatusToPaused(todo types.Todo) types.Todo {
	todo.Status.Paused = true
	return todo
}

func changeStatusToDone(todo types.Todo) types.Todo {
	todo.Status.InProgress = false
	todo.Status.Paused = false
	todo.Status.Complete = true
	return todo
}

func updateTodo(listName string, todoId string, updater func(types.Todo) types.Todo) bool {
	todoJson := getTodoJson(listName)

	for i, todo := range todoJson.Todos {
		if todo.Id == strToInt(todoId) {
			todoJson.Todos[i] = updater(todo)
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

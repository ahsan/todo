package dao

import (
	"fmt"
	"strconv"

	"github.com/ahsan/todo/dao/filesystem"
	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/types"
)

func AddTodoToList(listName string, todo string) bool {
	return filesystem.AddTodo(listName, todo)
}

func GetTodoList(listName string) []types.Todo {
	todoJson := filesystem.GetTodoJson(listName)
	return todoJson.Todos
}

func MarkTodoAsInProgress(listName string, todoId string) bool {
	return filesystem.UpdateTodo(listName, strToInt(todoId), filesystem.ChangeStatusToInProgress)
}

func MarkTodoAsPaused(listName string, todoId string) bool {
	return filesystem.UpdateTodo(listName, strToInt(todoId), filesystem.ChangeStatusToPaused)
}

func MarkTodoAsComplete(listName string, todoId string) bool {
	return filesystem.UpdateTodo(listName, strToInt(todoId), filesystem.ChangeStatusToDone)
}

func GetAllTodosByStatus(status types.TodoStatus) []types.Todo {
	panic("Not implemented")
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		logger.Error(fmt.Sprintf("Could not convert todoId %s to int", s))
	}
	return i
}

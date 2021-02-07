package dao

import (
	"github.com/ahsan/todo/dao/filesystem"
	"github.com/ahsan/todo/types"
)

func AddTodoToList(listName string, todo string) bool {
	return filesystem.AddTodo(listName, todo)
}

func GetTodoList(listName string) []types.Todo {
	todoJson := filesystem.GetTodoJson(listName)
	return todoJson.Todos
}

func MarkTodoAsInProgress(listName string, todoId int) bool {
	return true
}

func MarkTodoAsComplete(listName string, todoId int) bool {
	return true
}

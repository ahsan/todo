package dao

import "github.com/ahsan/todo/dao/filesystem"

func AddTodoToList(listName string, todo string) bool {
	return filesystem.AddTodo(listName, todo)
}

func GetTodoList(listName string) string {
	return ""
}

func MarkTodoAsInProgress(listName string, todoId int) bool {
	return true
}

func MarkTodoAsComplete(listName string, todoId int) bool {
	return true
}

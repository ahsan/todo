package dao

import "github.com/ahsan/todo/types"

type TodoDAO interface {

	// Add a new Todo item to a list
	AddTodoToList(listName string, todo string) bool

	// Get all Todo items by list
	GetTodoList(listName string) []types.Todo

	// Change Todo status to 'In Progress'
	MarkTodoAsInProgress(listName string, todoId string) bool

	// Change Todo status to 'Paused'
	MarkTodoAsPaused(listName string, todoId string) bool

	// Change Todo status to 'Complete'
	MarkTodoAsComplete(listName string, todoId string) bool
}

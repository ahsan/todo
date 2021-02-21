package filesystem

import (
	"github.com/ahsan/todo/dao/shared"
	"github.com/ahsan/todo/types"
)

type FilesystemDAO struct{}

func (f FilesystemDAO) AddTodoToList(listName string, todo string) bool {
	return addTodoToList(listName, todo)
}

func (f FilesystemDAO) GetTodoList(listName string) []types.Todo {
	return getTodoJson(listName).Todos
}

func (f FilesystemDAO) MarkTodoAsInProgress(listName string, todoId string) bool {
	return updateTodo(listName, todoId, shared.ChangeStatusToInProgress)
}

func (f FilesystemDAO) MarkTodoAsPaused(listName string, todoId string) bool {
	return updateTodo(listName, todoId, shared.ChangeStatusToPaused)
}

func (f FilesystemDAO) MarkTodoAsComplete(listName string, todoId string) bool {
	return updateTodo(listName, todoId, shared.ChangeStatusToDone)
}

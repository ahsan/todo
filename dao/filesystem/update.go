package filesystem

import (
	"github.com/ahsan/todo/types"
)

func ChangeStatusToInProgress(todo types.Todo) types.Todo {
	todo.Status.InProgress = true
	return todo
}

func ChangeStatusToPaused(todo types.Todo) types.Todo {
	todo.Status.Paused = true
	return todo
}

func ChangeStatusToDone(todo types.Todo) types.Todo {
	todo.Status.InProgress = false
	todo.Status.Paused = false
	todo.Status.Complete = true
	return todo
}

func UpdateTodo(listName string, todoId int, updater func(types.Todo) types.Todo) bool {
	todoJson := GetTodoJson(listName)

	for i, todo := range todoJson.Todos {
		if todo.Id == todoId {
			todoJson.Todos[i] = updater(todo)
		}
	}

	return writeJsonFile(listName, todoJson)
}

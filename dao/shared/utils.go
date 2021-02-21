package shared

import "github.com/ahsan/todo/types"

func ChangeStatusToInProgress(todo types.Todo) types.Todo {
	todo.Status = types.Statuses["inProgress"]
	return todo
}

func ChangeStatusToPaused(todo types.Todo) types.Todo {
	todo.Status = types.Statuses["paused"]
	return todo
}

func ChangeStatusToDone(todo types.Todo) types.Todo {
	todo.Status = types.Statuses["complete"]
	return todo
}

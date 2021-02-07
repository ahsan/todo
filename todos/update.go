package todos

import "github.com/ahsan/todo/dao"

func StartTodo(todoId string) {
	dao.MarkTodoAsInProgress(getTodaysTodoList(), todoId)
}

func PauseTodo(todoId string) {
	dao.MarkTodoAsPaused(getTodaysTodoList(), todoId)
}

func DoneTodo(todoId string) {
	dao.MarkTodoAsComplete(getTodaysTodoList(), todoId)
}

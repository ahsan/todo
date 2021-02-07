package todos

import "github.com/ahsan/todo/dao"

func StartTodo(todoId int) {
	dao.MarkTodoAsInProgress(getTodaysTodoList(), todoId)
}

func CompleteTodo(todoId int) {
	dao.MarkTodoAsComplete(getTodaysTodoList(), todoId)
}

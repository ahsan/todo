package todos

import (
	"time"

	"github.com/ahsan/todo/dao"
)

func getTodaysTodoList() string {
	return time.Now().Local().Format("20060102")
}

func AddTodo(todo string) {
	var todaysListName = getTodaysTodoList()
	dao.AddTodoToList(todaysListName, todo)
}

func StartTodo(todoId int) {
	dao.MarkTodoAsInProgress(getTodaysTodoList(), todoId)
}

func CompleteTodo(todoId int) {
	dao.MarkTodoAsComplete(getTodaysTodoList(), todoId)
}

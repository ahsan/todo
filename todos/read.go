package todos

import (
	"github.com/ahsan/todo/logger"
)

func (t Todos) GetTodosForToday() {
	l := t.dao.GetTodoList(getTodaysTodoList())
	if l != nil {
		logger.PrettyPrintTodos(l)
	}
}

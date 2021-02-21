package todos

import (
	"github.com/ahsan/todo/logger"
)

func (t Todos) GetTodosForToday() {
	logger.PrettyPrintTodos(t.dao.GetTodoList(getTodaysTodoList()))
}

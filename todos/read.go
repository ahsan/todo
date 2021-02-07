package todos

import (
	"github.com/ahsan/todo/dao"
	"github.com/ahsan/todo/logger"
)

func GetTodosForToday() {
	logger.PrettyPrintTodos(dao.GetTodoList(getTodaysTodoList()))
}

package todos

import (
	"github.com/ahsan/todo/dao"
)

func AddTodo(todo string) {
	var todaysListName = getTodaysTodoList()
	dao.AddTodoToList(todaysListName, todo)
}

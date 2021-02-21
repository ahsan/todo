package todos

import "github.com/ahsan/todo/logger"

func (t Todos) AddTodo(todo string) {
	t.dao.AddTodoToList(getTodaysTodoList(), todo)
}

func (t Todos) AddTodos(todosSlice []string) {
	for _, todo := range todosSlice {
		logger.Debug("Adding todo: '" + todo + "'")
		t.AddTodo(todo)
	}
}

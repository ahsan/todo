package todos

func (t Todos) StartTodo(todoId string) {
	t.dao.MarkTodoAsInProgress(getTodaysTodoList(), todoId)
}

func (t Todos) PauseTodo(todoId string) {
	t.dao.MarkTodoAsPaused(getTodaysTodoList(), todoId)
}

func (t Todos) DoneTodo(todoId string) {
	t.dao.MarkTodoAsComplete(getTodaysTodoList(), todoId)
}

package todos

func (t Todos) AddTodo(todo string) {
	var todaysListName = getTodaysTodoList()
	t.dao.AddTodoToList(todaysListName, todo)
}

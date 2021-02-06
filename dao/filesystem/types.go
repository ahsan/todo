package filesystem

type TodoStatus struct {
	InProgress bool `json:"inProgress"`
	Complete   bool `json:"complete"`
}

type Todo struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
}

type TodoJson struct {
	Todos []Todo `json:"todos"`
}

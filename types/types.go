package types

var Statuses = map[string]string{
	"created":    "created",
	"inProgress": "in_progress",
	"paused":     "paused",
	"complete":   "complete",
}

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TodoJson struct {
	Todos []Todo `json:"todos"`
}

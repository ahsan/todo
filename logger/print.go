package logger

import (
	"fmt"
	"os"

	"github.com/ahsan/todo/types"
	"github.com/olekukonko/tablewriter"
)

func PrettyPrintTodos(todos []types.Todo) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status", "Description"})

	for _, todo := range todos {
		table.Append([]string{fmt.Sprintf("%d", todo.Id), getStatusEmoji(todo.Status), todo.Description})
	}
	table.SetBorder(false)
	table.SetRowLine(true)
	table.Render()
}

func getStatusEmoji(todoStatus string) string {
	Info(todoStatus)
	var statusMap = map[string]string{
		types.Statuses["created"]:    "🦕",
		types.Statuses["inProgress"]: "🚧",
		types.Statuses["paused"]:     "⏸",
		types.Statuses["complete"]:   "✅",
	}
	return statusMap[todoStatus]
}

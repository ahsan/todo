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
	table.Render()
}

func getStatusEmoji(todoStatus types.TodoStatus) string {
	if todoStatus.InProgress {
		return "🚧"
	} else if todoStatus.Complete {
		return "✅"
	} else {
		return "🦕"
	}
}

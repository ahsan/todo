package router

import (
	"fmt"

	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/todos"
	"github.com/urfave/cli/v2"
)

func Route(function string, args cli.Args) {
	switch function {
	case "add":
		var l int = args.Len()
		if l <= 0 {
			logger.Error("Todo not specified.")
			break
		}

		for _, todo := range args.Slice() {
			logger.Debug("Adding todo: '" + todo + "'")
			todos.AddTodo(todo)
		}
		logger.Info(fmt.Sprintf("Added %d todo items.", l))
		break
	}
}

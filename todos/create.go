package todos

import (
	"time"

	"github.com/ahsan/todo/dao"
	"github.com/ahsan/todo/logger"
)

func getTodaysTodoList() string {
	return time.Now().Local().Format("20060102")
}

func AddTodo(todo string) {
	var todaysListName = getTodaysTodoList()

	// create todo if it doesn't exist
	if listExists, _ := dao.ListExists(todaysListName); listExists == false {
		logger.Debug("List did not exist for today, creating")
		dao.CreateList(todaysListName)
		logger.Debug("List created")
	}

}

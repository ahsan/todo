package filesystem

import (
	"os"

	"github.com/ahsan/todo/logger"
)

func CreateList(listName string) bool {
	if exists, err := ListExists(listName); exists == true {
		return logger.Error(err.Error())
	}

	_, err := os.Create(getListFullPath(listName))
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Created file: " + getListFullPath(listName))
	return false
}

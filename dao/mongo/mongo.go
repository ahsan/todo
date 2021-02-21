package mongo

import (
	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/types"
)

type MongodbDAO struct{}

func (m MongodbDAO) AddTodoToList(listName string, todo string) bool {
	logger.Error("Not implemented")
	return false
}

func (m MongodbDAO) GetTodoList(listName string) []types.Todo {
	logger.Error("Not implemented")
	return nil
}

func (m MongodbDAO) MarkTodoAsInProgress(listName string, todoId string) bool {
	logger.Error("Not implemented")
	return false
}

func (m MongodbDAO) MarkTodoAsPaused(listName string, todoId string) bool {
	logger.Error("Not implemented")
	return false
}

func (m MongodbDAO) MarkTodoAsComplete(listName string, todoId string) bool {
	logger.Error("Not implemented")
	return false
}

package mongo

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbDAO struct{}

func (m MongodbDAO) AddTodoToList(listName string, todo string) bool {
	c := getClient()
	defer c.cancel()

	return writeToList(*c.ctx, *c.client, listName, todo)
}

func (m MongodbDAO) GetTodoList(listName string) []types.Todo {
	c := getClient()
	defer c.cancel()

	return readList(*c.ctx, *c.client, listName)
}

func (m MongodbDAO) MarkTodoAsInProgress(listName string, todoId string) bool {
	return m.updateStatus(listName, todoId, types.Statuses["inProgress"])
}

func (m MongodbDAO) MarkTodoAsPaused(listName string, todoId string) bool {
	return m.updateStatus(listName, todoId, types.Statuses["paused"])
}

func (m MongodbDAO) MarkTodoAsComplete(listName string, todoId string) bool {
	return m.updateStatus(listName, todoId, types.Statuses["complete"])
}

type getClientResponse struct {
	client *mongo.Client
	ctx    *context.Context
	cancel context.CancelFunc
}

func getClient() getClientResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getMongoURI()))
	if err != nil {
		log.Fatal(err)
	}
	return getClientResponse{client: client, ctx: &ctx, cancel: cancel}
}

func (m MongodbDAO) updateStatus(listName string, todoId string, newStatus string) bool {
	c := getClient()
	defer c.cancel()

	todoIdInt, err := strconv.Atoi(todoId)
	if err != nil {
		logger.Error("Could not convert todo id to integer")
		return false
	}

	return updateTodoStatus(*c.ctx, *c.client, listName, todoIdInt, newStatus)
}

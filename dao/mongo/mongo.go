package mongo

import (
	"context"
	"log"
	"time"

	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbDAO struct{}

func (m MongodbDAO) AddTodoToList(listName string, todo string) bool {
	logger.Error("Not implemented")
	return false
}

func (m MongodbDAO) GetTodoList(listName string) []types.Todo {
	c := getClient()
	defer c.cancel()

	return readList(*c.ctx, *c.client, listName)
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

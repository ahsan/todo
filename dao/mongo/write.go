package mongo

import (
	"context"

	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/types"
	"go.mongodb.org/mongo-driver/mongo"
)

func writeToList(ctx context.Context, client mongo.Client, listName string, todo string) bool {
	collection := client.Database(getDBName()).Collection(TodoCollectionName)

	existingTodos := readList(ctx, client, listName)
	newId := getNextId(existingTodos)

	_, err := collection.InsertOne(ctx, MongoTodo{
		Id:          newId,
		ListName:    listName,
		Status:      types.Statuses["created"],
		Description: todo,
	})

	if err != nil {
		return logger.Error("Could not add todo to DB")
	}

	logger.Debug("Added todo to " + listName + ".")
	return true
}

package mongo

import (
	"context"

	"github.com/ahsan/todo/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func updateTodoStatus(ctx context.Context, client mongo.Client, listName string, todoId int, newStatus string) bool {
	collection := client.Database(getDBName()).Collection(TodoCollectionName)

	var todoUpdate = map[string]string{
		"status": newStatus,
	}

	update := bson.M{
		"$set": todoUpdate,
	}

	result := collection.FindOneAndUpdate(ctx, bson.D{{"listname", listName}, {"id", todoId}}, update)
	if result.Err() != nil {
		return logger.Error("Failed to update todo: " + result.Err().Error())
	}
	return true
}

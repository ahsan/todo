package mongo

import (
	"context"

	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Document model
/**
{
	_id: auto generated mongo id
	id: int number
	list_name: string
	description: string
	status: string
}
*/

func readList(ctx context.Context, client mongo.Client, listName string) []types.Todo {
	collection := client.Database("todo-app").Collection(TodoCollectionName)

	var todos []types.Todo

	opts := options.Find()

	opts.SetSort(bson.D{{"id", 1}})

	cursor, err := collection.Find(ctx, bson.D{{"listname", listName}})
	if err != nil {
		logger.Error("Could not read from MongoDB: " + err.Error())
		panic(err)
	}

	for cursor.Next(ctx) {
		var mongoTodo MongoTodo
		err := cursor.Decode(&mongoTodo)
		if err != nil {
			logger.Error("Could not decode")
			break
		}
		logger.Debug(mongoTodo.String())

		todos = append(todos, types.Todo{Id: mongoTodo.Id, Description: mongoTodo.Description, Status: mongoTodo.Status})
	}

	cursor.Close(ctx)
	return todos
}

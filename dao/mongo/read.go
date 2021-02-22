package mongo

import (
	"context"

	"github.com/ahsan/todo/logger"
	"github.com/ahsan/todo/types"
	"go.mongodb.org/mongo-driver/mongo"
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

// see types.go#Todo
type MongoTodo struct {
	Id          int    `json:"id" json:"id"`
	ListName    string `json:"listName" json:"listName"`
	Description string `json:"description" json:"description"`
	Status      string `json:"status" json:"status"`
}

func readList(ctx context.Context, client mongo.Client, listName string) []types.Todo {
	collection := client.Database("todo-app").Collection(TodoCollectionName)

	var todos []types.Todo

	cursor, err := collection.Find(ctx, map[string]string{"listName": listName})
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
		todos = append(todos, types.Todo{Id: mongoTodo.Id, Description: mongoTodo.Description, Status: mongoTodo.Status})
	}
	return todos
}

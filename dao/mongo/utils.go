package mongo

import (
	"strconv"
	"time"

	"github.com/ahsan/todo/environment"
	"github.com/ahsan/todo/types"
)

// see types.go#Todo
type MongoTodo struct {
	Id          int    `json:"id" json:"id"`
	ListName    string `json:"listname" json:"listname"`
	Description string `json:"description" json:"description"`
	Status      string `json:"status" json:"status"`
}

func (t MongoTodo) String() string {
	return "id: " + strconv.Itoa(t.Id) +
		", listName: " + t.ListName +
		", description: " + t.Description +
		", status: " + t.Status
}

const TodoCollectionName string = "todos"
const DefaultTimeout time.Duration = 5 * time.Second

func getMongoURI() string {
	username := environment.GetEnvVar("MONGO_USERNAME")
	password := environment.GetEnvVar("MONGO_PASSWORD")
	db := getDBName()
	url := environment.GetEnvVar("MONGO_URI")

	return "mongodb+srv://" + username + ":" + password + "@" + url + "/" + db
}

func getDBName() string {
	return environment.GetEnvVar("MONGO_DB")
}

func getNextId(todos []types.Todo) int {
	var biggestId int = -1

	for _, todo := range todos {
		if biggestId < todo.Id {
			biggestId = todo.Id
		}
	}
	return biggestId + 1
}

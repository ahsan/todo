package mongo

import (
	"time"

	"github.com/ahsan/todo/environment"
)

const TodoCollectionName string = "todos"
const DefaultTimeout time.Duration = 5 * time.Second

func getMongoURI() string {
	username := environment.GetEnvVar("MONGO_USERNAME")
	password := environment.GetEnvVar("MONGO_PASSWORD")
	db := environment.GetEnvVar("MONGO_DB")
	url := environment.GetEnvVar("MONGO_URI")

	return "mongodb+srv://" + username + ":" + password + "@" + url + "/" + db
}

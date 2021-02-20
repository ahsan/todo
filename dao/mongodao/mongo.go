package mongodao

import (
	"context"
	"log"
	"time"

	"github.com/ahsan/todo/environment"
	"github.com/ahsan/todo/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mongo_dao() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("testing").Collection("numbers")
	logger.Debug(collection.Name())

	defer client.Disconnect(ctx)
}

func getMongoURI() string {
	username := environment.GetEnvVar("MONGO_USERNAME")
	password := environment.GetEnvVar("MONGO_PASSWORD")
	db := environment.GetEnvVar("MONGO_DB")
	url := environment.GetEnvVar("MOGO_URI")

	return "mongodb+srv://" + username + ":" + password + "@" + url + "/" + db
}

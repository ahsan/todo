package mongo_test

import (
	"github.com/ahsan/todo/dao"
	"github.com/ahsan/todo/dao/mongo"
)

var _ dao.TodoDAO = mongo.MongodbDAO{}

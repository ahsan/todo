package dao

import (
	"github.com/ahsan/todo/dao/filesystem"
	"github.com/ahsan/todo/dao/mongo"
	"github.com/ahsan/todo/environment"
	"github.com/ahsan/todo/logger"
)

const envVarKey = "TODO_STORAGE_STRATEGY"

func DaoFactory() TodoDAO {
	var dao TodoDAO
	storageStrategy := environment.GetEnvVar(envVarKey)

	switch storageStrategy {
	case "filesystem":
		dao = filesystem.FilesystemDAO{}
	case "mongodb":
		dao = mongo.MongodbDAO{}
	default:
		logger.Warning("Defaulting strategy to 'filesystem'.")
		dao = filesystem.FilesystemDAO{}
	}

	return dao
}

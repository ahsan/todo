package filesystem_test

import (
	"github.com/ahsan/todo/dao"
	"github.com/ahsan/todo/dao/filesystem"
)

var _ dao.TodoDAO = filesystem.FilesystemDAO{}

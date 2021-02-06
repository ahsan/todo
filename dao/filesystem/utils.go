package filesystem

import (
	"path/filepath"

	"github.com/ahsan/todo/environment"
)

const DAEFAULT_FILE_PERM = 0644

func getListFullPath(listName string) string {
	return filepath.Join(environment.GetNotesDir(), listName)
}

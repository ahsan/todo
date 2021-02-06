package filesystem

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ahsan/todo/environment"
)

func ListExists(listName string) (bool, error) {
	if _, err := os.Stat(getListFullPath(listName)); err == nil {
		return true, fmt.Errorf("todo list %s already exists.", listName)
	}

	return false, nil
}

func getListFullPath(listName string) string {
	return filepath.Join(environment.GetNotesDir(), listName)
}

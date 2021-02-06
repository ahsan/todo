package dao

import "github.com/ahsan/todo/dao/filesystem"

func CreateList(listName string) bool {
	return filesystem.CreateList(listName)
}

func ListExists(listName string) (bool, error) {
	return filesystem.ListExists(listName)
}

package main

import "os"

func GetNotesDir() string {
	return os.Getenv("NOTES_DIR")
	// todo add sanity checks
}

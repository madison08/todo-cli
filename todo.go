package main

import (
	"encoding/json"
	"log"
	"os"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func saveTasks(tasks []Todo) {
	results, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	writeErr := os.WriteFile(FILE_PATH, results, os.FileMode(0644))

	if writeErr != nil {
		log.Fatal(writeErr)
	}

}

func loadTasks() []Todo {

	jsonData, err := os.ReadFile(FILE_PATH)

	if err != nil {
		return []Todo{}
	}

	var todos []Todo

	unmarshalErr := json.Unmarshal(jsonData, &todos)

	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}

	return todos
}

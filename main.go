package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const FILE_PATH string = "tasks.json"

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

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Veuillez saisir au moins un argument")
	}

	tasks := loadTasks()

	arg := os.Args[1]

	switch arg {
	case "add":
		fmt.Println("Ajout d'une tâche...")
		if len(os.Args) > 2 {
			tasks = append(tasks, Todo{ID: len(tasks) + 1, Title: os.Args[2], Done: false})
			saveTasks(tasks)
			fmt.Println(os.Args[2], "a été ajouter avec succès !")
		}
	case "list":
		for _, task := range tasks {
			fmt.Println(task.Title)
		}
	case "help", "":
		fmt.Println("Commandes disponibles :")
		fmt.Println("  add  <titre>  → Ajoute une tâche")
		fmt.Println("  list         → Affiche les tâches")
	default:
		fmt.Println("Commande inconnue.")
		fmt.Println("Tape `help` pour voir les commandes disponibles.")

	}

}

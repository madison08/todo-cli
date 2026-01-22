package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/fatih/color"
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
		log.Fatal("Veuillez saisir au moins un argument")
	}

	greenText := color.New(color.FgGreen)
	redText := color.New(color.FgRed)

	tasks := loadTasks()

	arg := os.Args[1]

	switch arg {
	case "add":
		fmt.Println("Ajout d'une tâche...")
		now := time.Now()
		unixTime := now.Unix()
		if len(os.Args) > 2 {
			tasks = append(tasks, Todo{ID: int(unixTime), Title: os.Args[2], Done: false})
			saveTasks(tasks)
			fmt.Println(os.Args[2], "a été ajouter avec succès !")
		}
	case "done":
		if len(os.Args) > 2 {
			taskId, err := strconv.Atoi(os.Args[2])

			if err != nil {
				log.Fatal(err)
			}

			for index, task := range tasks {
				if task.ID == taskId {
					tasks[index].Done = true
					break
				}
			}
			saveTasks(tasks)
		}

	case "delete":
		if len(os.Args) > 2 {

			taskId, err := strconv.Atoi(os.Args[2])

			if err != nil {
				log.Fatal(err)
			}

			for _, task := range tasks {
				if task.ID == taskId {
					index := slices.Index(tasks, task)
					if index != -1 {
						tasks = slices.Delete(tasks, index, index+1)
					}
				}
			}

			saveTasks(tasks)
		}

	case "list":

		for _, task := range tasks {
			if task.Done {
				greenText.Printf("[X] - %v - %v \n", task.ID, task.Title)

			} else {
				redText.Printf("[ ] - %v - %v \n", task.ID, task.Title)
			}

		}
	case "pending":
		for _, task := range tasks {
			if !task.Done {
				redText.Printf("[ ] - %v - %v \n", task.ID, task.Title)

			}
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

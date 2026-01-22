package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

const FILE_PATH string = "tasks.json"

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
		handleAddTask(tasks, os.Args[2])
	case "done":
		handleDoneTask(tasks, os.Args)
	case "delete":
		handleDeleteTask(tasks, os.Args)

	case "list":
		handleGetTasks(tasks, greenText, redText)
	case "pending":
		handleGetPendingTasks(tasks, redText)

	case "help", "":
		fmt.Println("Commandes disponibles :")
		fmt.Println("  add  <titre>  → Ajoute une tâche")
		fmt.Println("  list         → Affiche les tâches")
	default:
		fmt.Println("Commande inconnue.")
		fmt.Println("Tape `help` pour voir les commandes disponibles.")

	}

}

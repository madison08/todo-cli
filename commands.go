package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func handleAddTask(tasks []Todo, title string) {
	fmt.Println("Ajout d'une tâche...")
	now := time.Now()
	unixTime := now.Unix()
	if len(os.Args) > 2 {
		tasks = append(tasks, Todo{ID: int(unixTime), Title: title, Done: false})
		saveTasks(tasks)
		fmt.Println(os.Args[2], "a été ajouter avec succès !")
	}
}

func handleDoneTask(tasks []Todo, args []string) {
	if len(args) > 2 {
		taskId, err := strconv.Atoi(args[2])

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
}

func handleDeleteTask(tasks []Todo, args []string) {
	if len(args) > 2 {

		taskId, err := strconv.Atoi(args[2])

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
}

func handleGetTasks(tasks []Todo, doneColor *color.Color, undoneColor *color.Color) {

	for _, task := range tasks {
		if task.Done {
			doneColor.Printf("[X] - %v - %v \n", task.ID, task.Title)

		} else {
			undoneColor.Printf("[ ] - %v - %v \n", task.ID, task.Title)
		}

	}
}

func handleGetPendingTasks(tasks []Todo, undoneColor *color.Color) {
	for _, task := range tasks {
		if !task.Done {
			undoneColor.Printf("[ ] - %v - %v \n", task.ID, task.Title)

		}
	}
}

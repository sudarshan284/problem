package main

import (
	"fmt"
	"os"
	"strconv"
)

const filePath = "todos.json"

func main() {
	repo := NewRepository(filePath)
	service := NewService(repo)

	if len(os.Args) < 2 {
		fmt.Println("Usage : todo c-cli [add list]")
		return
	}
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli add <task title>")
			return
		}

		title := os.Args[2]
		err := service.AddTask(title)
		if err != nil {
			fmt.Println("error adding task : ", err)
		} else {
			fmt.Println("taks added successfully!!")
		}

	case "list":
		tasks, err := service.ListTasks()
		if err != nil {
			fmt.Println("Error listing tasks : ", err)
			return
		}
		fmt.Println("To do list")
		for _, task := range tasks {
			status := "Pending"
			if task.Completed {
				status = "Completed"
			}
			fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
		}
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo-cli complete <task ID>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := service.CompleteTask(id)
		if err != nil {
			fmt.Println("Error completing task:", err)
		} else {
			fmt.Println("Task marked as completed!")
		}
	default:
		fmt.Println("unknown command")
	}
}

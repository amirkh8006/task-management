package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-management/task"
)

func main() {
	todo := task.ToDoList{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n======= Todo List Menu =======")
		fmt.Println("1. Add Task")
		fmt.Println("2. Display Tasks")
		fmt.Println("3. Update Task Status")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title := scanner.Text()
			todo.AddTask(title)
		case "2":
			todo.DisplayTask()
		case "3":
			fmt.Print("Enter task ID to update: ")
			scanner.Scan()
		    taskID, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Is the task done? (true/false): ")
			scanner.Scan()
			status := scanner.Text()
			todo.UpdateTaskStatus(taskID, strings.ToLower(status) == "true")
		case "4":
			fmt.Print("Enter task ID to delete: ")
			scanner.Scan()
			taskID, _ := strconv.Atoi(scanner.Text())
			todo.DeleteTask(taskID)
		case "5":
			fmt.Println("Exiting...")
			os.Exit(0)
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}		
	}
}
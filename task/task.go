package task

import "fmt"

type Task struct {
	ID int
	Title string
	IsDone bool
}


type ToDoList struct {
	Tasks []Task
}


func (t *ToDoList) AddTask(title string) {
	task := Task{
		ID: len(t.Tasks) + 1,
		Title: title,
		IsDone: false,
	}
	t.Tasks = append(t.Tasks, task)
	fmt.Println("Task added successfully:", task)
}

func (t *ToDoList) DisplayTask() {
	if len(t.Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range t.Tasks {
		doneStatus := ""
		if task.IsDone {
			doneStatus = "X"
		}
		fmt.Printf("[%s] %d: %s\n", doneStatus, task.ID, task.Title)
	}
}

func (t *ToDoList) UpdateTaskStatus(taskID int, isDone bool) {
	for i, task := range t.Tasks {
		if task.ID == taskID {
			t.Tasks[i].IsDone = isDone
			fmt.Printf("Task %d status updated to %t\n", taskID, isDone)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", taskID)
}

func (t *ToDoList) DeleteTask(taskID int) {
	for i, task := range t.Tasks {
		if task.ID == taskID {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			fmt.Printf("Task %d deleted successfully.\n", taskID)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", taskID)
}
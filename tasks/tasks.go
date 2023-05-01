package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Comlete bool   `json:"complete"`
}

func SaveTasks(file *os.File, tasks []Task) error {
	// elimino todo el contenido del archivo
	err := file.Truncate(0)
	if err != nil {
		return err
	}

	// convierto el []Task a bytes
	bytes, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func LsTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet...")
		return
	}

	for _, task := range tasks {
		complete := "[ ]"
		if task.Comlete {
			complete = "[✔]"
		}
		fmt.Printf("%s (%d) %s\n", complete, task.ID, task.Name)
	}
}

func AddTask(file *os.File, tasks []Task) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your task?: ")
	task, _ := reader.ReadString(byte(10)) // byte(10) = \n

	newTask := Task{
		ID:      GenerateId(tasks),
		Name:    strings.TrimSpace(task),
		Comlete: false,
	}
	tasks = append(tasks, newTask)
	err := SaveTasks(file, tasks)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Task created (%d)\n", newTask.ID)
}

func DeleteTask(file *os.File, tasks []Task) {
	if len(os.Args) < 3 {
		fmt.Println("You must enter an id after the delete command")
		return
	}

	foundTask := false
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("The id must be a number")
		return
	}

	for index, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			foundTask = true
			break
		}
	}

	if !foundTask {
		fmt.Println("Task not found...")
		return
	}

	err = SaveTasks(file, tasks)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted task (%d)\n", id)
}

func CompleteTask(file *os.File, tasks []Task) {
	if len(os.Args) < 3 {
		fmt.Println("You must enter an id after the complete command")
		return
	}

	foundTask := false
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("The id must be a number")
		return
	}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks[i].Comlete = true
			foundTask = true
			break
		}
	}

	if !foundTask {
		fmt.Println("Task not found...")
		return
	}

	err = SaveTasks(file, tasks)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Task completed (%d)\n", id)
}

func GenerateId(tasks []Task) int {
	return len(tasks) + 1
}

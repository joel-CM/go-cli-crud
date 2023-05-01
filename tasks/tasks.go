package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Comlete bool   `json:"complete"`
}

func (t Task) Save(file *os.File) error {
	var tasks []Task

	bytes, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &tasks)
	if err != nil {
		return err
	}

	tasks = append(tasks, t)

	tasksToByte, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	_, err = file.Write(tasksToByte)
	if err != nil {
		return err
	}

	return nil
}

func LsTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas...")
		return
	}

	for _, task := range tasks {
		complete := "[x]"
		if task.Comlete {
			complete = "[-]"
		}
		fmt.Printf("%s %s\n", complete, task.Name)
	}
}

func AddTask(file *os.File) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("¿Cual es tu tarea?: ")
	task, _ := reader.ReadString(byte(10))

	newTask := Task{
		ID:      10,
		Name:    strings.TrimSpace(task),
		Comlete: false,
	}
	newTask.Save(file)
}

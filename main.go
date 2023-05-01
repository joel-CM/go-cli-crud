package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pkgTasks "github.com/joel-CM/go-cli-crud/tasks"
)

func main() {

	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fail(err)
	}

	var tasks []pkgTasks.Task

	fileInfo, err := file.Stat()
	if err != nil {
		fail(err)
	}

	if fileInfo.Size() != 0 {
		bytes, err := ioutil.ReadFile(file.Name())
		if err != nil {
			fail(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			fail(err)
		}
	} else {
		file.Write([]byte("[]"))
	}

	if len(os.Args) < 2 {
		HelpUse()
		return
	}

	switch os.Args[1] {
	case "ls":
		pkgTasks.LsTasks(tasks)
	case "add":
		pkgTasks.AddTask(file, tasks)
	case "delete":
		pkgTasks.DeleteTask(file, tasks)
	default:
		HelpUse()
	}
}

func fail(err error) {
	log.Fatal(err)
}

func HelpUse() {
	fmt.Println("comands: [ls|add|delete]")
}

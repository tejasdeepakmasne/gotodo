package operations

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	task "github.com/tejasdeepakmasne/gotodo/model"
)

func LoadFile() []task.Task {
	filename := "tasks.json"

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting fileInfo %v", err)
	}

	var tasks []task.Task

	if fileInfo.Size() > 0 {
		err = json.NewDecoder(file).Decode(&tasks)
		if err != nil {
			log.Fatalln("Error decoding json", err)
		}
	} else {
		fmt.Println("tasks.json is empty, initializing default format")
		err = json.NewEncoder(file).Encode(tasks)
		if err != nil {
			log.Fatalln("Error encoding json", err)
		}
	}

	return tasks
}

func WriteTasks(tasks *[]task.Task) error {
	filename := "tasks.json"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(*tasks)
	if err != nil {
		return err
	}
	return nil
}

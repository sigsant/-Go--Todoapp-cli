package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type item struct {
	Task        string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

var logFile, _ = os.Create("debug.log")

var LoggerHandler = log.New(logFile, "Info:", log.LstdFlags|log.Lshortfile)

//	Add an item to the task list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

//	Complete marks a task as complete (at x time)
func (l *List) Complete(index int) error {
	list := *l
	// Error si el dato introducido no es positivo o sobrepasa la longitud de la lista
	if index < 0 || index > len(list) {
		return fmt.Errorf("There is no item in position %d", index)
	}

	list[index-1].Completed = true
	list[index-1].CompletedAt = time.Now()
	LoggerHandler.Println("Task changed")

	return nil
}

//	Delete a task from the list
func (l *List) Delete(index int) error {
	list := *l
	if index < 0 || index > len(list) {
		return fmt.Errorf("There is no item in position %d", index)
	}
	//	Reacomoda la capacidad del slice
	*l = append(list[:index-1], list[index:]...)
	LoggerHandler.Println("Task deleted!")
	return nil
}

//	Save saves the notes in JSON format
func (l *List) Save(filename string) error {
	fileJSON, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		return err
	}
	os.WriteFile(filename, fileJSON, 0644)
	LoggerHandler.Println("File saved!")

	return nil
}

func (l *List) Read(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			LoggerHandler.Println("JSON not exist")
			return nil
		}
		return err
	}

	//	Ignora el error si el archivo esta vacio
	if len(filename) == 0 {
		LoggerHandler.Println("Empty JSON")
		return nil
	}

	// Devuelve los datos a partir del struct 'Item'
	LoggerHandler.Println("Read file done")
	return json.Unmarshal(file, l)
}

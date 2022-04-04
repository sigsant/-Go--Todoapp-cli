package todo

import (
	"encoding/json"
	"fmt"
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

	list[index].Completed = true
	list[index].CompletedAt = time.Now()

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
	return nil
}

func (l *List) Save(filename string) error {
	fileJSON, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		return err
	}
	os.WriteFile(filename, fileJSON, 0644)
	return nil
}

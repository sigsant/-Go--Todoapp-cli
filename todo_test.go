package todo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var lista = List{}

var taskExample = []struct {
	tasks string
}{
	{"Do shopping"},
	{"Read a book"},
	{"Testing in Go"},
}

func TestCreateTask(t *testing.T) {
	for index, value := range taskExample {
		lista.Add(value.tasks)
		assert.Equal(t, lista[index].Task, value.tasks)
	}
	assert.Equal(t, len(lista), 3)
}

func TestCompleteTask(t *testing.T) {
	var taskComplete = []struct {
		tasks      string
		isComplete bool
		errorMsg   string
	}{
		{"Sing a song", true, "Task 'Sing a song' should be not complete"},
		{"Take a walk", true, "Task 'Take a walk' should be complete"},
	}
	for index, value := range taskComplete {
		lista.Add(value.tasks)
		result := lista.Complete(index)
		assert.Nil(t, result)
		assert.Equal(t, lista[index].Completed, value.isComplete, value.errorMsg)
	}
}

func TestDeleteTask(t *testing.T) {
	lista.Delete(1)
	assert.Equal(t, len(lista), 4, "Error en la longitud de la lista")
}

func TestSaveTask(t *testing.T) {
	_ = lista.Save("test.json")

	_, err := os.Stat("test.json")
	assert.Nil(t, err, "No se ha encontrado el archivo")
	_ = os.Remove("test.json")
}

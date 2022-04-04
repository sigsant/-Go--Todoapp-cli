package main

import (
	"fmt"
	"os"

	"github.com/sigsant/todo"
)

//	Placeholder for testing
const filename = "task.json"

func main() {

	list := &todo.List{}

	// Try to read the file. If not it is possible display an error and quit the app
	if err := list.Read(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//	TODO: Considerar las acciones a realizar según si hay parametros al iniciarse el programa
	//	TODO: Logicamente, traducir los comentarios al inglés a medida que se desarrolle la aplicacion
}

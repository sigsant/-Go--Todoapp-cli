//	TODO: Traducir los comentarios al inglés a medida que se desarrolle la aplicacion
// TODO:  Observar las buenas prácticas recomendadas por Gophers.

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sigsant/todo"
)

//	Placeholder for testing
const filename = "task.json"

var list = &todo.List{}

//	TODO: Considerar las acciones a realizar según si hay parametros al iniciarse el programa

func main() {

	// Try to read the file. If not it is possible display an error and quit the app
	if err := list.Read(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch len(os.Args) {
	case 1:
		for i, task := range *list {
			fmt.Printf("\t%d. %s\n", i, task.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		list.Add(item)
	}

	if err := list.Save(filename); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

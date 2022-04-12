//	TODO: Traducir los comentarios al inglés a medida que se desarrolle la aplicacion
// TODO:  Observar las buenas prácticas recomendadas por Gophers.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sigsant/todo"
)

//	Placeholder for testing
const filename = "task.json"

var list = &todo.List{}

func main() {

	newTask := flag.String("task", "", "Task to be included in the notes")
	listTask := flag.Bool("list", false, "List all the saved task")
	isTaskCompleted := flag.Int("complete", 0, "Number of the item to be marked as complete")
	deleteTask := flag.Int("delete", 0, "Delete # task")

	flag.Parse()

	//	Show info/credits and usage about this program after its execution
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\nDesarrollado por Sigfrid Alex.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\nVersion 0.5\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\nUsage of %s\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Try to read the file. If not it is possible display an error and quit the app
	if err := list.Read(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *listTask:
		fmt.Print(list)
	case *isTaskCompleted > 0:
		if err := list.Complete(*isTaskCompleted); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := list.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	// BEWARE Can get multiple values if use '-task "task xxx xxx"'
	case *newTask != "":
		list.Add(*newTask)
		if err := list.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *deleteTask > 0:
		if err := list.Delete(*deleteTask); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		if err := list.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid flags (-task)(-list)(-complete), use -h for more details")
	}
}

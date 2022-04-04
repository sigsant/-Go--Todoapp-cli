package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sigsant/todo"
)

func main() {

	fmt.Print("\n\tIntroduce una tarea: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	todo.Add(text)

	for i, v := range todo.List {
		fmt.Printf("\n\t%d. %s", i, v)
	}
}

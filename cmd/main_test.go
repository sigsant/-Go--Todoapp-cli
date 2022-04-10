package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	filenameExecutable = "todo"
	dummyJSON          = "todo.json"
)

func TestMain(m *testing.M) {

	if runtime.GOOS == "windows" {
		filenameExecutable += ".exe"
	}

	fmt.Println("Creating file...")

	// Execute "go build" command
	build := exec.Command("go", "build", "-o", filenameExecutable)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error building %s: %s", filenameExecutable, err)
	}

	fmt.Println("Running file...")
	resultRunning := m.Run()

	os.Remove(filenameExecutable)
	os.Exit(resultRunning)
}

func TestCli(t *testing.T) {
	dummyTask := "Dummy Task"

	//Get the actual directory of the test
	filePath, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	// Create path with the bin file (...\todo\cmd\todo.exe)
	cmdPath := path.Join(filePath, filenameExecutable)

	//	Creating subtests with anonymous function...

	t.Run("Add New Tasks", func(t *testing.T) {
		// Execute bin file with variadic arguments
		cmd := exec.Command(cmdPath, strings.Split(dummyTask, " ")...)

		// Running test

		err := cmd.Run()

		assert.Nil(t, err, "Not possible to add any task")
	})
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	filenameExecutable = "todo"
	dummyJSON          = "task.json"
)

//	Similar to @beforeEach in Jest
func TestMain(m *testing.M) {

	if runtime.GOOS == "windows" {
		filenameExecutable += ".exe"
	}

	fmt.Println("\nCreating file...")

	// Execute "go build" command
	build := exec.Command("go", "build", "-o", filenameExecutable)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error building %s: %s", filenameExecutable, err)
	}

	fmt.Println("Running file...")
	resultRunning := m.Run()

	fmt.Println("Deleting file...")
	os.Remove(filenameExecutable)
	os.Remove(dummyJSON)
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
		cmd := exec.Command(cmdPath, "-task", dummyTask)

		// Running test

		err := cmd.Run()

		assert.Nil(t, err, "Unable to add any task")
	})

	t.Run("List task", func(t *testing.T) {

		expected := "\t[ ] 1: " + dummyTask + "\n"

		cmd := exec.Command(cmdPath, "-list")
		// Expect a return value
		taskOuput, err := cmd.CombinedOutput()

		assert.Nil(t, err, "Unable to list any task")
		assert.Equal(t, expected, string(taskOuput))
	})

	t.Run("Complete task", func(t *testing.T) {
		indexTask := strconv.Itoa(0)
		cmd := exec.Command(cmdPath, "-complete", indexTask)

		err := cmd.Run()

		assert.Nil(t, err, "Unable to change Completed as true")
		for _, v := range *list {
			assert.True(t, v.Completed)
		}
	})
}

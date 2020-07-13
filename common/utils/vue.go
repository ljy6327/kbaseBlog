package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func VueRun() error {

	//change directory
	err1 := os.Chdir("view")
	if err1 != nil {
		return err1
	}

	input := "npm run server"
	args := strings.Split(input, " ")

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and save it's output.
	err2 := cmd.Run()

	if err2 != nil {
		log.Panic(err2)
	}
	return nil
}

package main

import (
	"log"
	"os"
	"os/exec"
)

const ReturnCodeSuccess = 0
const ReturnCodeFail = 1

// RunCmd runs a command + arguments (cmd) with environment variables from env
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for k, v := range env {
		var err error

		if len(v) > 0 {
			err = os.Setenv(k, v)
		} else {
			err = os.Unsetenv(k)
		}

		if err != nil {
			log.Fatalf("Failed to set env vars: %s", err)

			return ReturnCodeFail
		}
	}

	var cmdName string
	var cmdArgs []string

	if len(cmd) == 0 {
		log.Fatal("Cmd should contain at least 1 argument")

		return ReturnCodeFail
	}

	cmdName = cmd[0]

	// nolint
	if len(cmd) > 1 {
		cmdArgs = cmd[1:]
	}

	c := exec.Command(cmdName, cmdArgs...)
	c.Env = os.Environ()
	c.Stdout = os.Stdout
	c.Stdin = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		log.Fatalf("Failed to execute command: %s", err)

		return ReturnCodeFail
	}

	return ReturnCodeSuccess
}

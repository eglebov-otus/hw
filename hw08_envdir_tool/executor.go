package main

import (
	"log"
	"os"
	"os/exec"
)

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

			return 1
		}
	}

	c := exec.Command(cmd[0], cmd[1:]...)
	c.Env = append(os.Environ())
	c.Stdout = os.Stdout
	c.Stdin = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		log.Fatalf("Failed to execute command: %s", err)

		return 1
	}

	return 0
}

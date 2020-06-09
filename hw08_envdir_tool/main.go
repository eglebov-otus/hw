package main

import "os"

func main() {
	env, err := ReadDir(os.Args[1])

	if err != nil {
		panic(err)
	}

	os.Exit(RunCmd(os.Args[2:], env))
}

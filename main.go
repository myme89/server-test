package main

import (
	"fmt"
	"os"
	"server-test/cmd"
)

var version string

func main() {
	fmt.Println(os.Args[1])
	if len(os.Args) == 2 {

		switch os.Args[1] {

		case "version":
			cmd.Version(version)
		case "help":
			cmd.Help()
		case "run":
			cmd.Run()
		default:
			cmd.Help()
		}
	} else {
		cmd.Help()
	}
}

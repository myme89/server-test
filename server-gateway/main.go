package main

import (
	"os"
	"server-test/server-gateway/cmd"
)

var version string

func main() {
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

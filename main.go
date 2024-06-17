package main

import (
	"os"
	"strings"

	"github.com/dimitardimitrov13/bbscli/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		if err, ok := err.(commands.BBSError); ok {
			os.Exit(err.ExitCode())
		}

		if strings.Contains(err.Error(), "invalid argument") {
			os.Exit(3)
		}

		os.Exit(-1)
	}
}

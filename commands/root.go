package commands

import (
	"errors"

	"code.cloudfoundry.org/lager/v3"
	"github.com/spf13/cobra"
)

var globalLogger = lager.NewLogger("bbs")

var RootCmd = &cobra.Command{
	Use:   "bbscli",
	Short: "BBS simple client",
	Long:  "A command-line tool to interact with a BBS deployment",
}

var (
	errMissingArguments   = errors.New("Missing arguments")
	errExtraArguments     = errors.New("Too many arguments specified")
	errInvalidProcessGuid = errors.New("Process guid should be non empty string")
)

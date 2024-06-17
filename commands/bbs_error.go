package commands

import (
	"code.cloudfoundry.org/bbs/models"
	"fmt"

	"github.com/spf13/cobra"
)

type BBSError struct {
	err      error
	exitCode int
}

func (a BBSError) Error() string {
	if err, ok := a.err.(*models.Error); ok {
		return fmt.Sprintf(`BBS error Type %d: %s Message: %s`, err.Type, err.Type.String(), err.Message)
	}

	return a.err.Error()
}

func (a BBSError) ExitCode() int {
	return a.exitCode
}

func NewBBSError(cmd *cobra.Command, err error) BBSError {
	cmd.SilenceUsage = true

	if _, ok := err.(*models.Error); ok {
		return BBSError{
			err:      err,
			exitCode: 4,
		}
	}

	return BBSError{
		err:      err,
		exitCode: 5,
	}
}

func NewBBSComponentError(cmd *cobra.Command, err error) BBSError {
	cmd.SilenceUsage = true

	return BBSError{
		err:      err,
		exitCode: 4,
	}
}

func NewBBSValidationError(cmd *cobra.Command, err error) BBSError {
	return BBSError{
		err:      err,
		exitCode: 3,
	}
}

package stscreds

import (
	"os"
	"fmt"
	"time"
)

type ExecCommand struct {
	Expiry time.Duration
	Profile string
}

func (cmd *ExecCommand) Execute() error {
	sess, err := newLimitedAccessSession(cmd.Profile)
	if err != nil {
		return err
	}
	username, err := currentUserName(sess)
	if err != nil {
		return fmt.Errorf("couldn't request current user: %s\n", err.Error())
	}

	fmt.Fprintf(os.Stderr, "Current user: %s. ", username)

	token, err := askUserForToken()
	if err != nil {
		return fmt.Errorf("error requesting mfa token: %s", err.Error())
	}

	_, err = requestNewSTSToken(sess, username, token, cmd.Expiry, cmd.Profile)
	if err != nil {
		return fmt.Errorf("error requesting credentials: %s", err.Error())
	}


	return nil
}
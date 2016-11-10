package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
	"os"
	"errors"
	"github.com/sonofbytes/stscreds"
)

var (
	_ = kingpin.Command("init", "Initialise stscreds. Creates ~/.stscreds/credentials.")
	expires     = kingpin.Flag("expires", "Credentials expiry").Default("1h").Duration()
	profile     = kingpin.Flag("profile", "AWS profile to manage credentials for.").Default("default").String()

	authCommand    = kingpin.Command("auth", "Authenticates with AWS and requests a temporary session token.")
	envVarTemplate = authCommand.Flag("output-env", "Additionally write environment variable exports to stdout.").Bool()

	readCommand = kingpin.Command("read", "Read keys from ~/.aws/credentials and print to stdout.")
	readKey     = readCommand.Arg("key", "Key to read from credentials file: aws_access_key_id, aws_secret_access_key, aws_session_token.").String()

	_ = kingpin.Command("whoami", "Print details about current user.")
)

var versionNumber string

func versionString() string {
	if versionNumber != "" {
		return versionNumber
	}
	return "DEVELOPMENT"
}

type Command interface {
	Execute() error
}

func cmdFailWithoutInitialisation(cmd Command) error {
	exist, err := stscreds.CredentialsExist()
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("no credentials found, please run init first.")
	}
	return cmd.Execute()
}

func handle() error {
	switch kingpin.Parse() {
	case "init":
		cmd := &stscreds.InitCommand{Profile: *profile}
		return cmd.Execute()
	case "whoami":
		return cmdFailWithoutInitialisation(&stscreds.WhoAmI{Profile: *profile})
	case "auth":
		return cmdFailWithoutInitialisation(&stscreds.AuthCommand{Expiry: *expires, OutputAsEnvVariable: *envVarTemplate, Profile: *profile})
	case "exec":
		return cmdFailWithoutInitialisation(&stscreds.ExecCommand{Expiry: *expires, Profile: *profile})
	case "read":
		return cmdFailWithoutInitialisation(&stscreds.ReadCommand{Key: *readKey, Expiry: *expires, Profile: *profile})
	}
	return nil
}

func main() {
	kingpin.Version(versionString())
	err := handle()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(2)
	}
}

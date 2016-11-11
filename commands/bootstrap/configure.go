package bootstrap

import (
	"github.com/alecthomas/kingpin"
	"fmt"
)

type BootstrapCommand struct {
	profile string
}

func Configure(app *kingpin.Application, bootstrapInput *BootstrapInput) {

	// See if the defaults get overridden
	if bootstrapInput == nil {
		bootstrapInput = &BootstrapInput{}
	}
	bootstrapInput.defaults()

	c := BootstrapCommand{}
	bc := app.Command("bootstrap", "Use root account to bootstrap into IAM Admin account").
		Action(c.run)
	bc.Flag("profile","AWS profile to manage credentials for.").
		Default("default").
		StringVar(&c.profile)

}

func (bc *BootstrapCommand)run(c *kingpin.ParseContext) error {
	fmt.Printf("HELLO %s\n", bc.profile)
	return nil
}
package bootstrap_test

import (
	"testing"
	"github.com/alecthomas/kingpin"
	"github.com/sonofbytes/stscreds/commands/bootstrap"
)

func TestConfigure(t *testing.T) {
	app := kingpin.New("test", "test app")
	bootstrap.Configure(app, nil)
}
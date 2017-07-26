package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
)

const (
	appname = "enrysrv"
)

var (
	version = "undefined"
	build   = "undefined"
	commit  = "undefined"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	parser := flags.NewNamedParser(appname, flags.Default)
	parser.AddCommand("server", "", "Run server", &serverCmd{})
	parser.AddCommand("client", "", "Run client", &clientCmd{})

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				os.Exit(0)
			} else {
				parser.WriteHelp(os.Stderr)
				os.Exit(1)
			}
		}

		logrus.Errorf("exiting with error: %s", err)
		os.Exit(1)
	}

	logrus.Debug("exiting without error")
}

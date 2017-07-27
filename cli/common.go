package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-errors.v0"
)

var (
	ErrMaxMessageSizeTooBig = errors.NewKind("max-message-size too big (limit is 2047MB): %u")
)

type commonCmd struct {
	LogLevel       string `long:"log-level" description:"log level" default:"debug"`
	MaxMessageSize string `long:"max-message-size" description:"maximum message size to send/receive to/from clients (in MB)" default:"100"`
	Version        bool   `short:"v" long:"version" description:"print version"`
}

var versionMessage = fmt.Sprintf("version:%s build:%s commit:%s", version, build, commit)

func (c *commonCmd) exec(args []string) error {
	if c.Version {
		fmt.Fprintf(os.Stdout, "%s\n", versionMessage)
		return nil
	}

	level, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)
	logrus.Info(versionMessage)
	return nil
}

func (c *commonCmd) parseMaxMessageSize() (int, error) {
	if c.MaxMessageSize == "" {
		return 0, nil
	}

	n, err := strconv.ParseUint(c.MaxMessageSize, 10, 16)
	if err != nil {
		return 0, err
	}

	if n >= 2048 {
		// Setting the hard limit of message size to less than 2GB since
		// it may overflow an int value, and it should be big enough
		return 0, ErrMaxMessageSizeTooBig.New(n)
	}

	return int(n * 1024 * 1024), nil // Convert MB to B
}

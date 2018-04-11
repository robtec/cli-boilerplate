package commands

import (
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:   "say-hello",
		Usage:  "Say hello command",
		Action: sayHello,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name, n",
				Value: "Jim",
			},
		},
	},
	{
		Name:   "config",
		Usage:  "test config",
		Action: testconfig,
	},
}

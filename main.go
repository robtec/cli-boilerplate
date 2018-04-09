package main

import (
	"fmt"
	"os"

	"github.com/robtec/cli-boilerplate/commands"
	"github.com/urfave/cli"
)

var (
	Version  = "0"
	CommitId = "0"
)

func main() {

	app := cli.NewApp()
	app.Name = "My Cli Example"
	app.Usage = "Demo of urfave Cli Sweetness!"

	app.Version = fmt.Sprintf("%s - %s", Version, CommitId)

	app.Commands = commands.Commands

	app.Run(os.Args)
}

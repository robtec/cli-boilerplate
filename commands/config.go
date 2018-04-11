package commands

import (
	"fmt"

	"github.com/robtec/cli-boilerplate/config"
	"github.com/urfave/cli"
)

type Data struct {
	Name    string
	Summary string
}

func testconfig(c *cli.Context) error {
	conf := config.NewConfig()
	conf.Save(Data{"54343", "ghlglg"})

	d, _ := conf.Load()

	fmt.Print(d)

	return nil
}

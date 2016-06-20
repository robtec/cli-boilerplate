package commands

import (
    "fmt"
    "github.com/urfave/cli"
)

func sayHello(c *cli.Context) error {
    fmt.Println("Hello",c.String("name"))
    return nil
}
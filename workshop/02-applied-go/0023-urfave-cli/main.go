package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Recipe DB",
		Usage: "recipe-db [command]",
		Action: func(c *cli.Context) error {
			fmt.Println("No recipe command provided")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

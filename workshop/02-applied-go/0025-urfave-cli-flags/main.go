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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "search",
				Usage: "search for a recipe",
				Value: "",
			},
		},
		Action: func(c *cli.Context) error {
			searchParam := c.String("search")
			if searchParam != "" {
				fmt.Println("Searching for recipe:", searchParam)
			} else {
				fmt.Println("No recipe command provided")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

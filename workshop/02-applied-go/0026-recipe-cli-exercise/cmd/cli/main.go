package main

import (
	"fmt"
	"goschool/0026-recipe-cli-exercise/api"
	"goschool/0026-recipe-cli-exercise/file"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	/// setup cli
	app := &cli.App{
		Name: "Recipe DB",
		Commands: []*cli.Command{
			{
				Name:  "search",
				Usage: "search for a recipe",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "save",
						Usage: "save recipe to file",
						Value: false,
					},
				},
				Action: func(c *cli.Context) error {
					searchParam := c.Args().First()
					if searchParam == "" {
						fmt.Println("No recipe name provided")
						return fmt.Errorf("no recipe name provided")
					} else {
						fmt.Println("Searching for recipe:", searchParam)
					}

					// TODO: implement search for recipe
					meals, err := api.SearchMealsByName(searchParam)
					if err != nil {
						fmt.Println("Error searching for meals:", err)
						return err
					}

					fmt.Println("Found meals:", meals)

					if c.Bool("save") {
						// TODO: save recipe to file
						// use SaveToFile from file package
					}

					return nil
				},
			},
			{
				Name:  "random",
				Usage: "get a random recipe",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "save",
						Usage: "save recipe to file",
						Value: false,
					},
				},
				Action: func(c *cli.Context) error {
					meal, err := api.GetRandomMeal()

					if err != nil {
						fmt.Println("Error getting random meal:", err)
						return err
					}

					fmt.Printf("Random meal:\nID:%s Name: %s\n", meal.IDMeal, meal.StrMeal)

					if c.Bool("save") {
						err := file.SaveToFile(meal.IDMeal+".txt", meal)
						if err != nil {
							fmt.Println("Error saving random meal:", err)
							return err
						}
					}

					return nil
				},
			},
			{
				Name:  "id",
				Usage: "get a recipe by id",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "save",
						Usage: "save recipe to file",
						Value: false,
					},
				},
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					if id == "" {
						fmt.Println("No recipe id provided")
						return fmt.Errorf("no recipe id provided")
					} else {
						fmt.Println("Searching for recipe by id:", id)
					}

					// TODO: implement GetMealById for recipe
					meal, err := api.GetMealById(id)
					if err != nil {
						fmt.Println("Error getting meal by id:", err)
					}
					fmt.Printf("Random meal:\nID:%s Name: %s\n", meal.IDMeal, meal.StrMeal)

					if c.Bool("save") {
						// TODO: save recipe to file
						// use SaveToFile from file package
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

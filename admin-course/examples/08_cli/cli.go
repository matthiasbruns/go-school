package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

//START_1 OMIT
func main() {
	app := &cli.App{ // HL1
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error { // HL1
			fmt.Println("boom! I say!")
			return nil
		},
	}

	err := app.Run(os.Args) // HL1
	if err != nil {
		log.Fatal(err)
	}
}

//END_1 OMIT

//START_2 OMIT
func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			fmt.Printf("Hello %q", c.Args().Get(0)) // HL2
			return nil
		},
	}

	err := app.Run(os.Args) // HL2
	if err != nil {
		log.Fatal(err)
	}
}

//END_2 OMIT

func main() {
	//START_3 OMIT

	app := &cli.App{
		Flags: []cli.Flag{ // HL3
			&cli.StringFlag{
				Name:  "lang",
				Value: "english", // HL3
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			name := "Matthias"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "spanish" { // HL3
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
		//END_3 OMIT
	}

}

//START_4 OMIT
var language string // HL4

app := &cli.App{
  Flags: []cli.Flag {
	&cli.StringFlag{
	  Name:        "lang",
	  Value:       "english",
	  Usage:       "language for the greeting",
	  Destination: &language, // HL4
	},
  },
  Action: func(c *cli.Context) error {
 	// ...
	if language == "spanish" { // HL4
	  fmt.Println("Hola", name)
	} else {
	  fmt.Println("Hello", name)
	}
	return nil
  },
}
// ...
//END_4 OMIT


//START_5 OMIT
Flags: []cli.Flag {
	&cli.StringFlag{
	  Name:    "lang",
	  Aliases: []string{"l"},
	  Value:   "english",
	  Usage:   "language for the greeting",
	  EnvVars: []string{"APP_LANG"}, // HL5
	},
  },
}
// ...
//END_5 OMIT

//START_6 OMIT
app.Flags = []cli.Flag {
    &cli.StringFlag{
      Name: "password",
      Aliases: []string{"p"},
      Usage: "password for the mysql database",
      FilePath: "/etc/mysql/password",  // HL6
    },
  }
// ...
//END_6 OMIT

//START_7 OMIT
app := &cli.App{
    Commands: []*cli.Command{ // HL7
      {
        Name:    "add",
        Aliases: []string{"a"},
        Usage:   "add a task to the list",
        Action:  func(c *cli.Context) error {
          fmt.Println("added task: ", c.Args().First())
          return nil
        },
      },
	  // More commands
    },
  }
//END_7 OMIT
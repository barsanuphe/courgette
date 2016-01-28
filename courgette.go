// Courgette deals with pictures
package main

import (
	"fmt"
	"os"

	"github.com/barsanuphe/courgette/courgette"
	"github.com/codegangsta/cli"
)

func main() {
	fmt.Printf("\n# # # C O U R G E T T E # # #\n\n")
	c := courgette.Collection{}
	if err := c.Load("courgette"); err != nil {
		fmt.Println("Could not load configuration")
		os.Exit(-1)
	}
	if err := c.Check(); err != nil {
		fmt.Println("Invalid Configuration")
		os.Exit(-1)
	}

	app := cli.NewApp()
	app.Name = "C O U R G E T T E"
	app.Usage = "Organize your photo collection."
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "import",
			Aliases: []string{"i"},
			Usage:   "import from card reader.",
			Action: func(c *cli.Context) {
				// TODO
			},
		},
		{
			Name:    "export",
			Aliases: []string{"x"},
			Usage:   "export [tag] to [folder].",
			Action: func(c *cli.Context) {
				// TODO
			},
		},
		{
			Name:    "collection",
			Aliases: []string{"c"},
			Usage:   "options for photo collection",
			Subcommands: []cli.Command{
				{
					Name:  "findorphans",
					Usage: "find orphan raw files in a subdirectory.",
					Action: func(c *cli.Context) {
						// TODO
					},
				},
				{
					Name:  "sortnew",
					Usage: "sort newly imported pictures.",
					Action: func(c *cli.Context) {
						// TODO
					},
				},
				{
					Name:  "refresh",
					Usage: "refresh filenames in a subdirectory.",
					Action: func(c *cli.Context) {
						// TODO
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

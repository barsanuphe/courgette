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
			Name:        "import",
			Aliases:     []string{"i"},
			Usage:       "import pictures from card reader or sort incoming",
			ArgsUsage:   "[DISK_NAME]",
			Description: "import from card reader if DISK_NAME is provided, or incoming directory.",
			Action: func(c *cli.Context) {
				// TODO if args, import from disk with that name
				// TODO if not, or after copying from disk, sort files in incoming
			},
		},
		{
			Name:        "export",
			Aliases:     []string{"x"},
			Usage:       "export pictures with a given tag to a folder.",
			ArgsUsage:   "TAG FOLDER",
			Description: "export pictures with a TAG to a local directory.",
			Action: func(c *cli.Context) {
				// TODO
			},
		},
		{
			Name:        "findorphans",
			Aliases:     []string{"f"},
			Usage:       "find raw files without jpg versions, to be removed.",
			ArgsUsage:   "[SUBDIR]",
			Description: "Raw pictures without jpg versions are flagged for deletion.\n   If SUBDIR is given, only this subdirectory is considered instead of the whole collection.",
			Action: func(c *cli.Context) {
				// TODO
			},
		},
		{
			Name:        "refresh",
			Aliases:     []string{"r"},
			ArgsUsage:   "[SUBDIR]",
			Usage:       "refresh filenames in a subdirectory.",
			Description: "Rename pictures according to configuration and metadata, after import.\n   If SUBDIR is given, only this subdirectory is considered instead of the whole collection.",
			Action: func(c *cli.Context) {
				// TODO
			},
		},
		{
			Name:        "diff",
			Aliases:     []string{"d"},
			ArgsUsage:   "[SUBDIR]",
			Usage:       "show changes relative to last commit.",
			Description: "Show changes since last commit.\n   If SUBDIR is given, only this subdirectory is considered instead of the whole collection.",
			Action: func(c *cli.Context) {
				// TODO
			},
		},
		{
			Name:        "commit",
			Aliases:     []string{"c"},
			ArgsUsage:   "[SUBDIR]",
			Usage:       "accept changes in the collection.",
			Description: "Accept changes in the collection and save the current state.\n   If SUBDIR is given, only this subdirectory is considered instead of the whole collection.",
			Action: func(c *cli.Context) {
				// TODO save hash and info
			},
		},
	}

	app.Run(os.Args)
}

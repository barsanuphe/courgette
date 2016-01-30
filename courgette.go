// Courgette deals with pictures
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/barsanuphe/courgette/courgette"
	"github.com/codegangsta/cli"
)

func importPictures(cc courgette.Collection, cardName string) (err error) {
	if cardName != "" {
		// check it's a valid external mount point.
		completeCardPath, err := cc.CheckValidCardReader(cardName)
		if os.IsNotExist(err) {
			fmt.Println("Card '" + cardName + "' not found. Is it mounted?")
			return err
		} else if err != nil {
			return err
		}
		numImported, err := cc.Import(completeCardPath)
		if err != nil {
			return err
		}
		fmt.Println("Imported " + strconv.Itoa(numImported) + " pictures.")
	}
	numSorted, err := cc.SortNew()
	if err != nil {
		return err
	}
	fmt.Println("Sorted " + strconv.Itoa(numSorted) + " pictures.")
	return
}

func main() {
	fmt.Printf("\n# # # C O U R G E T T E # # #\n\n")
	cc := courgette.Collection{}
	if err := cc.Load("courgette"); err != nil {
		fmt.Println("Could not load configuration")
		os.Exit(-1)
	}

	if err := cc.Check(); err != nil {
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
				// check input
				if len(c.Args()) > 1 {
					fmt.Printf("Too many arguments. See usage below.\n\n")
					cli.ShowCommandHelp(c, "import")
					os.Exit(-1)
				}
				// do things.
				err := importPictures(cc, c.Args().First())
				if err != nil {
					if !os.IsNotExist(err) {
						fmt.Println("ERR: " + err.Error() + "\nStopping everything.")
					}
					os.Exit(-1)
				}
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

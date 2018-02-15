package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

const version = "0.1"

func main() {
	app := cli.NewApp()
	app.Version = version
	app.EnableBashCompletion = true
	app.Name = "gmig"
	app.Usage = "GCP migrations tool"
	app.Commands = []cli.Command{
		{
			Name:   "new",
			Usage:  "gmig create \"create tester service account\"",
			Action: cmdCreateMigration,
		},
		{
			Name:   "up",
			Usage:  "gmig up",
			Action: cmdMigrationsUp,
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func hanAction(c *cli.Context) {
	fmt.Println("added task han: ", c)
}

func main() {

	// global level flags
	flagz := []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show more output",
		},
		cli.StringFlag{
			Name:  "e, errfile",
			Usage: "Specify an alternate err file (default: errfile.go)",
		},
		cli.StringFlag{
			Name:  "p, project-name",
			Usage: "Specify an alternate project name (default: directory name)",
		},
	}

	// Commands
	cmdz := []cli.Command{
		{
			Name: "han",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Handle errors in place",
			Action: hanAction,
		},
		{
			Name: "del",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Handle errors in errfile",
			Action: hanAction,
		},
		{
			Name: "init",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Create an empty errfile",
			Action: hanAction,
		},
		{
			Name: "ls",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "List unhandled",
			Action: hanAction,
		},
	}

	app := cli.NewApp()
	app.Flags = flagz
	app.Commands = cmdz

	fmt.Println("hello world")

	app.Run(os.Args)
}

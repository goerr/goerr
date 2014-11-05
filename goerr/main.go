package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/davecgh/go-spew/spew"
	//	"go/ast"
	"go/parser"
	"go/token"
	//	"io/ioutil"
	"os"
)

func i() {
	spew.Dump(0)
}

func massageAction(c *cli.Context) {
	fmt.Println("massaging the codebase")
}

func hanAction(c *cli.Context) {
	fmt.Println("added task han: ", c.Command.Flags)

	//	spew.Dump(c.globalSet)

	codefile := c.GlobalString("f")

	errfile := c.GlobalString("e")
	/*
		errf, _ := os.Open(errfile)
		defer errf.Close()

		codf, _ := os.Open(codefile)
		defer codf.Close()

		codeall, _ := ioutil.ReadAll(codf)
		errall, _ := ioutil.ReadAll(errf)
	*/
	fsete := token.NewFileSet()
	fsetc := token.NewFileSet()
	fc, _ := parser.ParseFile(fsetc, codefile, nil, 0)
	fe, _ := parser.ParseFile(fsete, errfile, nil, 0)

	spew.Dump(fc)
	spew.Dump(fe)
}

func delAction(c *cli.Context) {
	fmt.Println("added task del: ", c)
}

func missingAction(c *cli.Context) {
	//TODO
	fmt.Println("TODO :)")
}

func main() {

	// global level flags
	flagz := []cli.Flag{
		cli.StringFlag{
			Name: "e, " + error_file_name,
			Usage: "Specify an alternate " + error_file_name +
				" (default: " + error_file_name + ".go)",
		},
		cli.StringFlag{
			Name:  "f, file-name",
			Usage: "Operate on a single file",
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
			Action: missingAction,
		},
		{
			Name: "ls",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "u",
					Usage: "Unhandled errors only.",
				},
			},
			Usage:  "List the handling of all errors",
			Action: missingAction,
		},
		{
			Name: "show",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Show handling related to a function",
			Action: missingAction,
		},
		{
			Name: "status",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Show cumulative statistics",
			Action: missingAction,
		},
	}

	app := cli.NewApp()
	app.Flags = flagz
	app.Commands = cmdz
	app.Usage = "strip / add error handling to a go file"
	app.Version = app_version
	app.Action = massageAction

	app.Run(os.Args)
}

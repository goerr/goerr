package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/parser"
	"go/token"
	//	"io/ioutil"
	"os"
	//	"strings"
)

func i() {
	spew.Dump(0)
}

func massageAction(c *cli.Context) {
	fmt.Println("massaging the codebase")
}

type myfileinfo struct {
	is_goerr_import  bool
	is_dotted_import bool
}

type spewlord struct{}

func (spewlord) Visit(node ast.Node) ast.Visitor {

	var fun *ast.SelectorExpr
	var funx *ast.Ident
	var funsel *ast.Ident

	switch n := node.(type) {
	case *ast.CallExpr:

		switch fun := n.Fun.(type) {
		case *ast.SelectorExpr:
			/*
				switch funsel := fun.Sel.(type) {
				case *ast.Ident:
			*/

			switch funsel := interface{}(fun.Sel).(type) {
			case *ast.Ident:

				switch funx := fun.X.(type) {
				case *ast.Ident:

					_ = fun
					_ = funx
					_ = funsel

					spew.Dump(n, fun, funx, funsel)

					//				}
				}
			}
		}

	}

	_ = fun
	_ = funx
	_ = funsel

	/*
		if funx.Name == "goerr" {
			if funsel.Name == "XQZ" {
				fmt.Println("found recover wrapper")
			}
			if funsel.Name[:2] == "OR" {
				fmt.Println("found return wrapper")
			}

		}
	*/

	/*

		m := n

	*/

	/*
		var efns map[string]bool

		efns["errB"] = true
		efns["errA"] = true
	*/
	//	if efns[node.Ident] {
	//		fmt.Println("found cool node:")

	//	}

	return spewlord{}
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

	var myc, mye myfileinfo

	for _, s := range fc.Imports {
		if s.Path.Value == "\"github.com/goerr/goerr\"" {
			myc.is_goerr_import = true

			if s.Name != nil && s.Name.Name == "." {
				myc.is_dotted_import = true
			}

		}
	}

	_ = myc
	_ = mye
	_ = fc
	_ = fe
	for _, s := range fc.Decls {
		//		spewlord{}.Visit(s)

		ast.Walk(spewlord{}, s)
	}
	//	spew.Dump(fc.Imports)
	//	spew.Dump(fc.Decls)
	//	spew.Dump(fe.Decls)
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

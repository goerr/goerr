package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	//	"io/ioutil"
	"os"
	//	"strings"
)

func i() {
	spew.Dump(0)
	printer.Fprint(os.Stdout, nil, nil)
}

func massageAction(c *cli.Context) {
	fmt.Println("massaging the codebase")
}

type myfileinfo struct {
	is_goerr_import  bool
	is_dotted_import bool
}

type spewlord struct {
	f      func(string) int
	bodies []*ast.BlockStmt
}

func wesit(node ast.Node, f func(string) int) (rrr []*ast.CallExpr, bbb *ast.BlockStmt, offs []int, idz []int, er error) {
	var e = fmt.Errorf("Incorrect block node")
	er = e

	switch ggg := node.(type) {
	case *ast.BlockStmt:
		bbb = ggg
		switch lll := interface{}(ggg.List).(type) {
		case []ast.Stmt:
			for sssid, sss := range lll {
				_ = sssid
				_ = sss

				switch nod := sss.(type) {
				case *ast.ExprStmt:
					//			spew.Dump(nod)
					switch foo := interface{}(nod.X).(type) {
					case *ast.CallExpr:

						switch funnam := foo.Fun.(type) {
						case *ast.Ident:
							valuee := f(funnam.Name)
							if debag == 7 {

								spew.Dump("hello", funnam.Name, valuee)
							}

							if valuee != 0 {

								er = nil
								offs = append(offs, sssid)
								idz = append(idz, valuee)
								rrr = append(rrr, foo)
							}

						}

					}
				}

			}
		}
	}
	return
}

func witch(node ast.Node) (*ast.CallExpr, *ast.Ident, *ast.Ident, error) {
	var e = fmt.Errorf("Incorrect call node")

	switch n := node.(type) {
	case *ast.CallExpr:

		switch fun := n.Fun.(type) {
		case *ast.SelectorExpr:
			switch funsel := interface{}(fun.Sel).(type) {
			case *ast.Ident:

				switch funx := fun.X.(type) {
				case *ast.Ident:
					return n, funx, funsel, nil
				}

			}
		case *ast.Ident:
			return n, nil, fun, nil

		}
	}
	return nil, nil, nil, e
}

type errf struct {
	m      map[string]int
	bodies []*ast.BlockStmt
}

func (e *errf) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.FuncDecl:

		str := n.Name.String()
		//		spew.Dump(n.Body)
		e.bodies = append(e.bodies, n.Body)
		//		spew.Dump("FOUND:" + str)
		e.m[str] = len(e.bodies)

	default:
		return e
	}

	if debag == 4 {
		spew.Dump(node)
	}
	return e
}

func (s *spewlord) Visit(node ast.Node) ast.Visitor {

	var rewriter bool

	nnn, funx, funsel, err := witch(node)
	if err != nil {
		rewriter = true
	}
	rrr, bufflist, offz, idz, err2 := wesit(node, s.f)
	if err2 != nil {
		if rewriter {
			return s
		}
	}

	_ = rrr

	if rewriter {
		baff := &(bufflist.List)

		// first put together the statement "a = 42"
		identA := ast.NewIdent("a")
		fortyTwo := &ast.BasicLit{Kind: token.INT, Value: "42"}
		assignment := &ast.AssignStmt{Lhs: []ast.Expr{identA}, Rhs: []ast.Expr{fortyTwo}}
		nothing := []ast.Stmt(nil)
		well := []ast.Stmt{assignment}
		_ = well
		something := &ast.IfStmt{Body: &ast.BlockStmt{Lbrace: 398, List: nothing, Rbrace: 402}}
		_ = something
		_ = assignment
		var put []ast.Stmt

		spew.Dump("offz:$", len(offz))

		for i := range offz {

			toput := ((*s).bodies)[idz[i]-1]
			if debag == 1 {
				spew.Dump(toput)
				spew.Dump("**********$")
			}
			_ = i
			put = append(put, toput)

			var nargs *ast.CallExpr
			nargs = rrr[i].Args[0].(*ast.CallExpr)

			if debag == 1 {
				spew.Dump(rrr[i])
				spew.Dump("$$$$$$$$$$$")
				spew.Dump(nargs)
				spew.Dump("$********$")
			}

			rrr[i].Fun = nargs.Fun
			rrr[i].Args = nargs.Args
			rrr[i].Ellipsis = nargs.Ellipsis
			rrr[i].Lparen = nargs.Lparen
			rrr[i].Rparen = nargs.Rparen

			if debag == 1 {
				spew.Dump(rrr[i])
				spew.Dump("@@@@@@@@@@@@@@")

			}
		}

		sliceshift(baff, offz, put)
		/*

		*/
		if debag == 1 {
			for i := range offz {
				_ = i
				//		spew.Dump((*baff)[i])
			}
			spew.Dump("hello")
			spew.Dump(offz)
			spew.Dump(idz)
			//			spew.Dump(baff)
		}
		return s
	}

	if funx != nil && funx.Name == "goerr" {
		if funsel.Name == "XQZ" {
			nnn.Args = nil
			nnn.Fun = ast.NewIdent("recover")
			return s
		}
		if funsel.Name[:2] == "OR" {
			if len(nnn.Args) != 0 {
				nnn.Fun = nnn.Args[0]
				nnn.Args = nnn.Args[1:]
			}
			return s
		}
	}

	return s
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

	eh := errf{m: make(map[string]int)}

	for _, s := range fe.Decls {
		ast.Walk(&eh, s)
	}

	funny := func(s string) int { return eh.m[s] }

	for _, s := range fc.Decls {
		ast.Walk(&spewlord{f: funny, bodies: eh.bodies}, s)
	}

	if debag == 2 {
		printer.Fprint(os.Stdout, fsetc, fc)
	}
	//	spew.Dump(fc.Imports)
	if debag == 3 {
		spew.Dump(fc.Decls)
	}
	//	spew.Dump(fe.Decls)
}

func delAction(c *cli.Context) {
	fmt.Println("added task del: ", c)
}

func missingAction(c *cli.Context) {
	//TODO
	fmt.Println("TODO :)")
}

func sliceshift(baf *[]ast.Stmt, offs []int, put []ast.Stmt) {
	var out []ast.Stmt
	o := 0

	for i, j := range *baf {
		out = append(out, j)
		if len(offs) > 0 && (i+o) == offs[0] {
			out = append(out, put[0])
			offs = offs[1:]
			put = put[1:]
			o++
		}
	}

	*baf = out
}

func main() {
	/*
		if debag == 1337 {
			baf := []int{49868, 685498, 3218, 654, 6541, 6531, 486}
			put := []int{0, 1}
			where := []int{2, 4}
			sliceshift(&baf, where, put)
			spew.Dump(baf)
			os.Exit(0)
		}
	*/
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

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

type spewlord struct {
	f       func(string) int
	bodies  []*ast.BlockStmt
	eargtxt []*ast.Ident
	eargoff []int
	eargtot []int
}

type item struct {
	root *ast.Stmt
	rrr  *ast.CallExpr
	off  int
	idz  int
	lhs  *[]ast.Expr
	op   bool
}

func callmanage(op bool, baz *ast.Stmt, foo *ast.CallExpr, f func(string) int, st *[]item, sssid int, o *int, lhs *[]ast.Expr, e *error) {
	switch funnam := foo.Fun.(type) {
	case *ast.Ident:
		valuee := f(funnam.Name)

		if valuee != 0 {

			if debag == 31 {

				spew.Dump("PAT", funnam.Name, valuee, baz)
			}

			(*st) = append(*st, item{op: op, root: baz, off: sssid + *o, idz: valuee, rrr: foo, lhs: lhs})
			(*e) = nil
			(*o)++
		}

	}
}

func wesit(node ast.Node, f func(string) int) (st []item, bbb *ast.BlockStmt, er error) {
	er = fmt.Errorf("Incorrect block node")

	o := 0

	switch ggg := node.(type) {
	case *ast.BlockStmt:
		bbb = ggg

		switch lll := interface{}(ggg.List).(type) {
		case []ast.Stmt:
			/*
				if debag == 31 {
					spew.Dump("????????????????")
					spew.Dump(lll)
				}
			*/
			for sssid, sss := range lll {
				_ = sssid
				_ = sss

				switch nod := sss.(type) {
				case *ast.ExprStmt:
					/*
						if debag == 4 {
							spew.Dump("!!!!!!!!!!??")
							spew.Dump(sss)

						}
					*/
					switch foo := interface{}(nod.X).(type) {
					case *ast.CallExpr:
						callmanage(false, &lll[sssid], foo, f, &st, sssid, &o, nil, &er)

					}

				case *ast.AssignStmt:

					s := nod.Tok.String()

					lhs := &nod.Lhs

					if s != ":=" && s != "=" {
						spew.Dump("!operator?" + s)
					}

					op := s == "="

					switch foo := interface{}(nod.Rhs[0]).(type) {
					case *ast.CallExpr:
						callmanage(op, &lll[sssid], foo, f, &st, sssid, &o, lhs, &er)
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
	m       map[string]int
	bodies  []*ast.BlockStmt
	eargtxt []*ast.Ident
	eargoff []int
	eargtot []int
}

func (e *errf) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.FuncDecl:

		arglist := n.Type.Params.List

		errori := -1

		for i := range arglist {

			switch errargt := interface{}(arglist[i].Type).(type) {
			case *ast.Ident:
				if errargt.Name == "error" {
					//fmt.Fprintln(os.Stderr, "TODO arg type IS error")
					errori = i
					break
				}
			}
		}

		if errori == -1 {
			fmt.Fprintln(os.Stderr, "TODO no error arg", errori)
			return e
		}

		var nonstrerrp *ast.Ident

		switch errargn := interface{}(arglist[errori].Names).(type) {
		case []*ast.Ident:
			if len(errargn) != 1 {
				spew.Dump("TODO arg multiple names?")
				return e
			}
			nonstrerrp = errargn[0]
		}

		if debag == 6 {
			spew.Dump(nonstrerrp, errori, len(arglist))
		}

		e.eargtxt = append(e.eargtxt, nonstrerrp)
		e.eargoff = append(e.eargoff, errori)
		e.eargtot = append(e.eargtot, len(arglist))

		str := n.Name.String()
		//		spew.Dump(n.Body)
		e.bodies = append(e.bodies, n.Body)
		//		spew.Dump("FOUND:" + str)
		e.m[str] = len(e.bodies)

	case *ast.BlockStmt:

		for in, nod := range n.List {

			switch no := nod.(type) {
			case *ast.ReturnStmt:

				n.List = n.List[:in]
				return e

			case *ast.ExprStmt:
				switch call := no.X.(type) {
				case *ast.CallExpr:

					switch funx := call.Fun.(type) {
					case *ast.Ident:
						if funx.Name == "Return" {

							n.List[in] = &ast.ReturnStmt{Return: 0, Results: call.Args}

						}

					}
				}
			}

		}

	default:

		return e
	}

	return e
}

func (s *spewlord) Visit(node ast.Node) ast.Visitor {

	var rewriter bool

	nnn, funx, funsel, err := witch(node)
	if err != nil {
		rewriter = true
	}
	stek, bufflist, err2 := wesit(node, s.f)
	if err2 != nil {
		if rewriter {
			return s
		}
	}

	_ = stek

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

		put := []ast.Stmt {&ast.EmptyStmt{}}

		//		spew.Dump("offz:$", len(offz))

		offz := []int {0}

		varz := make(map[string]bool)

		for i := range stek {

			offz = append(offz, stek[i].off+1)

			toput := ((*s).bodies)[stek[i].idz-1]
			if debag == 18 {
				spew.Dump(toput)
				spew.Dump("**********$")
			}
			_ = i
			put = append(put, toput)

			var nargs *ast.CallExpr
			nargs = stek[i].rrr.Args[0].(*ast.CallExpr)

			if debag == 1 {
				spew.Dump(stek[i].rrr)
				spew.Dump("$$$$$$$$$$$")
				spew.Dump(nargs)
				spew.Dump("$********$")
			}

			stek[i].rrr.Fun = nargs.Fun
			stek[i].rrr.Args = nargs.Args
			stek[i].rrr.Ellipsis = nargs.Ellipsis
			stek[i].rrr.Lparen = nargs.Lparen
			stek[i].rrr.Rparen = nargs.Rparen

			if debag == 1 {
				spew.Dump(stek[i].rrr)
				spew.Dump("@@@@@@@@@@@@@@")

			}

			/*
				if stek[i].root != nil {



					*(stek[i].root) = nil
				}
			*/

			tput := ((*s).eargtxt)[stek[i].idz-1]
			puttoff := ((*s).eargoff)[stek[i].idz-1]
			puttot := ((*s).eargtot)[stek[i].idz-1]

			if stek[i].lhs != nil {

				if stek[i].op {
					spew.Dump("we must put vars because:")
					varz[tput.Name] = true
				}

				if debag == 4 {
					//					spew.Dump(stek[i].lhs)
					spew.Dump(tput, puttoff, puttot, "**********$")
				}

				_ = tput
				_ = puttoff

				argsliceshiftone(stek[i].lhs, puttoff, puttot, tput, ast.NewIdent("_"))
				/*
					if debag == 6 {
						spew.Dump(stek[i].lhs)
						spew.Dump(":D:D:**$")
					}
				*/

				if debag == 31 {
					spew.Dump("??????????????????????????/")
					spew.Dump(stek[i].root)

				}

			} else {

				if debag == 4 {
					spew.Dump("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
					//					spew.Dump(stek[i].root)
					//					spew.Dump(stek[i].rrr)

				}

				newlhs := []ast.Expr{}
				argsliceshiftone(&newlhs, puttoff, puttot, tput, ast.NewIdent("_"))

				assignment := ast.AssignStmt{
					Lhs:    newlhs,
					TokPos: 0,
					Tok:    token.DEFINE,
					Rhs:    []ast.Expr{stek[i].rrr}}

				*(stek[i].root) = &assignment

				if debag == 4 {
					spew.Dump("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
					//spew.Dump(stek[i].root)

				}
				/*

				*/
				/*
					if debag == 4 {
						spew.Dump(stek[i])
						spew.Dump("@@@@@@@@@@@@@@")

					}
				*/
			}
		}

		if len(varz) > 0 {
			idents := []*ast.Ident{}

			for item := range varz {
				idents = append(idents, ast.NewIdent(item))
			}

			//			rspec := ast.TypeSpec{Name: ast.NewIdent("error")}
			lspec := ast.ValueSpec{
				Names: idents,
				Type:  ast.NewIdent("error")}

			declaration := ast.GenDecl{
				Doc:    nil,
				TokPos: 0,
				Rparen: 0,
				Lparen: 0,
				Tok:    token.VAR,
				Specs:  []ast.Spec{&lspec}}

			_ = declaration

			smt := ast.DeclStmt{Decl: &declaration}

			hhh := []ast.Stmt{&smt}

			spew.Dump(hhh)

			for i := range offz {
				offz[i]++
			}

			offz[0] = 0
			put[0] = &smt

		} else {
			offz = offz[1:]
			put = put[1:]
		}

		sliceshift(baff, offz, put)
		/*

		*/
		if debag == 1 {
			for i := range offz {
				_ = i
				//		spew.Dump((*baff)[i])
			}
			//			spew.Dump("hello")
			//			spew.Dump(offz)
			//			spew.Dump(idz)
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
	//	fmt.Println("added task han: ", c.Command.Flags)

	//	spew.Dump(c.globalSet)

	codefile := c.GlobalString("f")

	errfile := c.GlobalString("e")

	need_use_stdout := c.GlobalString("o")

	use_stdout := codefile == "" && need_use_stdout != ""
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

	for _, s := range fc.Imports {
		if s.Path.Value == "\"github.com/goerr/goerr\"" {
			s.Path.Value = ""
		}
	}

	_ = fc
	_ = fe

	eh := errf{m: make(map[string]int)}

	for _, s := range fe.Decls {
		ast.Walk(&eh, s)
	}

	funny := func(s string) int { return eh.m[s] }

	for _, s := range fc.Decls {
		ast.Walk(&spewlord{f: funny,
			bodies:  eh.bodies,
			eargtxt: eh.eargtxt,
			eargoff: eh.eargoff,
			eargtot: eh.eargtot,
		}, s)
	}

	var outf *os.File

	if use_stdout {
		outf = os.Stdout
	} else {
		outf, _ = os.Create(codefile)
	}

	printer.Fprint(outf, fsetc, fc)
	outf.Sync()
	outf.Close()

	if debag == 2 {

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

func slicerm(baf *[]*ast.ImportSpec, n int) {
	end := len(*baf) - 1
	for i := range *baf {
		if i < n || i == end {
			continue
		}
		(*baf)[i] = (*baf)[i+1]
	}
	(*baf) = (*baf)[:end]
}

func argsliceshiftone(baf *[]ast.Expr, off int, tot int, put *ast.Ident, fill *ast.Ident) {
	var out []ast.Expr

	for i, j := range *baf {
		if i == off {
			out = append(out, put)
		}
		out = append(out, j)
	}

	for len(out) < tot {
		if len(out) == off {
			out = append(out, put)
		} else {
			out = append(out, fill)
		}
	}

	*baf = out
}

func sliceshift(baf *[]ast.Stmt, offs []int, put []ast.Stmt) {
	var out []ast.Stmt
	o := 0

	for i, j := range *baf {
		if len(offs) > 0 && (i+o) == offs[0] {
			out = append(out, put[0])
			offs = offs[1:]
			put = put[1:]
			o++
		}
		out = append(out, j)
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
			Name: "e",
			Usage: "Specify an alternate " + error_file_name +
				" (default: " + error_file_name + ".go)",
		},
		cli.StringFlag{
			Name:  "f",
			Usage: "Operate on a single file",
		},
		cli.StringFlag{
			Name:  "o",
			Usage: "Output to stdout. Use with f",
		},
	}

	// Commands
	cmdz := []cli.Command{
		{
			Name: "sep",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Separate errors handling to an errfile",
			Action: hanAction,
		},
		{
			Name: "merge",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "f",
					Usage: "Foo.",
				},
			},
			Usage:  "Merge handlers back",
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

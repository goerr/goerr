// this program contains a bug. the Println error is unhandled.
//--okey but that would be another example

package main

import (
	"fmt"
	"github.com/goerr/goerr"
)

var (
	ErrGopher = fmt.Errorf("Gopher failed")
)

func successer() error {
	return nil
}

func failer() error {
	return ErrGopher
}

func bar(i bool) (j string, k string) {
	fmt.Println("Branching")

	if i {
		errB(successer())
	} else {
		errB(failer())
	}
	return "hello", "world"
}

func foo() {

	errA(successer())

	fmt.Println("This will happen")

	defer func() {
		if r := goerr.XQZ(recover()); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()

	errA(failer())

	fmt.Println("This maybe happens")
}

func main() {
	goerr.OR0(foo)

	rts, str := goerr.OR2(bar, (false))

	fmt.Println("This surely happens", str, rts)
}

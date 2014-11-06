package main

import (
	"fmt"
	"github.com/goerr/goerr"
)

var (
	ErrGopher = fmt.Errorf("Gopher failed")
)

func pass0() (error, string, int) {
	return ErrGopher, "dog", 7
}

func pass1() (bool, error, int) {
	return true, ErrGopher, 9
}

func pass2() (bool, string, error) {
	return false, "cat", ErrGopher
}

func try() {
	a, b := errPASS0(pass0())
	p, q := errPASS1(pass1())
	x, y := errPASS2(pass2())

	fmt.Println(a,b,p,q,x,y)
}

func main() {
	goerr.OR0(try)
}

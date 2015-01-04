package main

import (
	"fmt"
	"github.com/anlhord/goerr"
)

var (
	ErrSome = fmt.Errorf("Error.")
	ErrException = fmt.Errorf("Exception.")
)

func successer() error {
	return nil
}

func failer() error {
	return ErrSome
}

func recursion1(i int) {
	if i > 0 {
		recursion2(i-1)
		return
	}
	errPropagate(successer())
}

func recursion2(i int) {
	if i > 0 {
		recursion1(i-2)
		return
	}
	errPropagate(failer())
}

func run(i int) error {
	if (i & 1) == 0 {
		recursion1(i/2)
	} else {
		recursion2(i/2)
	}
	return nil
}

func main() {
	a := goerr.OR1(run, 785)
	b := goerr.OR1(run, 578)
	c := goerr.OR1(run, 654)
	d := goerr.OR1(run, 457)

	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d)
}

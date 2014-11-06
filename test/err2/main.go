package main

import (
	"fmt"
	"github.com/goerr/goerr"
)

var (
	ErrSea = fmt.Errorf("Fish failed")
)

func fail0() (error, string, int) {
	return ErrSea, "tree", 4
}

func fail1() (bool, error, int) {
	return false, ErrSea, 8
}

func fail2() (bool, string, error) {
	return true, "list", ErrSea
}

func try() {
	var m, u, y string
	var n, v, q int
	var p, x, w bool

	u, _ = errFAIL0(fail0())
	_, v = errFAIL1(fail1())
	w, _ = errFAIL2(fail2())

	m, n = errFAIL0(fail0())
	p, q = errFAIL1(fail1())
	x, y = errFAIL2(fail2())

	fmt.Println(m, n, p, q, x, y, u, v, w)
}

func main() {
	goerr.OR0(try)
}

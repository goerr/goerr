package main

import (
	"fmt"
	. "github.com/goerr/goerr"
)

func errPASS0(a error, b string, c int) (string, int) {
	if a != nil {
		fmt.Println("ERR PASS0")
	}

	return b, c
	Return()
	return b, c
}

func errPASS1(a bool, b error, c int) (bool, int) {
	if b != nil {
		fmt.Println("ERR PASS1")
	}

	return a, c
}

func errPASS2(a bool, b string, c error) (bool, string) {
	if c != nil {
		fmt.Println("ERR PASS2")
	}

	return a, b
}

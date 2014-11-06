package main

import (
	"fmt"
	. "github.com/goerr/goerr"
)

func errFAIL0(a error, b string, c int) (string, int) {
	if a != nil {
		fmt.Println("ERR HERE0")
	}

	return b, c
	Return(1, 2, 3, 4)
	return b, c
}

func errFAIL1(a bool, b error, c int) (bool, int) {
	if b != nil {
		fmt.Println("ERR HERE1")
	}

	return a, c
}

func errFAIL2(a bool, b string, c error) (bool, string) {
	if c != nil {
		fmt.Println("ERR HERE2")
	}

	return a, b
}

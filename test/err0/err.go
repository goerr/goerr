package main

import (
	"fmt"
	. "github.com/goerr/goerr"
)

func errA(err error) {
	if err != nil {
		fmt.Println("Oh very")
//		Return()
	}
}

func errB(err error) {
	if err != nil {
		fmt.Println("Oh snap")
		Return("gopher","turtle")
	}
}

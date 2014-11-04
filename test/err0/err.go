package main

import (
	"fmt"
)

func errA(err error) {
	fmt.Println("Oh very")
//	err.Return()
}

func errB(err error) {
	fmt.Println("Oh snap")
	err.Return("gopher","turtle")
}

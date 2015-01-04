package main

import (
	"fmt"
	. "github.com/anlhord/goerr"
)

func errPropagate(err error) {
	if err != nil {
		fmt.Println("Throwing")
		Return(ErrException)
	}
}

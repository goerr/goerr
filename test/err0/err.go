package main

import "fmt"

func err(err error) (r ret) {
	if err != nil {
		fmt.Println("Oh very")
		return true
	}
	return
}

////////////////////////////////////////////////////////////////////////////////

type ret bool

func def() {
	if r := recover(); r != nil {
		if r != "ar86e7a6rh" {
			panic(r)
		}
	}
}

func (r ret) ret() {
	if r {
		panic("ar86e7a6rh")
	}
}

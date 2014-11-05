package main

import (
	"fmt"
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
		successer()
		{
			if err != nil {
				fmt.
					Println("Oh snap")
				Return("gopher", "turtle")
			}
		}

		{
		}
	} else {
		failer()
		{
			if err != nil {
				fmt.
					Println("Oh snap")
				Return("gopher", "turtle")
			}
		}

	}
	return "hello", "world"
}

func foo() {

	successer()
	{
		if err != nil {
			fmt.
				Println("Oh very")
		}
	}

	fmt.Println("This will happen")
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	successer()
	{
		if err != nil {
			fmt.
				Println("Oh very")
		}
	}

	fmt.Println("This will happen")
	{
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()

	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}

	failer()
	{
		if err != nil {
			fmt.
				Println("Oh very")
		}
	}

	fmt.Println("This maybe happens")
}

func main() {
	foo()

	rts, str := bar((false))

	fmt.Println("This surely happens", str, rts)
}

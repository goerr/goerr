// this program contains a bug. the Println error is unhandled.
//--okey but that would be another example

package main

import "fmt"

var (
	ErrGopher = fmt.Errorf("Gopher failed")
)

func successer() error {
	return nil
}

func failer() error {
	panic("yeah")

	return ErrGopher
}

/*
//Original foo with an unhandled error

func foo() {
	failer()

	fmt.Println("This maybe happens")
}
*/

/*
//Handled foo
func foo() {
	err := failer()
	if err != nil {
		fmt.Println("Oh very")
		return
	}

	fmt.Println("This maybe happens")
}


func main() {
	foo()
	fmt.Println("This surely happens")
}
*/

//deleted bar
func bar(i bool) (j string, k string) {
	fmt.Println("Branching")

	if i {
		errB(successer())
	} else {
		errB(failer())
	}
	return "hello","world"
}

//Deleted foo
func foo() {

	errA(successer())

	fmt.Println("This will happen")

	defer func() {
		if r := X(recover()); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()

	errA(failer())

	fmt.Println("This maybe happens")
}

func main() {
	OR0(foo)

	rts, str := OR2(bar,(true))

	fmt.Println("This surely happens", str, rts)
}

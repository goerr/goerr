// this program contains a bug. the Println error is unhandled.
//--okey but that would be another example

package main

import "fmt"

var (
	ErrGopher = fmt.Errorf("Gopher failed")
)

func failer() error {
	return ErrGopher
}

/*
//Original foo with an unhandled error

func foo() {
	failer()

	fmt.Println("This maybe happens")
}
*/


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

/*
//Deleted foo
func foo() {
	defer def()

	err(failer()).ret()

	fmt.Println("This maybe happens")
}

func main() {
	foo()
	fmt.Println("This surely happens")
}
*/

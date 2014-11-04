package main

import (
	"fmt"
	"reflect"
)

func i(i interface{}) {
	i = reflect.Kind(0)
}

func errA(err error) {
	fmt.Println("Oh very")
	returns()
}

func errB(err error) {
	fmt.Println("Oh snap")
	returns("very", "bad")
}

////////////////////////////////////////////////////////////////////////////////
/*
added only to functions that do defer
added only before statements that can return in an error
*/
// if i, he works like deferr

type ptype struct{
	ret interface{}
}

type err1337 interface{}

func recovererr(a ...interface{}) (r interface{}) {
	fmt.Println("HELLO")
	r = recover()

	switch r.(type) {
	case ptype:
	fmt.Println("set args here")
		return nil
	}
	if (len(a) == 0) {
		return r
	}
	panic(r)
}

func returns(a ...interface{}) {
	fmt.Println("okay we returns", a)

	panic(ptype{a})
}

func err2(fun err1337, arguments ...err1337) (err1337, err1337) {
	e := errvariadic(fun, arguments)
	return e[0], e[1]
}

func err0(fun err1337) {
	errvariadic(fun)
}

//this is for functions with no return values
func errvariadic(fun err1337, arguments ...err1337) (rets []err1337) {
	defer recovererr(len(arguments))

	var vals []reflect.Value

	for _, arg := range arguments {
		vals = append(vals, reflect.ValueOf(arg))
	}

	outs := reflect.ValueOf(fun).Call(vals)

	fmt.Println("outs are:", outs)

	for _, out := range outs {
		rets = append(rets, out)
	}
	return
}

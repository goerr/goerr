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
//	Return()
}

func errB(err error) {
	fmt.Println("Oh snap")
	Return("gopher","turtle")
}

////////////////////////////////////////////////////////////////////////////////
/*
added only to functions that do defer
added only before statements that can return in an error
*/
// if i, he works like deferr

// The function X wraps the recover() call in client code
// When your code is in the hidden-error mode, you must wrap every recover
// in this function X, in order to avoid your code handling the goerr panics
func X(r interface{}) (interface{}) {
	switch r.(type) {
	case panik:
	return nil
	default:
	return r
	}
}

// The function Return is called from the error handling functions
func Return(a ...interface{}) {
	panic(panik{0xDEAD, a})
}

// took from github.com/go-on/queue
func toValues(in []interface{}) []reflect.Value {
	out := make([]reflect.Value, len(in))
	for i := range in {
		if in[i] != nil {
			out[i] = reflect.ValueOf(in[i])
		} else {
			out[i] = reflect.ValueOf(&in[i]).Elem()
		}
	}
	return out
}

// took from github.com/go-on/queue
func toInterfaces(in []reflect.Value) []interface{} {
	out := make([]interface{}, len(in))
	for i, vl := range in {
		out[i] = vl.Interface()
	}
	return out
}

// reduce verbosity
type q interface{}

// the return- panic struct
type panik struct {
	magic uint16
	q []interface{}
}

// Returner for 2 return-valued functions
func OR2(fun q, args ...interface{}) (q,q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1]
}

// Returner for 0 return-valued functions
// Returners should wrap in the client's code every function call, that contains
// hidden-error mode error handling
func OR0(fun q) {
	errvariadic(fun, []reflect.Value{})
}

// internal variadic returner
func errvariadic(fun interface{}, vals []reflect.Value) (out []interface{}) {
	defer func(){
		r := recover()
		switch r.(type) {
		case panik:
		p := r.(panik)
		out = p.q
		default:
		if r != nil {
			panic(r)
		}
		}
	}()
	return toInterfaces(reflect.ValueOf(fun).Call(vals))
}

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
	Return()
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

func recovererr(a ...*[]interface{}) (r interface{}) {
	r = recover()
	if len(a) != 0 {
	switch r.(type) {
	case []interface{}:
	fmt.Println("got panictype :")
		for _, ptr := range a {
			(*ptr) = r.([]interface{})
		}

		return nil
	default:
		panic(r)
	}}
	return r
}

func Return(a ...interface{}) {
	panic(a)
}

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

// toValues is a helper function that creates and returns a slice of
// interface{} values based on a given slice of reflect.Value values
func toInterfaces(in []reflect.Value) []interface{} {
	out := make([]interface{}, len(in))
	for i, vl := range in {
		//		if vl.IsNil() {
		//		out[i] = nil
		//} else {
		out[i] = vl.Interface()
		//}
	}
	return out
}

func OR2(fun interface{}, args ...interface{}) (interface{},interface{}) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1]
}

func OR0(fun interface{}) {
	errvariadic(fun, []reflect.Value{})
}

func errvariadic(fun interface{}, vals []reflect.Value) (out []interface{}) {
	defer recovererr(&out)
	return toInterfaces(reflect.ValueOf(fun).Call(vals))
}

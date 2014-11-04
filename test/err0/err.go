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
	returns("gopher","turtle")
}

////////////////////////////////////////////////////////////////////////////////
/*
added only to functions that do defer
added only before statements that can return in an error
*/
// if i, he works like deferr


func recovererr(a ...interface{}) (r interface{}) {
	if r = recover(); r == "ar86e7a6rh" {
		return nil
	}
	if (len(a) == 1) && r != nil {
		panic(r)
	}
	return r
}

func returns(a ...interface{}) {
	panic("ar86e7a6rh")
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

func err2(fun interface{}, args ...interface{}) (interface{},interface{}) {
	o := errvariadic(fun, toValues(args))

	fmt.Println("got ", o, "#")
	return nil, nil
//	return o[0], o[1]
}

func err0(fun interface{}) {
	errvariadic(fun, []reflect.Value{})
}

func errvariadic(fun interface{}, vals []reflect.Value) ([]interface{}) {
	defer recovererr(len(vals))
	return toInterfaces(reflect.ValueOf(fun).Call(vals))
}

package goerr

import (
	"reflect"
)

//XQZ is used in the clients code to wrap all calls recover to
//avoid recovering goerr internal panic by the client's code, for example:
//XQZ(recover());.
//Elsewhere, it doesn't do anything, you can delete it
func XQZ(r interface{}) interface{} {
	switch r.(type) {
	case panik:
		return nil
	default:
		return r
	}
}

// Return is called from the error handling functions
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
	q     []interface{}
}

// OR9 is a returner for 9 return-valued functions
func OR9(fun q, args ...interface{}) (q, q, q, q, q, q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3], o[4], o[5], o[6], o[7], o[8]
}

// OR8 is a returner for 8 return-valued functions
func OR8(fun q, args ...interface{}) (q, q, q, q, q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3], o[4], o[5], o[6], o[7]
}

// OR7 is a returner for 7 return-valued functions
func OR7(fun q, args ...interface{}) (q, q, q, q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3], o[4], o[5], o[6]
}

// OR6 is a returner for 6 return-valued functions
func OR6(fun q, args ...interface{}) (q, q, q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3], o[4], o[5]
}

// OR5 is a returner for 5 return-valued functions
func OR5(fun q, args ...interface{}) (q, q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3], o[4]
}

// OR4 is a returner for 4 return-valued functions
func OR4(fun q, args ...interface{}) (q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3]
}

// OR3 is a returner for 3 return-valued functions
func OR3(fun q, args ...interface{}) (q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2]
}

// OR2 is a returner for 2 return-valued functions
func OR2(fun q, args ...interface{}) (q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1]
}

// OR1 is a returner for 1 return-valued functions
func OR1(fun q, args ...interface{}) q {
	o := errvariadic(fun, toValues(args))
	return o[0]
}

//OR0 is a returner for 0 return-valued functions.
//Every function that contains
//wrappers who call Return(N), must be wrapped in a matching returner
//to obtain N return values. Otherwise, you get a panic from the Return.
func OR0(fun q, args ...interface{}) {
	errvariadic(fun, toValues(args))
}

// internal variadic returner
func errvariadic(fun interface{}, vals []reflect.Value) (out []interface{}) {
	defer func() {
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

package goerr

import (
	"reflect"
)

//The function RecoWrap is used in the clients code to wrap all calls recover to
//avoid recovering goerr internal panic by the client's code, for example:
//RecoWrap(recover());.
//Elsewhere, it doesn't do anything, you can delete it
func RecoWrap(r interface{}) interface{} {
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
	q     []interface{}
}

// Returner for 9 return-valued functions
func OR9(fun q, args ...interface{}) (q, q, q, q, q, q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3], o[4], o[5], o[6], o[7], o[8]
}

// Returner for 8 return-valued functions
func OR8(fun q, args ...interface{}) (q, q, q, q, q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3], o[4], o[5], o[6], o[7]
}

// Returner for 7 return-valued functions
func OR7(fun q, args ...interface{}) (q, q, q, q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3], o[4], o[5], o[6]
}

// Returner for 6 return-valued functions
func OR6(fun q, args ...interface{}) (q, q, q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3], o[4], o[5]
}

// Returner for 5 return-valued functions
func OR5(fun q, args ...interface{}) (q, q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3], o[4]
}

// Returner for 4 return-valued functions
func OR4(fun q, args ...interface{}) (q, q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2], o[3]
}

// Returner for 3 return-valued functions
func OR3(fun q, args ...interface{}) (q, q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1], o[2]
}

// Returner for 2 return-valued functions
func OR2(fun q, args ...interface{}) (q, q) {
	o := errvariadic(fun, toValues(args))
	return o[0], o[1]
}

// Returner for 1 return-valued functions
func OR1(fun q, args ...interface{}) q {
	o := errvariadic(fun, toValues(args))
	return o[0]
}

//Returner for 0 return-valued functions.
//Every function that contains
//wrappers who call Return(N), must be wrapped in a matching returner
//to obtain N return values. Otherwise, you get a panic from the Return.
func OR0(fun q) {
	errvariadic(fun, []reflect.Value{})
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

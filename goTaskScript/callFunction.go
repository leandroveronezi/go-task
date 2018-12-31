package goTaskScript

import (
	"errors"
	"reflect"
)

func callFunc(m map[int]TaskFunction, idx int, params ...interface{}) (result []reflect.Value, err error) {

	f := reflect.ValueOf(m[idx].Fun)
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.Call(in)
	return

}

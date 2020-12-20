package main

import (
	"errors"
	"reflect"
)

// snailSort gets a 2D quadratic slice, sorts its elements in snail order and
// returns them as a 1D slice, panics in case of error
func snailSort(in interface{}) (out []interface{}) {
	const (
		typeError        = "type error"
		elementTypeError = "element type error"
		dimensionError   = "dimension error"
	)

	handleError := func(errorType, a, op, b interface{}) {
		if (op == "ne" && a != b) || (op == "eq" && a == b) {
			panic(errors.New(errorType.(string)))
		}
	}

	inValue := reflect.ValueOf(in)
	inLen := inValue.Len()

	getElementType := func(i int) reflect.Kind {
		return reflect.TypeOf(inValue.Index(i).Interface()).Kind()
	}

	getElementLen := func(i int) int {
		return reflect.ValueOf(inValue.Index(i).Interface()).Len()
	}

	handleError(typeError, reflect.TypeOf(in).Kind(), "ne", reflect.Slice)
	handleError(typeError, inLen, "eq", 0)
	handleError(elementTypeError, getElementType(0), "ne", reflect.Slice)
	handleError(dimensionError, getElementLen(0), "ne", inLen)

	top, left, bottom, right := 0, 0, inLen, getElementLen(0)
	out = make([]interface{}, 0, inLen*inLen)
	firstIteration := true

	getElement := func(y, x int) interface{} {
		return reflect.ValueOf(inValue.Index(y).Interface()).Index(x).Interface()
	}

	sliceTop := func() {
		for i := left; i < right; i++ {
			out = append(out, getElement(top, i))
		}
		top++
	}

	sliceRight := func() {
		for i := top; i < bottom; i++ {
			if firstIteration {
				handleError(elementTypeError, getElementType(i), "ne", reflect.Slice)
				handleError(dimensionError, getElementLen(i), "ne", inLen)
			}
			out = append(out, getElement(i, right-1))
		}
		firstIteration = false
		right--
	}

	sliceBottom := func() {
		for i := right - 1; i >= left; i-- {
			out = append(out, getElement(bottom-1, i))
		}
		bottom--
	}

	sliceLeft := func() {
		for i := bottom - 1; i >= top; i-- {
			out = append(out, getElement(i, left))
		}
		left++
	}

	for top < bottom {
		sliceTop()
		if bottom-top == 0 {
			return
		}
		sliceRight()
		sliceBottom()
		sliceLeft()
	}
	return
}

package arrutil

import (
	"fmt"
	"reflect"
)

// FindFn  Function to find element in array
type FindFn func(a interface{}) bool

// Find takes a slice of type T, a FindFn and returns the the first T for which the FindFn returns true, or an error. FindFn is a function that takes a T and returns a boolean. If src is not a slice, or is an empty slice, an corresponding error is returned. If no element is found that satisfies FindFn, an error is returned.
// AI: "Find" 函数接收一个类型为 T 的切片和一个 "FindFn" 函数作为参数，并返回第一个满足 "FindFn" 函数为真的元素的指针，或一个错误。"FindFn" 是一个接受 T 类型参数并返回布尔值的函数。如果 "src" 不是一个切片，或者是一个空切片，则会返回相应的错误。如果没有找到满足 "FindFn" 函数的元素，则会返回一个错误。
func Find[T interface{}](src []T, fn FindFn) (T, error) {
	sType := reflect.TypeOf(src)
	var t T
	if sType.Kind() != reflect.Slice {
		return t, FindErrorNotArray
	}

	val := reflect.ValueOf(src)
	if val.Len() == 0 {
		return t, FindErrorEmpty
	}

	for i := 0; i < val.Len(); i++ {
		s := val.Index(i).Interface()
		if fn(s) {
			fmt.Println(val.Type())
			t = val.Index(i).Interface().(T)
			return t, nil
		}
	}
	return t, FindErrorNotFound

}

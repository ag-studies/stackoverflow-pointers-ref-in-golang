package main

import (
	"C"
	"unsafe"
)
import "reflect"

func Function(s0, s1, s2 float64, N int) unsafe.Pointer {
	result := make([]float64, N)
	for i := 0; i < N; i++ {
		result[i] = (s0 + s1 + s2)
	}
	return unsafe.Pointer(&result)
}

//export ExportedFunction
func ExportedFunction(s0, s1, s2 float64, N int) uintptr {
	p := Function(s0, s1, s2, N)
	s := *(*[]float64)(p)
	return uintptr(unsafe.Pointer(&s[0]))
}

func main() {
	//test
	p := Function(1.0, 1.1, 1.2, 2)
	//convert pointer to slice
	a := *(*[]float64)(p)
	//print result
	println(reflect.TypeOf(a).Kind().String())
	println(a[0], a[1])
}

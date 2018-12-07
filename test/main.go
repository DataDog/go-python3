package main

/*
#cgo pkg-config: python3
#include "Python.h"
*/
import "C"

import (
	"fmt"
	python "github.com/DataDog/go-python3"
)

func main() {
	C.Py_Initialize()
	ret, err := python.PyRun_AnyFile("test.py")
	fmt.Printf("%d %v\n", ret, err)
}

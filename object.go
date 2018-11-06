package python3

/*
#cgo pkg-config: python3
#include "Python.h"
*/
import "C"
import (
	"fmt"
)

type PyObject C.PyObject

func (pyObject *PyObject) IncRef() {
	C.Py_IncRef((*C.PyObject)(pyObject))
}

func (pyObject *PyObject) DecRef() {
	C.Py_DecRef((*C.PyObject)(pyObject))
}

func (pyObject *PyObject) ReprEnter() (bool, error) {
	if ret := C.Py_ReprEnter((*C.PyObject)(pyObject)); ret < 0 {
		return false, fmt.Errorf("recursion limit is reached")
	} else if ret > 0 {
		return true, nil
	}
	return false, nil
}

func (pyObject *PyObject) ReprLeave() {
	C.Py_ReprLeave((*C.PyObject)(pyObject))
}

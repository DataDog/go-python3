package python3

/*
#cgo pkg-config: python3
#include "Python.h"
*/
import "C"

type PyObject C.PyObject

func (pyObject *PyObject) IncRef() {
	C.Py_IncRef((*C.PyObject)(pyObject))
}

func (pyObject *PyObject) DecRef() {
	C.Py_DecRef((*C.PyObject)(pyObject))
}

package python3

/*
#cgo pkg-config: python3
#include "Python.h"
*/
import "C"
import (
	"fmt"
)

//PyObject : https://docs.python.org/3/c-api/structures.html?highlight=pyobject#c.PyObject
type PyObject C.PyObject

//togo converts a *C.PyObject to a *PyObject
func togo(cobject *C.PyObject) *PyObject {
	return (*PyObject)(cobject)
}

func toc(object *PyObject) *C.PyObject {
	return (*C.PyObject)(object)
}

//IncRef : https://docs.python.org/3/c-api/refcounting.html#c.Py_INCREF
func (pyObject *PyObject) IncRef() {
	C.Py_IncRef(toc(pyObject))
}

//DecRef : https://docs.python.org/3/c-api/refcounting.html#c.Py_DECREF
func (pyObject *PyObject) DecRef() {
	C.Py_DecRef(toc(pyObject))
}

//ReprEnter : https://docs.python.org/3/c-api/exceptions.html#c.Py_ReprEnter
func (pyObject *PyObject) ReprEnter() (bool, error) {
	if ret := C.Py_ReprEnter(toc(pyObject)); ret < 0 {
		return false, fmt.Errorf("recursion limit is reached")
	} else if ret > 0 {
		return true, nil
	}
	return false, nil
}

//ReprLeave : https://docs.python.org/3/c-api/exceptions.html#c.Py_ReprLeave
func (pyObject *PyObject) ReprLeave() {
	C.Py_ReprLeave(toc(pyObject))
}

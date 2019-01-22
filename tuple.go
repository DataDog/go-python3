/*
Unless explicitly stated otherwise all files in this repository are licensed
under the $license_for_repo License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

//Tuple : https://docs.python.org/3/c-api/tuple.html#c.PyTuple_Type
var Tuple = togo((*C.PyObject)(unsafe.Pointer(&C.PyTuple_Type)))

//PyTuple_Check : https://docs.python.org/3/c-api/tuple.html#c.PyTuple_Check
func PyTuple_Check(p *PyObject) bool {
	return C._go_PyTuple_Check(toc(p)) != 0
}

//PyTuple_CheckExact : https://docs.python.org/3/c-api/tuple.html#c.PyTuple_CheckExact
func PyTuple_CheckExact(p *PyObject) bool {
	return C._go_PyTuple_CheckExact(toc(p)) != 0
}

//PyTuple_New : https://docs.python.org/3/c-api/tuple.html#c.PyTuple_New
func PyTuple_New(len int) *PyObject {
	return togo(C.PyTuple_New(C.Py_ssize_t(len)))
}

//PyTuple_Size : https://docs.python.org/3/c-api/tuple.html#c.PyTuple_Size
func PyTuple_Size(p *PyObject) int {
	return int(C.PyTuple_Size(toc(p)))
}

//PyTuple_GetItem : https://docs.python.org/3/c-api/tuple.html#c.PyTuple_GetItem
func PyTuple_GetItem(p *PyObject, pos int) *PyObject {
	return togo(C.PyTuple_GetItem(toc(p), C.Py_ssize_t(pos)))
}

//PyTuple_GetSlice : https://docs.python.org/3/c-api/tuple.html#c.PyTuple_GetSlice
func PyTuple_GetSlice(p *PyObject, low, high int) *PyObject {
	return togo(C.PyTuple_GetSlice(toc(p), C.Py_ssize_t(low), C.Py_ssize_t(high)))
}

//PyTuple_SetItem : https://docs.python.org/3/c-api/tuple.html#c.PyTuple_SetItem
func PyTuple_SetItem(p *PyObject, pos int, o *PyObject) int {
	return int(C.PyTuple_SetItem(toc(p), C.Py_ssize_t(pos), toc(o)))
}

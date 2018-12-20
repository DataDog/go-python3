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

//List : https://docs.python.org/3/c-api/list.html#c.PyList_Type
var List = togo((*C.PyObject)(unsafe.Pointer(&C.PyList_Type)))

//PyList_Check : https://docs.python.org/3/c-api/list.html#c.PyList_Check
func PyList_Check(p *PyObject) bool {
	return C._go_PyList_Check(toc(p)) != 0
}

//PyList_CheckExact : https://docs.python.org/3/c-api/list.html#c.PyList_CheckExact
func PyList_CheckExact(p *PyObject) bool {
	return C._go_PyList_CheckExact(toc(p)) != 0
}

//PyList_New : https://docs.python.org/3/c-api/list.html#c.PyList_New
func PyList_New(len int) *PyObject {
	return togo(C.PyList_New(C.Py_ssize_t(len)))
}

//PyList_Size : https://docs.python.org/3/c-api/list.html#c.PyList_Size
func PyList_Size(p *PyObject) int {
	return int(C.PyList_Size(toc(p)))
}

//PyList_GetItem : https://docs.python.org/3/c-api/list.html#c.PyList_GetItem
func PyList_GetItem(p *PyObject, pos int) *PyObject {
	return togo(C.PyList_GetItem(toc(p), C.Py_ssize_t(pos)))
}

//PyList_SetItem : https://docs.python.org/3/c-api/list.html#c.PyList_SetItem
func PyList_SetItem(p *PyObject, pos int, o *PyObject) int {
	return int(C.PyList_SetItem(toc(p), C.Py_ssize_t(pos), toc(o)))
}

//PyList_Insert : https://docs.python.org/3/c-api/list.html#c.PyList_Insert
func PyList_Insert(p *PyObject, index int, item *PyObject) int {
	return int(C.PyList_Insert(toc(p), C.Py_ssize_t(index), toc(item)))
}

//PyList_Append : https://docs.python.org/3/c-api/list.html#c.PyList_Append
func PyList_Append(p, item *PyObject) int {
	return int(C.PyList_Append(toc(p), toc(item)))
}

//PyList_GetSlice : https://docs.python.org/3/c-api/list.html#c.PyList_GetSlice
func PyList_GetSlice(p *PyObject, low, high int) *PyObject {
	return togo(C.PyList_GetSlice(toc(p), C.Py_ssize_t(low), C.Py_ssize_t(high)))
}

//PyList_SetSlice : https://docs.python.org/3/c-api/list.html#c.PyList_SetSlice
func PyList_SetSlice(p *PyObject, low, high int, itemlist *PyObject) int {
	return int(C.PyList_SetSlice(toc(p), C.Py_ssize_t(low), C.Py_ssize_t(high), toc(itemlist)))
}

//PyList_Sort : https://docs.python.org/3/c-api/list.html#c.PyList_Sort
func PyList_Sort(list *PyObject) int {
	return int(C.PyList_Sort(toc(list)))
}

//PyList_Reverse : https://docs.python.org/3/c-api/list.html#c.PyList_Reverse
func PyList_Reverse(list *PyObject) int {
	return int(C.PyList_Reverse(toc(list)))
}

//PyList_AsTuple : https://docs.python.org/3/c-api/list.html#c.PyList_AsTuple
func PyList_AsTuple(list *PyObject) *PyObject {
	return togo(C.PyList_AsTuple(toc(list)))
}

//PyList_ClearFreeList : https://docs.python.org/3/c-api/list.html#c.PyList_ClearFreeList
func PyList_ClearFreeList() int {
	return int(C.PyList_ClearFreeList())
}

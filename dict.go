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
import (
	"unsafe"
)

//Dict : https://docs.python.org/3/c-api/dict.html#c.PyDict_Type
var Dict = togo((*C.PyObject)(unsafe.Pointer(&C.PyDict_Type)))

//PyDict_Check : https://docs.python.org/3/c-api/dict.html#c.PyDict_Check
func PyDict_Check(p *PyObject) bool {
	return C._go_PyDict_Check(toc(p)) != 0
}

//PyDict_CheckExact : https://docs.python.org/3/c-api/dict.html#c.PyDict_CheckExact
func PyDict_CheckExact(p *PyObject) bool {
	return C._go_PyDict_CheckExact(toc(p)) != 0
}

//PyDict_New : https://docs.python.org/3/c-api/dict.html#c.PyDict_New
func PyDict_New() *PyObject {
	return togo(C.PyDict_New())
}

//PyDictProxy_New : https://docs.python.org/3/c-api/dict.html#c.PyDictProxy_New
func PyDictProxy_New(mapping *PyObject) *PyObject {
	return togo(C.PyDictProxy_New(toc(mapping)))
}

//PyDict_Clear : https://docs.python.org/3/c-api/dict.html#c.PyDict_Clear
func PyDict_Clear(p *PyObject) {
	C.PyDict_Clear(toc(p))
}

//PyDict_Contains : https://docs.python.org/3/c-api/dict.html#c.PyDict_Contains
func PyDict_Contains(p, key *PyObject) int {
	return int(C.PyDict_Contains(toc(p), toc(key)))
}

//PyDict_Copy : https://docs.python.org/3/c-api/dict.html#c.PyDict_Copy
func PyDict_Copy(p *PyObject) *PyObject {
	return togo(C.PyDict_Copy(toc(p)))
}

//PyDict_SetItem : https://docs.python.org/3/c-api/dict.html#c.PyDict_SetItem
func PyDict_SetItem(p, key, val *PyObject) int {
	return int(C.PyDict_SetItem(toc(p), toc(key), toc(val)))
}

//PyDict_SetItemString : https://docs.python.org/3/c-api/dict.html#c.PyDict_SetItemString
func PyDict_SetItemString(p *PyObject, key string, val *PyObject) int {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	return int(C.PyDict_SetItemString(toc(p), ckey, toc(val)))
}

//PyDict_DelItem : https://docs.python.org/3/c-api/dict.html#c.PyDict_DelItem
func PyDict_DelItem(p, key *PyObject) int {
	return int(C.PyDict_DelItem(toc(p), toc(key)))
}

//PyDict_DelItemString : https://docs.python.org/3/c-api/dict.html#c.PyDict_DelItemString
func PyDict_DelItemString(p *PyObject, key string) int {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))
	return int(C.PyDict_DelItemString(toc(p), ckey))
}

//PyDict_GetItem : https://docs.python.org/3/c-api/dict.html#c.PyDict_GetItem
func PyDict_GetItem(p, key *PyObject) *PyObject {
	return togo(C.PyDict_GetItem(toc(p), toc(key)))
}

//PyDict_GetItemWithError : https://docs.python.org/3/c-api/dict.html#c.PyDict_GetItemWithError
func PyDict_GetItemWithError(p, key *PyObject) *PyObject {
	return togo(C.PyDict_GetItemWithError(toc(p), toc(key)))
}

//PyDict_GetItemString : https://docs.python.org/3/c-api/dict.html#c.PyDict_GetItemString
func PyDict_GetItemString(p *PyObject, key string) *PyObject {
	ckey := C.CString(key)
	defer C.free(unsafe.Pointer(ckey))

	return togo(C.PyDict_GetItemString(toc(p), ckey))
}

//PyDict_SetDefault : https://docs.python.org/3/c-api/dict.html#c.PyDict_SetDefault
func PyDict_SetDefault(p, key, pyDefault *PyObject) *PyObject {
	return togo(C.PyDict_SetDefault(toc(p), toc(key), toc(pyDefault)))
}

//PyDict_Items : https://docs.python.org/3/c-api/dict.html#c.PyDict_Items
func PyDict_Items(p *PyObject) *PyObject {
	return togo(C.PyDict_Items(toc(p)))
}

//PyDict_Keys : https://docs.python.org/3/c-api/dict.html#c.PyDict_Keys
func PyDict_Keys(p *PyObject) *PyObject {
	return togo(C.PyDict_Keys(toc(p)))
}

//PyDict_Values : https://docs.python.org/3/c-api/dict.html#c.PyDict_Values
func PyDict_Values(p *PyObject) *PyObject {
	return togo(C.PyDict_Values(toc(p)))
}

//PyDict_Size : https://docs.python.org/3/c-api/dict.html#c.PyDict_Size
func PyDict_Size(p *PyObject) int {
	return int(C.PyDict_Size(toc(p)))
}

//PyDict_Next : https://docs.python.org/3/c-api/dict.html#c.PyDict_Next
func PyDict_Next(p *PyObject, ppos *int, pkey, pvalue **PyObject) bool {
	cpos := C.Py_ssize_t(*ppos)
	ckey := toc(*pkey)
	cvalue := toc(*pvalue)

	res := C.PyDict_Next(toc(p), &cpos, &ckey, &cvalue) != 0

	*ppos = int(cpos)
	*pkey = togo(ckey)
	*pvalue = togo(cvalue)

	return res
}

//PyDict_ClearFreeList : https://docs.python.org/3/c-api/dict.html#c.PyDict_ClearFreeList
func PyDict_ClearFreeList() int {
	return int(C.PyDict_ClearFreeList())
}

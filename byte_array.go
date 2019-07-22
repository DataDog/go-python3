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

//ByteArray : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Type
var ByteArray = togo((*C.PyObject)(unsafe.Pointer(&C.PyByteArray_Type)))

//PyByteArray_Check : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Check
func PyByteArray_Check(o *PyObject) bool {
	return C._go_PyByteArray_Check(toc(o)) != 0
}

//PyByteArray_CheckExact : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_CheckExact
func PyByteArray_CheckExact(o *PyObject) bool {
	return C._go_PyByteArray_CheckExact(toc(o)) != 0
}

//PyByteArray_FromObject : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_FromObject
func PyByteArray_FromObject(o *PyObject) *PyObject {
	return togo(C.PyByteArray_FromObject(toc(o)))
}

//PyByteArray_FromStringAndSize : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_FromStringAndSize
func PyByteArray_FromStringAndSize(str string) *PyObject {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return togo(C.PyByteArray_FromStringAndSize(cstr, C.Py_ssize_t(len(str))))
}

//PyByteArray_Concat : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Concat
func PyByteArray_Concat(a, b *PyObject) *PyObject {
	return togo(C.PyByteArray_Concat(toc(a), toc(b)))
}

//PyByteArray_Size : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Size
func PyByteArray_Size(o *PyObject) int {
	return int(C.PyByteArray_Size(toc(o)))
}

//PyByteArray_AsString : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_AsString
func PyByteArray_AsString(o *PyObject) string {
	return C.GoStringN(C.PyByteArray_AsString(toc(o)), C.int(C.PyByteArray_Size(toc(o))))
}

//PyByteArray_Resize : https://docs.python.org/3/c-api/bytearray.html#c.PyByteArray_Resize
func PyByteArray_Resize(bytearray *PyObject, len int) {
	C.PyByteArray_Resize(toc(bytearray), C.Py_ssize_t(len))
}

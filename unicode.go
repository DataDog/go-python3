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

//Unicode : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_Type
var Unicode = togo((*C.PyObject)(unsafe.Pointer(&C.PyUnicode_Type)))

//PyUnicode_Check : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_Check
func PyUnicode_Check(o *PyObject) bool {
	return C._go_PyUnicode_Check(toc(o)) != 0
}

//PyUnicode_CheckExact : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_CheckExact
func PyUnicode_CheckExact(o *PyObject) bool {
	return C._go_PyUnicode_CheckExact(toc(o)) != 0
}

//PyUnicode_New : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_New
func PyUnicode_New(size int, maxchar rune) *PyObject {
	return togo(C.PyUnicode_New(C.Py_ssize_t(size), C.Py_UCS4(maxchar)))
}

//PyUnicode_FromString : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_FromString
func PyUnicode_FromString(u string) *PyObject {
	cu := C.CString(u)
	defer C.free(unsafe.Pointer(cu))

	return togo(C.PyUnicode_FromString(cu))
}

//PyUnicode_FromEncodedObject : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_FromEncodedObject
func PyUnicode_FromEncodedObject(obj *PyObject, encoding, errors string) *PyObject {
	cencoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(cencoding))

	cerrors := C.CString(errors)
	defer C.free(unsafe.Pointer(cerrors))

	return togo(C.PyUnicode_FromEncodedObject(toc(obj), cencoding, cerrors))
}

//PyUnicode_GetLength : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_GetLength
func PyUnicode_GetLength(unicode *PyObject) int {
	return int(C.PyUnicode_GetLength(toc(unicode)))
}

//PyUnicode_CopyCharacters : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_CopyCharacters
func PyUnicode_CopyCharacters(to, from *PyObject, to_start, from_start, how_many int) int {
	return int(C.PyUnicode_CopyCharacters(toc(to), C.Py_ssize_t(to_start), toc(from), C.Py_ssize_t(from_start), C.Py_ssize_t(how_many)))

}

//PyUnicode_Fill : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_Fill
func PyUnicode_Fill(unicode *PyObject, start, length int, fill_char rune) int {
	return int(C.PyUnicode_Fill(toc(unicode), C.Py_ssize_t(start), C.Py_ssize_t(length), C.Py_UCS4(fill_char)))
}

//PyUnicode_WriteChar : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_WriteChar
func PyUnicode_WriteChar(unicode *PyObject, index int, character rune) int {
	return int(C.PyUnicode_WriteChar(toc(unicode), C.Py_ssize_t(index), C.Py_UCS4(character)))
}

//PyUnicode_ReadChar : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_ReadChar
func PyUnicode_ReadChar(unicode *PyObject, index int) rune {
	return rune(C.PyUnicode_ReadChar(toc(unicode), C.Py_ssize_t(index)))
}

//PyUnicode_Substring : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_Substring
func PyUnicode_Substring(unicode *PyObject, start, end int) *PyObject {
	return togo(C.PyUnicode_Substring(toc(unicode), C.Py_ssize_t(start), C.Py_ssize_t(end)))
}

//PyUnicode_AsUTF8 : https://docs.python.org/3/c-api/unicode.html#c.PyUnicode_AsUTF8
func PyUnicode_AsUTF8(unicode *PyObject) string {
	cutf8 := C.PyUnicode_AsUTF8(toc(unicode))
	return C.GoString(cutf8)
}

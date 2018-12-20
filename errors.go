/*
Unless explicitly stated otherwise all files in this repository are licensed
under the $license_for_repo License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

/*
#include "Python.h"
*/
import "C"

import (
	"unsafe"
)

//PyErr_Clear : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_Clear
func PyErr_Clear() {
	C.PyErr_Clear()
}

//PyErr_PrintEx : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_PrintEx
func PyErr_PrintEx(setSysLastVars bool) {
	if setSysLastVars {
		C.PyErr_PrintEx(1)
	} else {
		C.PyErr_PrintEx(0)
	}
}

//PyErr_Print : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_Print
func PyErr_Print() {
	C.PyErr_PrintEx(1)
}

//PyErr_WriteUnraisable : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WriteUnraisable
func PyErr_WriteUnraisable(obj *PyObject) {
	C.PyErr_WriteUnraisable(toc(obj))
}

//PyErr_SetString : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetString
func PyErr_SetString(pyType *PyObject, message string) {
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))

	C.PyErr_SetString(toc(pyType), cmessage)

}

//PyErr_SetObject : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetObject
func PyErr_SetObject(pyType, value *PyObject) {
	C.PyErr_SetObject(toc(pyType), toc(value))
}

//PyErr_SetNone : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetNone
func PyErr_SetNone(pyType *PyObject) {
	C.PyErr_SetNone(toc(pyType))
}

//PyErr_BadArgument : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_BadArgument
func PyErr_BadArgument() {
	C.PyErr_BadArgument()
}

//PyErr_NoMemory : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_NoMemory
func PyErr_NoMemory() *PyObject {
	return togo(C.PyErr_NoMemory())
}

//PyErr_SetImportErrorSubclass : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetImportErrorSubclass
func PyErr_SetImportErrorSubclass(msg, name, path, subclass *PyObject) *PyObject {
	return togo(C.PyErr_SetImportErrorSubclass(toc(msg), toc(name), toc(path), toc(subclass)))
}

//PyErr_SetImportError : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetImportError
func PyErr_SetImportError(msg, name, path *PyObject) *PyObject {
	return togo(C.PyErr_SetImportError(toc(msg), toc(name), toc(path)))
}

//PyErr_SyntaxLocationObject : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SyntaxLocationObject
func PyErr_SyntaxLocationObject(filename *PyObject, lineno, col_offset int) {
	C.PyErr_SyntaxLocationObject(toc(filename), C.int(lineno), C.int(col_offset))
}

//PyErr_SyntaxLocationEx : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SyntaxLocationEx
func PyErr_SyntaxLocationEx(filename string, lineno, col_offset int) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	C.PyErr_SyntaxLocationEx(cfilename, C.int(lineno), C.int(col_offset))
}

//PyErr_SyntaxLocation : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SyntaxLocation
func PyErr_SyntaxLocation(filename string, lineno int) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	C.PyErr_SyntaxLocation(cfilename, C.int(lineno))

}

//PyErr_BadInternalCall : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_BadInternalCall
func PyErr_BadInternalCall() {
	C.PyErr_BadInternalCall()
}

//PyErr_Occurred : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_Occurred
func PyErr_Occurred() *PyObject {
	return togo(C.PyErr_Occurred())
}

//PyErr_GivenExceptionMatches : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_GivenExceptionMatches
func PyErr_GivenExceptionMatches(given, exc *PyObject) bool {
	ret := C.PyErr_GivenExceptionMatches(toc(given), toc(exc))
	return ret == 1
}

//PyErr_ExceptionMatches : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_ExceptionMatches
func PyErr_ExceptionMatches(exc *PyObject) bool {
	ret := C.PyErr_ExceptionMatches(toc(exc))
	return ret == 1
}

//PyErr_Fetch : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_Fetch
func PyErr_Fetch() (*PyObject, *PyObject, *PyObject) {
	var pyType, value, traceback *C.PyObject
	C.PyErr_Fetch(&pyType, &value, &traceback)
	return togo(pyType), togo(value), togo(traceback)
}

//PyErr_Restore : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_Restore
func PyErr_Restore(pyType *PyObject, value *PyObject, traceback *PyObject) {
	C.PyErr_Restore(toc(pyType), toc(value), toc(traceback))
}

//PyErr_NormalizeException : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_NormalizeException
func PyErr_NormalizeException(exc, val, tb *PyObject) (*PyObject, *PyObject, *PyObject) {
	cexc := toc(exc)
	cval := toc(val)
	ctb := toc(tb)
	C.PyErr_NormalizeException(&cexc, &cval, &ctb)
	return togo(cexc), togo(cval), togo(ctb)
}

//PyErr_GetExcInfo : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_GetExcInfo
func PyErr_GetExcInfo() (*PyObject, *PyObject, *PyObject) {
	var pyType, value, traceback *C.PyObject
	C.PyErr_GetExcInfo(&pyType, &value, &traceback)
	return togo(pyType), togo(value), togo(traceback)
}

//PyErr_SetExcInfo : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetExcInfo
func PyErr_SetExcInfo(pyType *PyObject, value *PyObject, traceback *PyObject) {
	C.PyErr_SetExcInfo(toc(pyType), toc(value), toc(traceback))
}

//PyErr_CheckSignals : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_CheckSignals
func PyErr_CheckSignals() int {
	return int(C.PyErr_CheckSignals())
}

//PyErr_SetInterrupt : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_SetInterrupt
func PyErr_SetInterrupt() {
	C.PyErr_SetInterrupt()
}

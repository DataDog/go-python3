package python3

/*
#include "Python.h"
#include "errors.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func PyErr_Clear() {
	C.PyErr_Clear()
}

func PyErr_PrintEx(setSysLastVars bool) {
	if setSysLastVars {
		C.PyErr_PrintEx(1)
	} else {
		C.PyErr_PrintEx(0)
	}
}

func PyErr_Print() {
	C.PyErr_PrintEx(1)
}

func PyErr_WriteUnraisable(obj *PyObject) {
	C.PyErr_WriteUnraisable((*C.PyObject)(obj))
}

func PyErr_SetString(pyType *PyObject, message string) {
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))

	C.PyErr_SetString((*C.PyObject)(pyType), cmessage)

}

func PyErr_SetObject(pyType, value *PyObject) {
	C.PyErr_SetObject((*C.PyObject)(pyType), (*C.PyObject)(value))
}

func PyErr_SetNone(pyType *PyObject) {
	C.PyErr_SetNone((*C.PyObject)(pyType))
}

func PyErr_BadArgument() {
	C.PyErr_BadArgument()
}

func PyErr_NoMemory() *PyObject {
	return (*PyObject)(C.PyErr_NoMemory())
}

func PyErr_SetImportErrorSubclass(msg, name, path, subclass *PyObject) *PyObject {
	return (*PyObject)(C.PyErr_SetImportErrorSubclass((*C.PyObject)(msg), (*C.PyObject)(name), (*C.PyObject)(path), (*C.PyObject)(subclass)))
}

func PyErr_SetImportError(msg, name, path *PyObject) *PyObject {
	return (*PyObject)(C.PyErr_SetImportError((*C.PyObject)(msg), (*C.PyObject)(name), (*C.PyObject)(path)))
}

func PyErr_SyntaxLocationObject(filename *PyObject, lineno, col_offset int) {
	C.PyErr_SyntaxLocationObject((*C.PyObject)(filename), C.int(lineno), C.int(col_offset))
}

func PyErr_SyntaxLocationEx(filename string, lineno, col_offset int) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	C.PyErr_SyntaxLocationEx(cfilename, C.int(lineno), C.int(col_offset))
}

func PyErr_SyntaxLocation(filename string, lineno int) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	C.PyErr_SyntaxLocation(cfilename, C.int(lineno))

}

func PyErr_BadInternalCall() {
	C.PyErr_BadInternalCall()
}

func PyErr_Occurred() *PyObject {
	return (*PyObject)(C.PyErr_Occurred())
}

func PyErr_GivenExceptionMatches(given, exc *PyObject) bool {
	ret := C.PyErr_GivenExceptionMatches((*C.PyObject)(given), (*C.PyObject)(exc))
	return ret == 1
}

func PyErr_ExceptionMatches(exc *PyObject) bool {
	ret := C.PyErr_ExceptionMatches((*C.PyObject)(exc))
	return ret == 1
}

func PyErr_Fetch() (*PyObject, *PyObject, *PyObject) {
	var pyType, value, traceback *C.PyObject
	C.PyErr_Fetch(&pyType, &value, &traceback)
	return (*PyObject)(pyType), (*PyObject)(value), (*PyObject)(traceback)
}

func PyErr_Restore(pyType *PyObject, value *PyObject, traceback *PyObject) {
	C.PyErr_Restore((*C.PyObject)(pyType), (*C.PyObject)(value), (*C.PyObject)(traceback))
}

func PyErr_NormalizeException(exc, val, tb *PyObject) (*PyObject, *PyObject, *PyObject) {
	cexc := (*C.PyObject)(exc)
	cval := (*C.PyObject)(val)
	ctb := (*C.PyObject)(tb)
	C.PyErr_NormalizeException(&cexc, &cval, &ctb)
	return (*PyObject)(cexc), (*PyObject)(cval), (*PyObject)(ctb)
}

func PyErr_GetExcInfo() (*PyObject, *PyObject, *PyObject) {
	var pyType, value, traceback *C.PyObject
	C.PyErr_GetExcInfo(&pyType, &value, &traceback)
	return (*PyObject)(pyType), (*PyObject)(value), (*PyObject)(traceback)
}

func PyErr_SetExcInfo(pyType *PyObject, value *PyObject, traceback *PyObject) {
	C.PyErr_SetExcInfo((*C.PyObject)(pyType), (*C.PyObject)(value), (*C.PyObject)(traceback))
}

func PyErr_CheckSignals() error {
	if C.PyErr_CheckSignals() == -1 {
		return fmt.Errorf("Exception raised during signal handle")
	}
	return nil
}

func PyErr_SetInterrupt() {
	C.PyErr_SetInterrupt()
}

func PySignal_SetWakeupFd(fd uintptr) uintptr {
	return uintptr(C.PySignal_SetWakeupFd(C.int(fd)))
}

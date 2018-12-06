package python3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"

import (
	"unsafe"
)

//Py_EnterRecursiveCall : https://docs.python.org/3/c-api/exceptions.html#c.Py_EnterRecursiveCall
func Py_EnterRecursiveCall(where string) int {
	cwhere := C.CString(where)
	C.free(unsafe.Pointer(cwhere))

	return int(C._go_Py_EnterRecursiveCall(cwhere))
}

//Py_LeaveRecursiveCall : https://docs.python.org/3/c-api/exceptions.html#c.Py_LeaveRecursiveCall
func Py_LeaveRecursiveCall() {
	C._go_Py_LeaveRecursiveCall()
}

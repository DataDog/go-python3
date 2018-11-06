package python3

/*
#include "Python.h"
#include "recursion.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//Py_EnterRecursiveCall : https://docs.python.org/3/c-api/exceptions.html#c.Py_EnterRecursiveCall
func Py_EnterRecursiveCall(where string) error {
	cwhere := C.CString(where)
	C.free(unsafe.Pointer(cwhere))

	if C._Py_EnterRecursiveCall(cwhere) != 0 {
		return fmt.Errorf("An error has been set in the python interpreter")
	}
	return nil
}

//Py_LeaveRecursiveCall : https://docs.python.org/3/c-api/exceptions.html#c.Py_LeaveRecursiveCall
func Py_LeaveRecursiveCall() {
	C._Py_LeaveRecursiveCall()
}

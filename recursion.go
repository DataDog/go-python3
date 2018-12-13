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

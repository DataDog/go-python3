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

//python boolean constants
var (
	Py_False = togo(C.Py_False)
	Py_True  = togo(C.Py_True)
)

//Bool : https://docs.python.org/3/c-api/bool.html#c.PyBool_Type
var Bool = togo((*C.PyObject)(unsafe.Pointer(&C.PyBool_Type)))

//PyBool_Check : https://docs.python.org/3/c-api/bool.html#c.PyBool_Check
func PyBool_Check(o *PyObject) bool {
	return C._go_PyBool_Check(toc(o)) != 0
}

//PyBool_FromLong : https://docs.python.org/3/c-api/bool.html#c.PyBool_FromLong
func PyBool_FromLong(v int) *PyObject {
	return togo(C.PyBool_FromLong(C.long(v)))
}

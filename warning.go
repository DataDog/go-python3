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

/*
All standard Python warning categories are available as global variables whose names are PyExc_ followed by the Python exception name.
These have the type PyObject*; they are all class objects.
*/
var (
	PyExc_Warning                   = togo(C.PyExc_Warning)
	PyExc_BytesWarning              = togo(C.PyExc_BytesWarning)
	PyExc_DeprecationWarning        = togo(C.PyExc_DeprecationWarning)
	PyExc_FutureWarning             = togo(C.PyExc_FutureWarning)
	PyExc_ImportWarning             = togo(C.PyExc_ImportWarning)
	PyExc_PendingDeprecationWarning = togo(C.PyExc_PendingDeprecationWarning)
	PyExc_ResourceWarning           = togo(C.PyExc_ResourceWarning)
	PyExc_RuntimeWarning            = togo(C.PyExc_RuntimeWarning)
	PyExc_SyntaxWarning             = togo(C.PyExc_SyntaxWarning)
	PyExc_UnicodeWarning            = togo(C.PyExc_UnicodeWarning)
	PyExc_UserWarning               = togo(C.PyExc_UserWarning)
)

//PyErr_WarnEx : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WarnEx
func PyErr_WarnEx(category *PyObject, message string, stack_level int) int {
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))

	return int(C.PyErr_WarnEx(toc(category), cmessage, C.Py_ssize_t(stack_level)))
}

//PyErr_WarnExplicitObject : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WarnExplicitObject
func PyErr_WarnExplicitObject(category *PyObject, message *PyObject, filename *PyObject, lineno int, module *PyObject, registry *PyObject) int {
	return int(C.PyErr_WarnExplicitObject(toc(category), toc(message), toc(filename), C.int(lineno), toc(module), toc(registry)))
}

//PyErr_WarnExplicit : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WarnExplicit
func PyErr_WarnExplicit(category *PyObject, message string, filename string, lineno int, module string, registry *PyObject) int {
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	cmodule := C.CString(module)
	defer C.free(unsafe.Pointer(cmodule))

	return int(C.PyErr_WarnExplicit(toc(category), cmessage, cfilename, C.int(lineno), cmodule, toc(registry)))
}

package python3

/*
#include "Python.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

/*
All standard Python warning categories are available as global variables whose names are PyExc_ followed by the Python exception name.
These have the type PyObject*; they are all class objects.
*/
var (
	Warning                   = (*PyObject)(C.PyExc_Warning)
	BytesWarning              = (*PyObject)(C.PyExc_BytesWarning)
	DeprecationWarning        = (*PyObject)(C.PyExc_DeprecationWarning)
	FutureWarning             = (*PyObject)(C.PyExc_FutureWarning)
	ImportWarning             = (*PyObject)(C.PyExc_ImportWarning)
	PendingDeprecationWarning = (*PyObject)(C.PyExc_PendingDeprecationWarning)
	ResourceWarning           = (*PyObject)(C.PyExc_ResourceWarning)
	RuntimeWarning            = (*PyObject)(C.PyExc_RuntimeWarning)
	SyntaxWarning             = (*PyObject)(C.PyExc_SyntaxWarning)
	UnicodeWarning            = (*PyObject)(C.PyExc_UnicodeWarning)
	UserWarning               = (*PyObject)(C.PyExc_UserWarning)
)

//PyErr_WarnEx : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WarnEx
func PyErr_WarnEx(category *PyObject, message string, stack_level int) error {
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))

	ret := C.PyErr_WarnEx((*C.PyObject)(category), cmessage, C.Py_ssize_t(stack_level))

	if ret == -1 {
		return fmt.Errorf("An exception was raised during warning execution")
	}
	return nil
}

//PyErr_WarnExplicitObject : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WarnExplicitObject
func PyErr_WarnExplicitObject(category *PyObject, message *PyObject, filename *PyObject, lineno int, module *PyObject, registry *PyObject) error {

	ret := C.PyErr_WarnExplicitObject((*C.PyObject)(category), (*C.PyObject)(message), (*C.PyObject)(filename), C.int(lineno), (*C.PyObject)(module), (*C.PyObject)(registry))

	if ret == -1 {
		return fmt.Errorf("An exception was raised during warning execution")
	}
	return nil
}

//PyErr_WarnExplicit : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WarnExplicit
func PyErr_WarnExplicit(category *PyObject, message string, filename string, lineno int, module string, registry *PyObject) error {
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	cmodule := C.CString(module)
	defer C.free(unsafe.Pointer(cmodule))

	ret := C.PyErr_WarnExplicit((*C.PyObject)(category), cmessage, cfilename, C.int(lineno), cmodule, (*C.PyObject)(registry))

	if ret == -1 {
		return fmt.Errorf("An exception was raised during warning execution")
	}
	return nil
}

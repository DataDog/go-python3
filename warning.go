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
	Warning                   = togo(C.PyExc_Warning)
	BytesWarning              = togo(C.PyExc_BytesWarning)
	DeprecationWarning        = togo(C.PyExc_DeprecationWarning)
	FutureWarning             = togo(C.PyExc_FutureWarning)
	ImportWarning             = togo(C.PyExc_ImportWarning)
	PendingDeprecationWarning = togo(C.PyExc_PendingDeprecationWarning)
	ResourceWarning           = togo(C.PyExc_ResourceWarning)
	RuntimeWarning            = togo(C.PyExc_RuntimeWarning)
	SyntaxWarning             = togo(C.PyExc_SyntaxWarning)
	UnicodeWarning            = togo(C.PyExc_UnicodeWarning)
	UserWarning               = togo(C.PyExc_UserWarning)
)

//PyErr_WarnEx : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WarnEx
func PyErr_WarnEx(category *PyObject, message string, stack_level int) error {
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))

	ret := C.PyErr_WarnEx(toc(category), cmessage, C.Py_ssize_t(stack_level))

	if ret == -1 {
		return fmt.Errorf("An exception was raised during warning execution")
	}
	return nil
}

//PyErr_WarnExplicitObject : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_WarnExplicitObject
func PyErr_WarnExplicitObject(category *PyObject, message *PyObject, filename *PyObject, lineno int, module *PyObject, registry *PyObject) error {

	ret := C.PyErr_WarnExplicitObject(toc(category), toc(message), toc(filename), C.int(lineno), toc(module), toc(registry))

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

	ret := C.PyErr_WarnExplicit(toc(category), cmessage, cfilename, C.int(lineno), cmodule, toc(registry))

	if ret == -1 {
		return fmt.Errorf("An exception was raised during warning execution")
	}
	return nil
}

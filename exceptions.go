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
All standard Python exceptions are available as global variables whose names are PyExc_ followed by the Python exception name.
These have the type PyObject*; they are all class objects.
*/
var (
	PyExc_BaseException          = togo(C.PyExc_BaseException)
	PyExc_Exception              = togo(C.PyExc_Exception)
	PyExc_ArithmeticError        = togo(C.PyExc_ArithmeticError)
	PyExc_AssertionError         = togo(C.PyExc_AssertionError)
	PyExc_AttributeError         = togo(C.PyExc_AttributeError)
	PyExc_BlockingIOError        = togo(C.PyExc_BlockingIOError)
	PyExc_BrokenPipeError        = togo(C.PyExc_BrokenPipeError)
	PyExc_BufferError            = togo(C.PyExc_BufferError)
	PyExc_ChildProcessError      = togo(C.PyExc_ChildProcessError)
	PyExc_ConnectionAbortedError = togo(C.PyExc_ConnectionAbortedError)
	PyExc_ConnectionError        = togo(C.PyExc_ConnectionError)
	PyExc_ConnectionRefusedError = togo(C.PyExc_ConnectionRefusedError)
	PyExc_ConnectcionResetError  = togo(C.PyExc_ConnectionResetError)
	PyExc_EOFError               = togo(C.PyExc_EOFError)
	PyExc_FileExistsError        = togo(C.PyExc_FileExistsError)
	PyExc_FileNotFoundError      = togo(C.PyExc_FileNotFoundError)
	PyExc_FloatingPointError     = togo(C.PyExc_FloatingPointError)
	PyExc_GeneratorExit          = togo(C.PyExc_GeneratorExit)
	PyExc_ImportError            = togo(C.PyExc_ImportError)
	PyExc_IndentationError       = togo(C.PyExc_IndentationError)
	PyExc_IndexError             = togo(C.PyExc_IndexError)
	PyExc_InterruptedError       = togo(C.PyExc_InterruptedError)
	PyExc_IsADirectoryError      = togo(C.PyExc_IsADirectoryError)
	PyExc_KeyError               = togo(C.PyExc_KeyError)
	PyExc_KeyboardInterrupt      = togo(C.PyExc_KeyboardInterrupt)
	PyExc_LookupError            = togo(C.PyExc_LookupError)
	PyExc_MemoryError            = togo(C.PyExc_MemoryError)
	PyExc_ModuleNotFoundError    = togo(C.PyExc_ModuleNotFoundError)
	PyExc_NameError              = togo(C.PyExc_NameError)
	PyExc_NotADirectoryError     = togo(C.PyExc_NotADirectoryError)
	PyExc_NotImplementedError    = togo(C.PyExc_NotImplementedError)
	PyExc_OSError                = togo(C.PyExc_OSError)
	PyExc_OverflowError          = togo(C.PyExc_OverflowError)
	PyExc_PermissionError        = togo(C.PyExc_PermissionError)
	PyExc_ProcessLookupError     = togo(C.PyExc_ProcessLookupError)
	PyExc_RecursionError         = togo(C.PyExc_RecursionError)
	PyExc_ReferenceError         = togo(C.PyExc_ReferenceError)
	PyExc_RuntimeError           = togo(C.PyExc_RuntimeError)
	PyExc_StopAsyncIteration     = togo(C.PyExc_StopAsyncIteration)
	PyExc_StopIteration          = togo(C.PyExc_StopIteration)
	PyExc_SyntaxError            = togo(C.PyExc_SyntaxError)
	PyExc_SystemError            = togo(C.PyExc_SystemError)
	PyExc_SystemExit             = togo(C.PyExc_SystemExit)
	PyExc_TabError               = togo(C.PyExc_TabError)
	PyExc_TimeoutError           = togo(C.PyExc_TimeoutError)
	PyExc_TypeError              = togo(C.PyExc_TypeError)
	PyExc_UnboundLocalError      = togo(C.PyExc_UnboundLocalError)
	PyExc_UnicodeDecodeError     = togo(C.PyExc_UnicodeDecodeError)
	PyExc_UnicodeEncodeError     = togo(C.PyExc_UnicodeEncodeError)
	PyExc_UnicodeError           = togo(C.PyExc_UnicodeError)
	PyExc_UnicodeTranslateError  = togo(C.PyExc_UnicodeTranslateError)
	PyExc_ValueError             = togo(C.PyExc_ValueError)
	PyExc_ZeroDivisionError      = togo(C.PyExc_ZeroDivisionError)
)

//PyErr_NewException : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_NewException
func PyErr_NewException(name string, base, dict *PyObject) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyErr_NewException(cname, toc(base), toc(dict)))
}

//PyErr_NewExceptionWithDoc : https://docs.python.org/3/c-api/exceptions.html#c.PyErr_NewExceptionWithDoc
func PyErr_NewExceptionWithDoc(name, doc string, base, dict *PyObject) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cdoc := C.CString(doc)
	defer C.free(unsafe.Pointer(cdoc))

	return togo(C.PyErr_NewExceptionWithDoc(cname, cdoc, toc(base), toc(dict)))
}

//PyException_GetTraceback : https://docs.python.org/3/c-api/exceptions.html#c.PyException_GetTraceback
func PyException_GetTraceback(ex *PyObject) *PyObject {
	return togo(C.PyException_GetTraceback(toc(ex)))
}

//PyException_SetTraceback : https://docs.python.org/3/c-api/exceptions.html#c.PyException_SetTraceback
func PyException_SetTraceback(ex, tb *PyObject) int {
	return int(C.PyException_SetTraceback(toc(ex), toc(tb)))
}

//PyException_GetContext : https://docs.python.org/3/c-api/exceptions.html#c.PyException_GetContext
func PyException_GetContext(ex *PyObject) *PyObject {
	return togo(C.PyException_GetContext(toc(ex)))
}

//PyException_SetContext : https://docs.python.org/3/c-api/exceptions.html#c.PyException_SetContext
func PyException_SetContext(ex, ctx *PyObject) {
	C.PyException_SetContext(toc(ex), toc(ctx))
}

//PyException_GetCause : https://docs.python.org/3/c-api/exceptions.html#c.PyException_GetCause
func PyException_GetCause(ex *PyObject) *PyObject {
	return togo(C.PyException_GetCause(toc(ex)))
}

//PyException_SetCause : https://docs.python.org/3/c-api/exceptions.html#c.PyException_SetCause
func PyException_SetCause(ex, cause *PyObject) {
	C.PyException_SetCause(toc(ex), toc(cause))
}

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
	BaseException          = togo(C.PyExc_BaseException)
	Exception              = togo(C.PyExc_Exception)
	ArithmeticError        = togo(C.PyExc_ArithmeticError)
	AssertionError         = togo(C.PyExc_AssertionError)
	AttributeError         = togo(C.PyExc_AttributeError)
	BlockingIOError        = togo(C.PyExc_BlockingIOError)
	BrokenPipeError        = togo(C.PyExc_BrokenPipeError)
	BufferError            = togo(C.PyExc_BufferError)
	ChildProcessError      = togo(C.PyExc_ChildProcessError)
	ConnectionAbortedError = togo(C.PyExc_ConnectionAbortedError)
	ConnectionError        = togo(C.PyExc_ConnectionError)
	ConnectionRefusedError = togo(C.PyExc_ConnectionRefusedError)
	ConnectcionResetError  = togo(C.PyExc_ConnectionResetError)
	EOFError               = togo(C.PyExc_EOFError)
	FileExistsError        = togo(C.PyExc_FileExistsError)
	FileNotFoundError      = togo(C.PyExc_FileNotFoundError)
	FloatingPointError     = togo(C.PyExc_FloatingPointError)
	GeneratorExit          = togo(C.PyExc_GeneratorExit)
	ImportError            = togo(C.PyExc_ImportError)
	IndentationError       = togo(C.PyExc_IndentationError)
	IndexError             = togo(C.PyExc_IndexError)
	InterruptedError       = togo(C.PyExc_InterruptedError)
	IsADirectoryError      = togo(C.PyExc_IsADirectoryError)
	KeyError               = togo(C.PyExc_KeyError)
	KeyboardInterrupt      = togo(C.PyExc_KeyboardInterrupt)
	LookupError            = togo(C.PyExc_LookupError)
	MemoryError            = togo(C.PyExc_MemoryError)
	ModuleNotFoundError    = togo(C.PyExc_ModuleNotFoundError)
	NameError              = togo(C.PyExc_NameError)
	NotADirectoryError     = togo(C.PyExc_NotADirectoryError)
	NotImplementedError    = togo(C.PyExc_NotImplementedError)
	OSError                = togo(C.PyExc_OSError)
	OverflowError          = togo(C.PyExc_OverflowError)
	PermissionError        = togo(C.PyExc_PermissionError)
	ProcessLookupError     = togo(C.PyExc_ProcessLookupError)
	RecursionError         = togo(C.PyExc_RecursionError)
	ReferenceError         = togo(C.PyExc_ReferenceError)
	RuntimeError           = togo(C.PyExc_RuntimeError)
	StopAsyncIteration     = togo(C.PyExc_StopAsyncIteration)
	StopIteration          = togo(C.PyExc_StopIteration)
	SyntaxError            = togo(C.PyExc_SyntaxError)
	SystemError            = togo(C.PyExc_SystemError)
	SystemExit             = togo(C.PyExc_SystemExit)
	TabError               = togo(C.PyExc_TabError)
	TimeoutError           = togo(C.PyExc_TimeoutError)
	TypeError              = togo(C.PyExc_TypeError)
	UnboundLocalError      = togo(C.PyExc_UnboundLocalError)
	UnicodeDecodeError     = togo(C.PyExc_UnicodeDecodeError)
	UnicodeEncodeError     = togo(C.PyExc_UnicodeEncodeError)
	UnicodeError           = togo(C.PyExc_UnicodeError)
	UnicodeTranslateError  = togo(C.PyExc_UnicodeTranslateError)
	ValueError             = togo(C.PyExc_ValueError)
	ZeroDivisionError      = togo(C.PyExc_ZeroDivisionError)
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
func PyException_SetTraceback(ex, tb *PyObject) {
	C.PyException_SetTraceback(toc(ex), toc(tb))
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

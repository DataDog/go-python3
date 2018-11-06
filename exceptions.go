package python3

/*
#include "Python.h"
*/
import "C"

import (
	"unsafe"
)

var (
	BaseException          = (*PyObject)(C.PyExc_BaseException)
	Exception              = (*PyObject)(C.PyExc_Exception)
	ArithmeticError        = (*PyObject)(C.PyExc_ArithmeticError)
	AssertionError         = (*PyObject)(C.PyExc_AssertionError)
	AttributeError         = (*PyObject)(C.PyExc_AttributeError)
	BlockingIOError        = (*PyObject)(C.PyExc_BlockingIOError)
	BrokenPipeError        = (*PyObject)(C.PyExc_BrokenPipeError)
	BufferError            = (*PyObject)(C.PyExc_BufferError)
	ChildProcessError      = (*PyObject)(C.PyExc_ChildProcessError)
	ConnectionAbortedError = (*PyObject)(C.PyExc_ConnectionAbortedError)
	ConnectionError        = (*PyObject)(C.PyExc_ConnectionError)
	ConnectionRefusedError = (*PyObject)(C.PyExc_ConnectionRefusedError)
	ConnectionResetError   = (*PyObject)(C.PyExc_ConnectionResetError)
	EOFError               = (*PyObject)(C.PyExc_EOFError)
	FileExistsError        = (*PyObject)(C.PyExc_FileExistsError)
	FileNotFoundError      = (*PyObject)(C.PyExc_FileNotFoundError)
	FloatingPointError     = (*PyObject)(C.PyExc_FloatingPointError)
	GeneratorExit          = (*PyObject)(C.PyExc_GeneratorExit)
	ImportError            = (*PyObject)(C.PyExc_ImportError)
	IndentationError       = (*PyObject)(C.PyExc_IndentationError)
	IndexError             = (*PyObject)(C.PyExc_IndexError)
	InterruptedError       = (*PyObject)(C.PyExc_InterruptedError)
	IsADirectoryError      = (*PyObject)(C.PyExc_IsADirectoryError)
	KeyError               = (*PyObject)(C.PyExc_KeyError)
	KeyboardInterrupt      = (*PyObject)(C.PyExc_KeyboardInterrupt)
	LookupError            = (*PyObject)(C.PyExc_LookupError)
	MemoryError            = (*PyObject)(C.PyExc_MemoryError)
	ModuleNotFoundError    = (*PyObject)(C.PyExc_ModuleNotFoundError)
	NameError              = (*PyObject)(C.PyExc_NameError)
	NotADirectoryError     = (*PyObject)(C.PyExc_NotADirectoryError)
	NotImplementedError    = (*PyObject)(C.PyExc_NotImplementedError)
	OSError                = (*PyObject)(C.PyExc_OSError)
	OverflowError          = (*PyObject)(C.PyExc_OverflowError)
	PermissionError        = (*PyObject)(C.PyExc_PermissionError)
	ProcessLookupError     = (*PyObject)(C.PyExc_ProcessLookupError)
	RecursionError         = (*PyObject)(C.PyExc_RecursionError)
	ReferenceError         = (*PyObject)(C.PyExc_ReferenceError)
	RuntimeError           = (*PyObject)(C.PyExc_RuntimeError)
	StopAsyncIteration     = (*PyObject)(C.PyExc_StopAsyncIteration)
	StopIteration          = (*PyObject)(C.PyExc_StopIteration)
	SyntaxError            = (*PyObject)(C.PyExc_SyntaxError)
	SystemError            = (*PyObject)(C.PyExc_SystemError)
	SystemExit             = (*PyObject)(C.PyExc_SystemExit)
	TabError               = (*PyObject)(C.PyExc_TabError)
	TimeoutError           = (*PyObject)(C.PyExc_TimeoutError)
	TypeError              = (*PyObject)(C.PyExc_TypeError)
	UnboundLocalError      = (*PyObject)(C.PyExc_UnboundLocalError)
	UnicodeDecodeError     = (*PyObject)(C.PyExc_UnicodeDecodeError)
	UnicodeEncodeError     = (*PyObject)(C.PyExc_UnicodeEncodeError)
	UnicodeError           = (*PyObject)(C.PyExc_UnicodeError)
	UnicodeTranslateError  = (*PyObject)(C.PyExc_UnicodeTranslateError)
	ValueError             = (*PyObject)(C.PyExc_ValueError)
	ZeroDivisionError      = (*PyObject)(C.PyExc_ZeroDivisionError)
)

func PyErr_NewException(name string, base, dict *PyObject) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return (*PyObject)(C.PyErr_NewException(cname, (*C.PyObject)(base), (*C.PyObject)(dict)))
}

func PyErr_NewExceptionWithDoc(name, doc string, base, dict *PyObject) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cdoc := C.CString(doc)
	defer C.free(unsafe.Pointer(cdoc))

	return (*PyObject)(C.PyErr_NewExceptionWithDoc(cname, cdoc, (*C.PyObject)(base), (*C.PyObject)(dict)))
}

func PyException_GetTraceback(ex *PyObject) *PyObject {
	return (*PyObject)(C.PyException_GetTraceback((*C.PyObject)(ex)))
}

func PyException_SetTraceback(ex, tb *PyObject) {
	C.PyException_SetTraceback((*C.PyObject)(ex), (*C.PyObject)(tb))
}

func PyException_GetContext(ex *PyObject) *PyObject {
	return (*PyObject)(C.PyException_GetContext((*C.PyObject)(ex)))
}

func PyException_SetContext(ex, ctx *PyObject) {
	C.PyException_SetContext((*C.PyObject)(ex), (*C.PyObject)(ctx))
}

func PyException_GetCause(ex *PyObject) *PyObject {
	return (*PyObject)(C.PyException_GetCause((*C.PyObject)(ex)))
}

func PyException_SetCause(ex, cause *PyObject) {
	C.PyException_SetCause((*C.PyObject)(ex), (*C.PyObject)(cause))
}

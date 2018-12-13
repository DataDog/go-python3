package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorSetString(t *testing.T) {
	Py_Initialize()

	PyErr_SetString(PyExc_BaseException, "test message")

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorSetObject(t *testing.T) {
	Py_Initialize()

	message := PyUnicode_FromString("test message")
	defer message.DecRef()

	PyErr_SetObject(PyExc_BaseException, message)

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Print()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorSetNone(t *testing.T) {
	Py_Initialize()

	message := PyUnicode_FromString("test message")
	defer message.DecRef()

	PyErr_SetNone(PyExc_BaseException)

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Print()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorSetObjectEx(t *testing.T) {
	Py_Initialize()

	message := PyUnicode_FromString("test message")
	defer message.DecRef()

	PyErr_SetObject(PyExc_BaseException, message)

	assert.NotNil(t, PyErr_Occurred())
	PyErr_PrintEx(false)
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorWriteUnraisable(t *testing.T) {
	Py_Initialize()

	message := PyUnicode_FromString("unraisable exception")
	defer message.DecRef()

	PyErr_WriteUnraisable(message)

	assert.Nil(t, PyErr_Occurred())
}

func TestErrorBadArgument(t *testing.T) {
	Py_Initialize()

	PyErr_BadArgument()

	assert.NotNil(t, PyErr_Occurred())

	PyErr_Clear()

	assert.Nil(t, PyErr_Occurred())
}

func TestErrorNoMemory(t *testing.T) {
	Py_Initialize()

	PyErr_NoMemory()

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorBadInternalCall(t *testing.T) {
	Py_Initialize()

	PyErr_BadInternalCall()

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorImportError(t *testing.T) {
	Py_Initialize()

	message := PyUnicode_FromString("test message")
	defer message.DecRef()

	PyErr_SetImportError(message, nil, nil)

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorImportErrorSubclass(t *testing.T) {
	Py_Initialize()

	message := PyUnicode_FromString("test message")
	defer message.DecRef()

	PyErr_SetImportErrorSubclass(message, nil, nil, Dict)

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorSyntax(t *testing.T) {
	Py_Initialize()

	PyErr_SetNone(PyExc_SyntaxError)

	filename := "test.py"
	PyErr_SyntaxLocation(filename, 0)

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorSyntaxEx(t *testing.T) {
	Py_Initialize()

	PyErr_SetNone(PyExc_SyntaxError)

	filename := "test.py"
	PyErr_SyntaxLocationEx(filename, 0, 0)

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorSyntaxLocation(t *testing.T) {
	Py_Initialize()

	PyErr_SetNone(PyExc_SyntaxError)

	filename := PyUnicode_FromString("test.py")
	defer filename.DecRef()

	PyErr_SyntaxLocationObject(filename, 0, 0)

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorExceptionMatches(t *testing.T) {
	Py_Initialize()

	PyErr_SetNone(PyExc_BufferError)

	assert.True(t, PyErr_ExceptionMatches(PyExc_BufferError))

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorGivenExceptionMatches(t *testing.T) {
	Py_Initialize()

	assert.True(t, PyErr_GivenExceptionMatches(PyExc_BufferError, PyExc_BufferError))
}

func TestErrorFetchRestore(t *testing.T) {
	Py_Initialize()

	PyErr_SetNone(PyExc_BufferError)

	exc, value, traceback := PyErr_Fetch()
	assert.Nil(t, PyErr_Occurred())

	assert.True(t, PyErr_GivenExceptionMatches(exc, PyExc_BufferError))
	assert.Nil(t, value)
	assert.Nil(t, traceback)

	PyErr_Restore(exc, value, traceback)

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorNormalizeExceptionRestore(t *testing.T) {
	Py_Initialize()

	PyErr_SetNone(PyExc_BufferError)

	exc, value, traceback := PyErr_Fetch()
	exc, value, traceback = PyErr_NormalizeException(exc, value, traceback)
	assert.Nil(t, PyErr_Occurred())

	assert.True(t, PyErr_GivenExceptionMatches(exc, PyExc_BufferError))
	assert.Equal(t, 1, value.IsInstance(exc))
	assert.Nil(t, traceback)

	PyErr_Restore(exc, value, traceback)

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorGetSetExcInfo(t *testing.T) {
	Py_Initialize()

	PyErr_SetNone(PyExc_BufferError)

	exc, value, traceback := PyErr_GetExcInfo()

	assert.True(t, PyErr_GivenExceptionMatches(exc, Py_None), PyUnicode_AsUTF8(exc.Repr()))
	assert.Nil(t, value)
	assert.Nil(t, traceback)

	PyErr_SetExcInfo(exc, value, traceback)

	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

func TestErrorInterrupt(t *testing.T) {
	Py_Initialize()

	PyErr_SetInterrupt()

	assert.Equal(t, -1, PyErr_CheckSignals())

	exc := PyErr_Occurred()
	assert.True(t, PyErr_GivenExceptionMatches(exc, PyExc_TypeError))

	assert.NotNil(t, PyErr_Occurred())
	PyErr_Clear()
	assert.Nil(t, PyErr_Occurred())
}

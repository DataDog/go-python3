package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflectionBuiltins(t *testing.T) {
	Py_Initialize()

	builtins := PyEval_GetBuiltins()
	assert.NotNil(t, builtins)

	len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(len))
}

func TestReflectionLocals(t *testing.T) {
	Py_Initialize()

	locals := PyEval_GetLocals()
	assert.Nil(t, locals)
}

func TestReflectionGlobals(t *testing.T) {
	Py_Initialize()

	globals := PyEval_GetGlobals()
	assert.Nil(t, globals)
}

func TestReflectionFuncName(t *testing.T) {
	Py_Initialize()

	builtins := PyEval_GetBuiltins()
	assert.NotNil(t, builtins)

	len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(len))

	assert.Equal(t, "len", PyEval_GetFuncName(len))
}
func TestReflectionFuncDesc(t *testing.T) {
	Py_Initialize()

	builtins := PyEval_GetBuiltins()
	assert.NotNil(t, builtins)

	len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(len))

	assert.Equal(t, "()", PyEval_GetFuncDesc(len))
}

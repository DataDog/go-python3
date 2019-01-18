package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExceptionNew(t *testing.T) {
	Py_Initialize()

	exc := PyErr_NewException("test_module.TestException", nil, nil)
	assert.NotNil(t, exc)
	defer exc.DecRef()
}

func TestExceptionNewDoc(t *testing.T) {
	Py_Initialize()

	exc := PyErr_NewExceptionWithDoc("test_module.TestException", "docstring", nil, nil)
	assert.NotNil(t, exc)
	defer exc.DecRef()
}

func TestExceptionContext(t *testing.T) {
	Py_Initialize()

	exc := PyErr_NewException("test_module.TestException", nil, nil)
	assert.NotNil(t, exc)
	defer exc.DecRef()

	PyException_SetContext(exc, PyExc_BrokenPipeError)

	assert.Equal(t, PyExc_BrokenPipeError, PyException_GetContext(exc))
}

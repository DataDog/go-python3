package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWarnEx(t *testing.T) {
	Py_Initialize()

	assert.Zero(t, PyErr_WarnEx(PyExc_RuntimeWarning, "test warning", 3))
}

func TestWarnExplicitObject(t *testing.T) {
	Py_Initialize()

	message := PyUnicode_FromString("test warning")
	defer message.DecRef()

	filename := PyUnicode_FromString("test.py")
	defer filename.DecRef()

	module := PyUnicode_FromString("test_module")
	defer module.DecRef()

	assert.Zero(t, PyErr_WarnExplicitObject(PyExc_RuntimeError, message, filename, 4, module, nil))
}

func TestWarnExplicit(t *testing.T) {
	Py_Initialize()

	assert.Zero(t, PyErr_WarnExplicit(PyExc_RuntimeError, "test warning", "test.py", 4, "test_module", nil))
}

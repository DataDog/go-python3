package python3

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPyFloatCheck(t *testing.T) {
	Py_Initialize()

	pyFloat := PyFloat_FromDouble(345.)
	assert.True(t, PyFloat_Check(pyFloat))
	assert.True(t, PyFloat_CheckExact(pyFloat))
	pyFloat.DecRef()
}

func TestPyFloatFromAsDouble(t *testing.T) {
	Py_Initialize()
	v := 2354.
	pyFloat := PyFloat_FromDouble(v)
	assert.NotNil(t, pyFloat)
	assert.Equal(t, v, PyFloat_AsDouble(pyFloat))
	pyFloat.DecRef()
}

func TestPyFloatFromAsString(t *testing.T) {
	Py_Initialize()
	v := 2354
	s := strconv.Itoa(v)
	pyString := PyUnicode_FromString(s)
	defer pyString.DecRef()

	pyFloat := PyFloat_FromString(pyString)
	assert.NotNil(t, pyFloat)
	assert.Equal(t, float64(v), PyFloat_AsDouble(pyFloat))
	pyFloat.DecRef()
}

func TestPyFloatMinMax(t *testing.T) {
	Py_Initialize()

	assert.Equal(t, math.MaxFloat64, PyFloat_GetMax())

	assert.Equal(t, 2.2250738585072014e-308, PyFloat_GetMin())
}

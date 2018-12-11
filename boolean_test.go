package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolCheck(t *testing.T) {
	Py_Initialize()

	assert.True(t, PyBool_Check(Py_True))
	assert.True(t, PyBool_Check(Py_False))
}

func TestBoolFromLong(t *testing.T) {
	Py_Initialize()

	assert.Equal(t, Py_True, PyBool_FromLong(1))
	assert.Equal(t, Py_False, PyBool_FromLong(0))
}

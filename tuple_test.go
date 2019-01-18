package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTupleCheck(t *testing.T) {
	Py_Initialize()

	tuple := PyTuple_New(0)
	assert.True(t, PyTuple_Check(tuple))
	assert.True(t, PyTuple_CheckExact(tuple))
	defer tuple.DecRef()
}

func TestTupleNew(t *testing.T) {
	Py_Initialize()

	tuple := PyTuple_New(0)
	assert.NotNil(t, tuple)
	defer tuple.DecRef()
}

func TestTupleSize(t *testing.T) {
	Py_Initialize()

	size := 45
	tuple := PyTuple_New(size)
	assert.Equal(t, size, PyTuple_Size(tuple))
	defer tuple.DecRef()
}

func TestTupleGetSetItem(t *testing.T) {
	Py_Initialize()

	s := PyUnicode_FromString("test")

	i := PyLong_FromGoInt(34)

	tuple := PyTuple_New(2)
	defer tuple.DecRef()

	assert.Zero(t, PyTuple_SetItem(tuple, 0, s))
	assert.Zero(t, PyTuple_SetItem(tuple, 1, i))

	assert.Equal(t, i, PyTuple_GetItem(tuple, 1))
}

func TestTupleGetSlice(t *testing.T) {
	Py_Initialize()

	s := PyUnicode_FromString("test")

	i := PyLong_FromGoInt(34)

	tuple := PyTuple_New(2)
	defer tuple.DecRef()

	assert.Zero(t, PyTuple_SetItem(tuple, 0, s))
	assert.Zero(t, PyTuple_SetItem(tuple, 1, i))

	slice := PyTuple_GetSlice(tuple, 0, 1)
	defer slice.DecRef()

	assert.True(t, PyTuple_Check(slice))
	assert.Equal(t, 1, PyTuple_Size(slice))
	assert.Equal(t, s, PyTuple_GetItem(slice, 0))
}

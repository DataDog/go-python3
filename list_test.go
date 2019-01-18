package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	Py_Initialize()

	list := PyList_New(0)
	assert.True(t, PyList_Check(list))
	assert.True(t, PyList_CheckExact(list))
	defer list.DecRef()

	s := PyUnicode_FromString("hello")
	assert.NotNil(t, s)

	i := PyLong_FromGoInt(123)
	assert.NotNil(t, i)

	f := PyFloat_FromDouble(123.4)
	assert.NotNil(t, f)

	assert.Zero(t, PyList_Append(list, i))
	assert.Zero(t, PyList_Insert(list, 0, s))

	assert.Equal(t, 2, PyList_Size(list))

	assert.Zero(t, PyList_SetItem(list, 0, f))

	assert.Equal(t, f, PyList_GetItem(list, 0))

	assert.Zero(t, PyList_Sort(list))
	assert.Equal(t, i, PyList_GetItem(list, 0))

	assert.Zero(t, PyList_Reverse(list))
	assert.Equal(t, f, PyList_GetItem(list, 0))

	s = PyUnicode_FromString("world")
	assert.NotNil(t, s)

	list2 := PyList_New(1)
	defer list2.DecRef()

	assert.Zero(t, PyList_SetItem(list2, 0, s))

	assert.Zero(t, PyList_SetSlice(list, 0, 1, list2))

	list3 := PyList_GetSlice(list, 0, 1)
	assert.NotNil(t, list3)
	defer list3.DecRef()

	assert.Equal(t, 1, list2.RichCompareBool(list3, Py_EQ))

	tuple := PyList_AsTuple(list)
	assert.NotNil(t, tuple)
	defer tuple.DecRef()

	world := PyTuple_GetItem(tuple, 0)
	assert.NotNil(t, world)

	assert.Equal(t, "world", PyUnicode_AsUTF8(world))

	PyList_ClearFreeList()

}

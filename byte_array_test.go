package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteArray(t *testing.T) {
	Py_Initialize()

	s1 := "aaaaaaaa"
	s2 := "bbbbbbbb"

	array1 := PyByteArray_FromStringAndSize(s1)
	assert.True(t, PyByteArray_Check(array1))
	assert.True(t, PyByteArray_CheckExact(array1))
	defer array1.DecRef()

	bytes := PyBytes_FromString(s2)
	assert.NotNil(t, bytes)
	defer bytes.DecRef()

	array2 := PyByteArray_FromObject(bytes)
	assert.NotNil(t, array2)
	defer array2.DecRef()

	newArray := PyByteArray_Concat(array1, array2)
	defer newArray.DecRef()

	length := 20
	PyByteArray_Resize(newArray, 20)

	assert.Equal(t, length, PyByteArray_Size(newArray))

	assert.Equal(t, s1+s2, PyByteArray_AsString(newArray))

}

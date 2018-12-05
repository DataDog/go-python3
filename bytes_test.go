package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	Py_Initialize()

	s1 := "aaaaaaaa"
	s2 := "bbbbbbbb"

	bytes1 := PyBytes_FromString(s1)
	assert.True(t, PyBytes_Check(bytes1))
	assert.True(t, PyBytes_CheckExact(bytes1))
	defer bytes1.DecRef()

	array := PyByteArray_FromStringAndSize(s2)
	assert.NotNil(t, array)
	defer array.DecRef()

	bytes2 := PyBytes_FromObject(array)
	assert.NotNil(t, bytes2)
	defer bytes2.DecRef()

	newBytes := PyBytes_Concat(bytes1, bytes2)
	assert.NotNil(t, newBytes)

	newBytes = PyBytes_ConcatAndDel(bytes1, newBytes)
	assert.NotNil(t, newBytes)
	defer newBytes.DecRef()

	assert.Equal(t, 24, PyBytes_Size(newBytes))
	assert.Equal(t, s1+s1+s2, PyBytes_AsString(newBytes))

}

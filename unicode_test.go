package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnicodeNew(t *testing.T) {
	Py_Initialize()

	s := PyUnicode_New(20, 'z')
	assert.NotNil(t, s)
	defer s.DecRef()
}

func TestUnicodeFromString(t *testing.T) {
	Py_Initialize()

	u := PyUnicode_FromString("aaa")
	assert.True(t, PyUnicode_Check(u))
	assert.True(t, PyUnicode_CheckExact(u))
	defer u.DecRef()

	assert.Equal(t, 3, PyUnicode_GetLength(u))
}

func TestUnicodeFromEncodedObject(t *testing.T) {
	Py_Initialize()

	b := PyBytes_FromString("bbb")
	assert.NotNil(t, b)
	defer b.DecRef()
	ub := PyUnicode_FromEncodedObject(b, "utf-8", "strict")
	assert.NotNil(t, ub)
	defer ub.DecRef()
}

func TestUnicodeChar(t *testing.T) {
	Py_Initialize()

	u := PyUnicode_FromString("aaa")
	assert.True(t, PyUnicode_Check(u))
	assert.True(t, PyUnicode_CheckExact(u))
	defer u.DecRef()

	assert.Equal(t, 0, PyUnicode_WriteChar(u, 1, 'd'))

	assert.Equal(t, 'd', PyUnicode_ReadChar(u, 1))
}

func TestUnicodeFill(t *testing.T) {
	Py_Initialize()

	u := PyUnicode_FromString("aaa")
	assert.True(t, PyUnicode_Check(u))
	assert.True(t, PyUnicode_CheckExact(u))
	defer u.DecRef()

	assert.Equal(t, 3, PyUnicode_Fill(u, 0, 3, 'c'))

	assert.Equal(t, "ccc", PyUnicode_AsUTF8(u))
}

func TestUnicodeCopyCharacters(t *testing.T) {
	Py_Initialize()

	u := PyUnicode_FromString("aaa")
	assert.True(t, PyUnicode_Check(u))
	assert.True(t, PyUnicode_CheckExact(u))
	defer u.DecRef()

	b := PyBytes_FromString("bbb")
	assert.NotNil(t, b)
	defer b.DecRef()
	ub := PyUnicode_FromEncodedObject(b, "utf-8", "strict")
	assert.NotNil(t, ub)
	defer ub.DecRef()

	assert.Equal(t, 3, PyUnicode_CopyCharacters(ub, u, 0, 0, 3))

	assert.Equal(t, "aaa", PyUnicode_AsUTF8(ub))
}

func TestUnicodeSubstring(t *testing.T) {
	Py_Initialize()

	u := PyUnicode_FromString("aaa")
	assert.True(t, PyUnicode_Check(u))
	assert.True(t, PyUnicode_CheckExact(u))
	defer u.DecRef()

	sub := PyUnicode_Substring(u, 0, 2)
	assert.NotNil(t, sub)
	sub.DecRef()

	assert.Equal(t, "aa", PyUnicode_AsUTF8(sub))
}

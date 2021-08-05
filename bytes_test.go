/*
Unless explicitly stated otherwise all files in this repository are licensed
under the MIT License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesCheck(t *testing.T) {
	Py_Initialize()

	s1 := "aaaaaaaa"

	bytes1 := PyBytes_FromString(s1)
	assert.True(t, PyBytes_Check(bytes1))
	assert.True(t, PyBytes_CheckExact(bytes1))
	defer bytes1.DecRef()
}

func TestBytesFromAsString(t *testing.T) {
	Py_Initialize()

	s1 := "aaaaaaaa"

	bytes1 := PyBytes_FromString(s1)
	defer bytes1.DecRef()

	assert.Equal(t, s1, PyBytes_AsString(bytes1))
}

func TestBytesSize(t *testing.T) {
	Py_Initialize()

	s1 := "aaaaaaaa"

	bytes1 := PyBytes_FromString(s1)
	defer bytes1.DecRef()

	assert.Equal(t, 8, PyBytes_Size(bytes1))
}

func TestBytesConcat(t *testing.T) {
	Py_Initialize()

	s1 := "aaaaaaaa"
	s2 := "bbbbbbbb"

	bytes1 := PyBytes_FromString(s1)

	array := PyByteArray_FromStringAndSize(s2)
	defer array.DecRef()

	bytes2 := PyBytes_FromObject(array)
	assert.NotNil(t, bytes2)

	bytes1 = PyBytes_Concat(bytes1, bytes2)
	assert.NotNil(t, bytes1)
	defer bytes1.DecRef()

	assert.Equal(t, s1+s2, PyBytes_AsString(bytes1))
}

func TestBytesConcatAndDel(t *testing.T) {
	Py_Initialize()

	s1 := "aaaaaaaa"
	s2 := "bbbbbbbb"

	bytes1 := PyBytes_FromString(s1)

	bytes2 := PyBytes_FromString(s2)
	assert.NotNil(t, bytes2)

	bytes1 = PyBytes_ConcatAndDel(bytes1, bytes2)
	assert.NotNil(t, bytes1)
	defer bytes1.DecRef()

	assert.Equal(t, s1+s2, PyBytes_AsString(bytes1))
}

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

func TestByteArrayCheck(t *testing.T) {
	Py_Initialize()

	s1 := "aaaaaaaa"

	array1 := PyByteArray_FromStringAndSize(s1)
	assert.True(t, PyByteArray_Check(array1))
	assert.True(t, PyByteArray_CheckExact(array1))
	defer array1.DecRef()
}

func TestByteArrayFromAsString(t *testing.T) {
	Py_Initialize()

	s1 := "aaaaaaaa"

	array1 := PyByteArray_FromStringAndSize(s1)
	defer array1.DecRef()

	assert.Equal(t, s1, PyByteArray_AsString(array1))
}

func TestByteArrayConcat(t *testing.T) {
	Py_Initialize()

	s1 := "aaaaaaaa"
	s2 := "bbbbbbbb"

	array1 := PyByteArray_FromStringAndSize(s1)
	defer array1.DecRef()

	bytes := PyBytes_FromString(s2)
	assert.NotNil(t, bytes)
	defer bytes.DecRef()

	array2 := PyByteArray_FromObject(bytes)
	assert.NotNil(t, array2)
	defer array2.DecRef()

	newArray := PyByteArray_Concat(array1, array2)
	defer newArray.DecRef()

	assert.Equal(t, s1+s2, PyByteArray_AsString(newArray))
}

func TestByteArrayResize(t *testing.T) {
	Py_Initialize()

	s1 := "aaaaaaaa"

	array1 := PyByteArray_FromStringAndSize(s1)
	defer array1.DecRef()

	length := 20
	PyByteArray_Resize(array1, 20)

	assert.Equal(t, length, PyByteArray_Size(array1))
}

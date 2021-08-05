/*
Unless explicitly stated otherwise all files in this repository are licensed
under the MIT License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPyLongCheck(t *testing.T) {
	Py_Initialize()

	pyLong := PyLong_FromGoInt(345)
	assert.True(t, PyLong_Check(pyLong))
	assert.True(t, PyLong_CheckExact(pyLong))
	pyLong.DecRef()
}

func TestPyLongFromAsLong(t *testing.T) {
	Py_Initialize()
	v := 2354
	pyLong := PyLong_FromLong(v)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsLong(pyLong))
	pyLong.DecRef()
}

func TestPyLongFromAsUnsignedLong(t *testing.T) {
	Py_Initialize()
	v := uint(2354)
	pyLong := PyLong_FromUnsignedLong(v)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsUnsignedLong(pyLong))
	pyLong.DecRef()
}

func TestPyLongFromAsLongLong(t *testing.T) {
	Py_Initialize()
	v := int64(2354)
	pyLong := PyLong_FromLongLong(v)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsLongLong(pyLong))
	pyLong.DecRef()
}

func TestPyLongFromAsUnsignedLongLong(t *testing.T) {
	Py_Initialize()
	v := uint64(2354)
	pyLong := PyLong_FromUnsignedLongLong(v)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsUnsignedLongLong(pyLong))
	pyLong.DecRef()
}

func TestPyLongFromAsDouble(t *testing.T) {
	Py_Initialize()
	v := float64(2354.0)
	pyLong := PyLong_FromDouble(v)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsDouble(pyLong))
	pyLong.DecRef()
}

func TestPyLongFromAsGoFloat64(t *testing.T) {
	Py_Initialize()
	v := float64(2354.0)
	pyLong := PyLong_FromGoFloat64(v)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsDouble(pyLong))
	pyLong.DecRef()
}

func TestPyLongFromAsString(t *testing.T) {
	Py_Initialize()
	v := 2354
	s := strconv.Itoa(v)
	pyLong := PyLong_FromString(s, 10)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsLong(pyLong))
	pyLong.DecRef()
}

func TestPyLongFromAsUnicodeObject(t *testing.T) {
	Py_Initialize()
	v := 2354
	s := strconv.Itoa(v)
	pyUnicode := PyUnicode_FromString(s)
	assert.NotNil(t, pyUnicode)
	pyLong := PyLong_FromUnicodeObject(pyUnicode, 10)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsLong(pyLong))
	pyLong.DecRef()
	pyUnicode.DecRef()
}

func TestPyLongFromAsGoInt(t *testing.T) {
	Py_Initialize()
	v := 2354
	pyLong := PyLong_FromGoInt(v)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsLong(pyLong))
	pyLong.DecRef()
}

func TestPyLongFromAsGoUint(t *testing.T) {
	Py_Initialize()
	v := uint(2354)
	pyLong := PyLong_FromGoUint(v)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsUnsignedLong(pyLong))
	pyLong.DecRef()
}

func TestPyLongFromAsGoInt64(t *testing.T) {
	Py_Initialize()
	v := int64(2354)
	pyLong := PyLong_FromGoInt64(v)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsLongLong(pyLong))
	pyLong.DecRef()
}

func TestPyLongFromAsGoUint64(t *testing.T) {
	Py_Initialize()
	v := uint64(2354)
	pyLong := PyLong_FromGoUint64(v)
	assert.NotNil(t, pyLong)
	assert.Equal(t, v, PyLong_AsUnsignedLongLong(pyLong))
	pyLong.DecRef()
}

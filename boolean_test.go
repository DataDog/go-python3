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

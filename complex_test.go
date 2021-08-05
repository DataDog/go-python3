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

func TestComplex(t *testing.T) {
	Py_Initialize()

	real := 2.
	imaginary := 5.

	complex := PyComplex_FromDoubles(real, imaginary)
	assert.True(t, PyComplex_Check(complex))
	assert.True(t, PyComplex_CheckExact(complex))
	defer complex.DecRef()

	assert.Equal(t, real, PyComplex_RealAsDouble(complex))
	assert.Equal(t, imaginary, PyComplex_ImagAsDouble(complex))
}

package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplex(t *testing.T) {
	real := 2.
	imaginary := 5.

	complex := PyComplex_FromDoubles(real, imaginary)
	assert.True(t, PyComplex_Check(complex))
	assert.True(t, PyComplex_CheckExact(complex))
	defer complex.DecRef()

	assert.Equal(t, real, PyComplex_RealAsDouble(complex))
	assert.Equal(t, imaginary, PyComplex_ImagAsDouble(complex))
}

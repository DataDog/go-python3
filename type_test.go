package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeCheck(t *testing.T) {
	Py_Initialize()

	assert.True(t, PyType_Check(Type))
	assert.True(t, PyType_CheckExact(Type))
}

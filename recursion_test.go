package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursion(t *testing.T) {
	Py_Initialize()

	assert.Zero(t, Py_EnterRecursiveCall("in test function"))

	Py_LeaveRecursiveCall()

}

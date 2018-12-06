package python3

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunFile(t *testing.T) {
	Py_Initialize()

	pyErr, err := PyRun_AnyFile("tests/test.py")
	assert.Zero(t, pyErr)
	assert.Nil(t, err)

	stdout := PySys_GetObject("stdout")

	result := stdout.CallMethodArgs("getvalue")
	defer result.DecRef()

	assert.Equal(t, "hello world\n", PyUnicode_AsUTF8(result))
}

func TestRunString(t *testing.T) {
	Py_Initialize()

	pythonCode, err := ioutil.ReadFile("tests/test.py")
	assert.Nil(t, err)

	assert.Zero(t, PyRun_SimpleString(string(pythonCode)))

	stdout := PySys_GetObject("stdout")

	result := stdout.CallMethodArgs("getvalue")
	defer result.DecRef()

	assert.Equal(t, "hello world\n", PyUnicode_AsUTF8(result))
}

func TestPyMain(t *testing.T) {
	Py_Initialize()

	pyErr, err := Py_Main([]string{"tests/test.py"})
	assert.Zero(t, pyErr)
	assert.Nil(t, err)
}

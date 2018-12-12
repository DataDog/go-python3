package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSysGetSetObject(t *testing.T) {
	Py_Initialize()

	platform := PySys_GetObject("platform")
	assert.NotNil(t, platform)
	assert.True(t, PyUnicode_Check(platform))
	platform.IncRef()

	newPlatform := PyUnicode_FromString("test")
	defer newPlatform.DecRef()

	assert.Zero(t, PySys_SetObject("platform", newPlatform))

	assert.Equal(t, newPlatform, PySys_GetObject("platform"))

	assert.Zero(t, PySys_SetObject("platform", platform))
}

func TestSysWarnOption(t *testing.T) {
	Py_Finalize()

	assert.Nil(t, PySys_AddWarnOption("ignore"))

	Py_Initialize()

	warnoptions := PySys_GetObject("warnoptions")
	assert.Equal(t, "ignore", PyUnicode_AsUTF8(PyList_GetItem(warnoptions, 0)))

	Py_Finalize()

	PySys_ResetWarnOptions()

	Py_Initialize()

	warnoptions = PySys_GetObject("warnoptions")
	assert.Zero(t, PyList_Size(warnoptions))
}

func TestSysXOption(t *testing.T) {
	Py_Finalize()

	assert.Nil(t, PySys_AddXOption("faulthandler"))

	Py_Initialize()

	XOptions := PySys_GetXOptions()
	faulthandler := PyDict_GetItemString(XOptions, "faulthandler")

	assert.Equal(t, Py_True, faulthandler)
}

func TestSysPath(t *testing.T) {
	Py_Initialize()

	path := PySys_GetObject("path")
	path.IncRef()

	assert.Nil(t, PySys_SetPath("test"))

	newPath := PySys_GetObject("path")
	assert.Equal(t, "test", PyUnicode_AsUTF8(PyList_GetItem(newPath, 0)))

	assert.Zero(t, PySys_SetObject("path", path))
}

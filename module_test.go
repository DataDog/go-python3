package python3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModuleCheck(t *testing.T) {
	Py_Initialize()

	name := "test_module"

	module := PyModule_New(name)
	assert.True(t, PyModule_Check(module))
	assert.True(t, PyModule_CheckExact(module))
	defer module.DecRef()
}

func TestModuleNew(t *testing.T) {
	Py_Initialize()

	name := "test_module"

	module := PyModule_New(name)
	assert.NotNil(t, module)
	defer module.DecRef()
}

func TestModuleNewObject(t *testing.T) {
	Py_Initialize()

	name := "test_module"

	pyName := PyUnicode_FromString(name)
	assert.NotNil(t, pyName)
	defer pyName.DecRef()

	module := PyModule_NewObject(pyName)
	assert.NotNil(t, module)
	defer module.DecRef()
}

func TestModuleGetDict(t *testing.T) {
	Py_Initialize()

	name := "sys"
	pyName := PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := PyImport_ImportModule(name)
	defer sys.DecRef()

	dict := PyModule_GetDict(sys)
	assert.True(t, PyDict_Check(dict))
}

func TestModuleGetName(t *testing.T) {
	Py_Initialize()

	name := "sys"
	pyName := PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := PyImport_ImportModule(name)
	defer sys.DecRef()

	assert.Equal(t, name, PyModule_GetName(sys))
}

func TestModuleGetNameObject(t *testing.T) {
	Py_Initialize()

	name := "sys"
	pyName := PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := PyImport_ImportModule(name)
	defer sys.DecRef()

	assert.Equal(t, 1, pyName.RichCompareBool(PyModule_GetNameObject(sys), Py_EQ))
}

func TestModuleGetState(t *testing.T) {
	Py_Initialize()

	name := "sys"
	pyName := PyUnicode_FromString(name)
	defer pyName.DecRef()

	sys := PyImport_ImportModule(name)
	defer sys.DecRef()

	state := PyModule_GetState(sys)
	assert.True(t, state == nil)
}

func TestModuleGetFilenameObject(t *testing.T) {
	Py_Initialize()

	name := "test"
	pyName := PyUnicode_FromString(name)
	defer pyName.DecRef()

	test := PyImport_ImportModule(name)
	defer test.DecRef()

	pyFilename := PyModule_GetFilenameObject(test)
	assert.NotNil(t, pyFilename)
	filename := PyUnicode_AsUTF8(pyFilename)

	assert.True(t, strings.Contains(filename, "/test/__init__.py"))
}

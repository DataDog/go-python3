package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportModule(t *testing.T) {
	Py_Initialize()

	mathName := PyUnicode_FromString("math")
	defer mathName.DecRef()
	platformName := PyUnicode_FromString("platform")
	defer platformName.DecRef()

	os := PyImport_ImportModule("os")
	test := PyImport_ImportModuleEx("test", nil, nil, nil)
	math := PyImport_ImportModuleLevelObject(mathName, nil, nil, nil, 0)
	sys := PyImport_ImportModuleLevel("sys", nil, nil, nil, 0)
	platform := PyImport_Import(platformName)

	assert.NotNil(t, os)
	assert.NotNil(t, test)
	assert.NotNil(t, math)
	assert.NotNil(t, sys)
	assert.NotNil(t, platform)

	os.DecRef()
	test.DecRef()
	math.DecRef()
}

func TestReloadModule(t *testing.T) {
	Py_Initialize()

	os := PyImport_ImportModule("os")
	assert.NotNil(t, os)
	defer os.DecRef()

	newOs := PyImport_ReloadModule(os)
	assert.NotNil(t, newOs)
	defer newOs.DecRef()

	// PyImport_ReloadModule return a new reference, pointer should be the same
	assert.Equal(t, os, newOs)
}

func TestAddModuleObject(t *testing.T) {
	Py_Initialize()

	os := PyImport_ImportModule("os")
	assert.NotNil(t, os)
	defer os.DecRef()

	pyName := PyUnicode_FromString("os.new")
	defer pyName.DecRef()

	new := PyImport_AddModuleObject(pyName)
	assert.NotNil(t, new)
}

func TestAddModule(t *testing.T) {
	Py_Initialize()

	os := PyImport_ImportModule("os")
	assert.NotNil(t, os)
	defer os.DecRef()

	new := PyImport_AddModule("os.new")
	assert.NotNil(t, new)
}

func TestExecCodeModule(t *testing.T) {
	Py_Initialize()

	// fake module
	source := PyUnicode_FromString("__version__ = '2.0'")
	defer source.DecRef()
	filename := PyUnicode_FromString("test_module.py")
	defer filename.DecRef()
	mode := PyUnicode_FromString("exec")
	defer mode.DecRef()

	// perform module load
	builtins := PyEval_GetBuiltins()
	assert.True(t, PyDict_Check(builtins))

	compile := PyDict_GetItemString(builtins, "compile")
	assert.True(t, PyCallable_Check(compile))

	code, _ := compile.CallFunctionObjArgs(source, filename, mode)
	assert.NotNil(t, code)
	defer code.DecRef()

	module := PyImport_ExecCodeModule("test_module", code)
	assert.NotNil(t, module)

}

func TestGetMagicNumber(t *testing.T) {
	Py_Initialize()

	magicNumber := PyImport_GetMagicNumber()
	assert.NotNil(t, magicNumber)
}

func TestGetMagicTag(t *testing.T) {
	Py_Initialize()

	magicTag := PyImport_GetMagicTag()
	assert.NotNil(t, magicTag)
}

func TestGetModuleDict(t *testing.T) {
	Py_Initialize()

	moduleDict := PyImport_GetModuleDict()
	defer moduleDict.DecRef()

	assert.True(t, PyDict_Check(moduleDict))

}

func TestGetModule(t *testing.T) {
	Py_Initialize()

	os := PyImport_ImportModule("os")
	assert.NotNil(t, os)
	defer os.DecRef()

	name := PyUnicode_FromString("os")
	defer name.DecRef()

	new := PyImport_GetModule(name)
	assert.Equal(t, new, os)
}

func TestGetImporter(t *testing.T) {
	Py_Initialize()

	paths := PySys_GetObject("path")
	path := PyList_GetItem(paths, 0)

	assert.NotNil(t, path)
	importer := PyImport_GetImporter(path)
	defer importer.DecRef()

	assert.NotNil(t, importer)

}

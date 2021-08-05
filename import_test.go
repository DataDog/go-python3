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

func TestImportModule(t *testing.T) {
	Py_Initialize()

	os := PyImport_ImportModule("os")
	assert.NotNil(t, os)
	os.DecRef()
}

func TestImportModuleEx(t *testing.T) {
	Py_Initialize()

	test := PyImport_ImportModuleEx("test", nil, nil, nil)
	assert.NotNil(t, test)
	test.DecRef()
}

func TestImportModuleLevelObject(t *testing.T) {
	Py_Initialize()

	mathName := PyUnicode_FromString("math")
	defer mathName.DecRef()

	math := PyImport_ImportModuleLevelObject(mathName, nil, nil, nil, 0)
	assert.NotNil(t, math)
	math.DecRef()
}

func TestImportModuleLevel(t *testing.T) {
	Py_Initialize()

	sys := PyImport_ImportModuleLevel("sys", nil, nil, nil, 0)
	assert.NotNil(t, sys)
	sys.DecRef()
}

func TestImportImport(t *testing.T) {
	Py_Initialize()

	platformName := PyUnicode_FromString("platform")
	defer platformName.DecRef()

	platform := PyImport_Import(platformName)
	assert.NotNil(t, platform)
	platform.DecRef()
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

	code := compile.CallFunctionObjArgs(source, filename, mode)
	assert.NotNil(t, code)
	defer code.DecRef()

	module := PyImport_ExecCodeModule("test_module", code)
	assert.NotNil(t, module)

}

func TestExecCodeModuleEx(t *testing.T) {
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

	code := compile.CallFunctionObjArgs(source, filename, mode)
	assert.NotNil(t, code)
	defer code.DecRef()

	module := PyImport_ExecCodeModuleEx("test_module", code, "test_module.py")
	assert.NotNil(t, module)

}

func TestExecCodeModuleWithPathnames(t *testing.T) {
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

	code := compile.CallFunctionObjArgs(source, filename, mode)
	assert.NotNil(t, code)
	defer code.DecRef()

	module := PyImport_ExecCodeModuleWithPathnames("test_module", code, "test_module.py", "test_module.py")
	assert.NotNil(t, module)

}

func TestExecCodeModuleObject(t *testing.T) {
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

	code := compile.CallFunctionObjArgs(source, filename, mode)
	assert.NotNil(t, code)
	defer code.DecRef()

	moduleName := PyUnicode_FromString("test_module")
	defer moduleName.DecRef()

	module := PyImport_ExecCodeModuleObject(moduleName, code, filename, filename)
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

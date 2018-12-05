package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialization(t *testing.T) {

	Py_Initialize()
	assert.True(t, Py_IsInitialized())
	Py_Finalize()
	assert.False(t, Py_IsInitialized())

}

func TestInitializationEx(t *testing.T) {

	Py_Initialize()
	assert.True(t, Py_IsInitialized())
	assert.Nil(t, Py_FinalizeEx())
	assert.False(t, Py_IsInitialized())

}

func TestProgramName(t *testing.T) {
	Py_Finalize()
	defaultName := Py_GetProgramName()
	name := "py3é"
	Py_SetProgramName(name)
	assert.Equal(t, name, Py_GetProgramName())

	Py_SetProgramName(defaultName)

}

func TestPrefix(t *testing.T) {
	assert.IsType(t, "", Py_GetPrefix())

}

func TestExecPrefix(t *testing.T) {
	assert.IsType(t, "", Py_GetExecPrefix())

}

func TestProgramFullPath(t *testing.T) {
	assert.IsType(t, "", Py_GetProgramFullPath())

}

func TestPath(t *testing.T) {
	Py_Finalize()
	defaultPath := Py_GetPath()
	name := "påth"
	Py_SetPath(name)
	assert.Equal(t, name, Py_GetPath())
	Py_SetPath(defaultPath)

}

func TestVersion(t *testing.T) {
	version := Py_GetVersion()
	assert.IsType(t, "", version)
}

func TestPlatform(t *testing.T) {
	platform := Py_GetPlatform()
	assert.IsType(t, "", platform)
}

func TestCopyright(t *testing.T) {
	copyright := Py_GetCopyright()
	assert.IsType(t, "", copyright)
}

func TestCompiler(t *testing.T) {
	compiler := Py_GetCompiler()
	assert.IsType(t, "", compiler)
}

func TestBuildInfo(t *testing.T) {
	buildInfo := Py_GetBuildInfo()
	assert.IsType(t, "", buildInfo)
}

func TestPythonHome(t *testing.T) {
	name := "høme"
	defaultHome := Py_GetPythonHome()
	Py_SetPythonHome(name)
	assert.Equal(t, name, Py_GetPythonHome())

	Py_SetPythonHome(defaultHome)
}

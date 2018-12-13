/*
Unless explicitly stated otherwise all files in this repository are licensed
under the $license_for_repo License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

/*
#include "Python.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// The argument for Py_SetProgramName, Py_SetPath and Py_SetPythonHome should point to a zero-terminated wide character string in static storage
// whose contents will not change for the duration of the program’s execution
var (
	programName *C.wchar_t
	pythonPath  *C.wchar_t
	pythonHome  *C.wchar_t
)

//Py_Initialize : https://docs.python.org/3/c-api/init.html#c.Py_Initialize
func Py_Initialize() {
	C.Py_Initialize()
}

//Py_InitializeEx : https://docs.python.org/3/c-api/init.html#c.Py_InitializeEx
func Py_InitializeEx(initsigs bool) {
	if initsigs {
		C.Py_InitializeEx(1)
	} else {
		C.Py_InitializeEx(0)
	}
}

//Py_IsInitialized : https://docs.python.org/3/c-api/init.html#c.Py_IsInitialized
func Py_IsInitialized() bool {
	return C.Py_IsInitialized() != 0
}

//Py_FinalizeEx : https://docs.python.org/3/c-api/init.html#c.Py_FinalizeEx
func Py_FinalizeEx() int {
	return int(C.Py_FinalizeEx())
}

//Py_Finalize : https://docs.python.org/3/c-api/init.html#c.Py_Finalize
func Py_Finalize() {
	C.Py_Finalize()
}

//Py_SetStandardStreamEncoding : https://docs.python.org/3/c-api/init.html#c.Py_SetStandardStreamEncoding
func Py_SetStandardStreamEncoding(encoding, errors string) int {
	cencoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(cencoding))

	cerrors := C.CString(errors)
	defer C.free(unsafe.Pointer(cerrors))

	return int(C.Py_SetStandardStreamEncoding(cencoding, cerrors))

}

//Py_SetProgramName : https://docs.python.org/3/c-api/init.html#c.Py_SetProgramName
func Py_SetProgramName(name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	newProgramName := C.Py_DecodeLocale(cname, nil)
	if newProgramName == nil {
		return fmt.Errorf("fail to call Py_DecodeLocale on '%s'", name)
	}
	C.Py_SetProgramName(newProgramName)

	//no operation is performed if nil
	C.PyMem_RawFree(unsafe.Pointer(programName))
	programName = newProgramName

	return nil
}

//Py_GetProgramName : https://docs.python.org/3/c-api/init.html#c.Py_GetProgramName
func Py_GetProgramName() (string, error) {
	wcname := C.Py_GetProgramName()
	if wcname == nil {
		return "", nil
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

//Py_GetPrefix : https://docs.python.org/3/c-api/init.html#c.Py_GetPrefix
func Py_GetPrefix() (string, error) {
	wcname := C.Py_GetPrefix()
	if wcname == nil {
		return "", nil
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

//Py_GetExecPrefix : https://docs.python.org/3/c-api/init.html#c.Py_GetExecPrefix
func Py_GetExecPrefix() (string, error) {
	wcname := C.Py_GetExecPrefix()
	if wcname == nil {
		return "", nil
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

//Py_GetProgramFullPath : https://docs.python.org/3/c-api/init.html#c.Py_GetProgramFullPath
func Py_GetProgramFullPath() (string, error) {
	wcname := C.Py_GetProgramFullPath()
	if wcname == nil {
		return "", nil
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

//Py_GetPath : https://docs.python.org/3/c-api/init.html#c.Py_GetPath
func Py_GetPath() (string, error) {
	wcname := C.Py_GetPath()
	if wcname == nil {
		return "", nil
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	if cname == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname), nil
}

//Py_SetPath : https://docs.python.org/3/c-api/init.html#c.Py_SetPath
func Py_SetPath(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	newPath := C.Py_DecodeLocale(cpath, nil)
	if newPath == nil {
		return fmt.Errorf("fail to call Py_DecodeLocale on '%s'", path)
	}
	C.Py_SetPath(newPath)

	C.PyMem_RawFree(unsafe.Pointer(pythonPath))
	pythonHome = newPath

	return nil
}

//Py_GetVersion : https://docs.python.org/3/c-api/init.html#c.Py_GetVersion
func Py_GetVersion() string {
	cversion := C.Py_GetVersion()
	return C.GoString(cversion)
}

//Py_GetPlatform : https://docs.python.org/3/c-api/init.html#c.Py_GetPlatform
func Py_GetPlatform() string {
	cplatform := C.Py_GetPlatform()
	return C.GoString(cplatform)
}

//Py_GetCopyright : https://docs.python.org/3/c-api/init.html#c.Py_GetCopyright
func Py_GetCopyright() string {
	ccopyright := C.Py_GetCopyright()
	return C.GoString(ccopyright)
}

//Py_GetCompiler : https://docs.python.org/3/c-api/init.html#c.Py_GetCompiler
func Py_GetCompiler() string {
	ccompiler := C.Py_GetCompiler()
	return C.GoString(ccompiler)
}

//Py_GetBuildInfo : https://docs.python.org/3/c-api/init.html#c.Py_GetBuildInfo
func Py_GetBuildInfo() string {
	cbuildInfo := C.Py_GetBuildInfo()
	return C.GoString(cbuildInfo)
}

//PySys_SetArgvEx : https://docs.python.org/3/c-api/init.html#c.PySys_SetArgvEx
func PySys_SetArgvEx(args []string, updatepath bool) error {
	argc := C.int(len(args))
	argv := make([]*C.wchar_t, argc, argc)
	for i, arg := range args {
		carg := C.CString(arg)
		defer C.free(unsafe.Pointer(carg))

		warg := C.Py_DecodeLocale(carg, nil)
		if warg == nil {
			return fmt.Errorf("fail to call Py_DecodeLocale on '%s'", arg)
		}
		// Py_DecodeLocale requires a call to PyMem_RawFree to free the memory
		defer C.PyMem_RawFree(unsafe.Pointer(warg))

		argv[i] = warg
	}

	if updatepath {
		C.PySys_SetArgvEx(argc, (**C.wchar_t)(unsafe.Pointer(&argv[0])), 1)
	} else {
		C.PySys_SetArgvEx(argc, (**C.wchar_t)(unsafe.Pointer(&argv[0])), 0)
	}

	return nil
}

//PySys_SetArgv : https://docs.python.org/3/c-api/init.html#c.PySys_SetArgv
func PySys_SetArgv(args []string) error {
	argc := C.int(len(args))
	argv := make([]*C.wchar_t, argc, argc)
	for i, arg := range args {
		carg := C.CString(arg)
		defer C.free(unsafe.Pointer(carg))

		warg := C.Py_DecodeLocale(carg, nil)
		if warg == nil {
			return fmt.Errorf("fail to call Py_DecodeLocale on '%s'", arg)
		}
		// Py_DecodeLocale requires a call to PyMem_RawFree to free the memory
		defer C.PyMem_RawFree(unsafe.Pointer(warg))

		argv[i] = warg
	}
	C.PySys_SetArgv(argc, (**C.wchar_t)(unsafe.Pointer(&argv[0])))

	return nil
}

//Py_SetPythonHome : https://docs.python.org/3/c-api/init.html#c.Py_SetPythonHome
func Py_SetPythonHome(home string) error {
	chome := C.CString(home)
	defer C.free(unsafe.Pointer(chome))

	newHome := C.Py_DecodeLocale(chome, nil)
	if newHome == nil {
		return fmt.Errorf("fail to call Py_DecodeLocale on '%s'", home)
	}
	C.Py_SetPythonHome(newHome)

	C.PyMem_RawFree(unsafe.Pointer(pythonHome))
	pythonHome = newHome

	return nil
}

//Py_GetPythonHome : https://docs.python.org/3/c-api/init.html#c.Py_GetPythonHome
func Py_GetPythonHome() (string, error) {
	wchome := C.Py_GetPythonHome()
	if wchome == nil {
		return "", nil
	}
	chome := C.Py_EncodeLocale(wchome, nil)
	if chome == nil {
		return "", fmt.Errorf("fail to call Py_EncodeLocale")
	}
	defer C.PyMem_Free(unsafe.Pointer(chome))

	return C.GoString(chome), nil
}

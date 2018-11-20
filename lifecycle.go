package python3

/*
#include "Python.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

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
func Py_FinalizeEx() error {
	if C.Py_FinalizeEx() == -1 {
		return fmt.Errorf("An error occured during finalization")
	}
	return nil
}

//Py_Finalize : https://docs.python.org/3/c-api/init.html#c.Py_Finalize
func Py_Finalize() {
	C.Py_Finalize()
}

//Py_SetStandardStreamEncoding : https://docs.python.org/3/c-api/init.html#c.Py_SetStandardStreamEncoding
func Py_SetStandardStreamEncoding(encoding, errors string) error {
	cencoding := C.CString(encoding)
	defer C.free(unsafe.Pointer(cencoding))

	cerrors := C.CString(errors)
	defer C.free(unsafe.Pointer(cerrors))

	if C.Py_SetStandardStreamEncoding(cencoding, cerrors) != 0 {
		return fmt.Errorf("Py_SetStandardStreamEncoding: an error occured")
	}

	return nil
}

//Py_SetProgramName : https://docs.python.org/3/c-api/init.html#c.Py_SetProgramName
func Py_SetProgramName(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	newProgramName := C.Py_DecodeLocale(cname, nil)
	C.Py_SetProgramName(newProgramName)

	C.PyMem_RawFree(unsafe.Pointer(programName))
	programName = newProgramName

}

//Py_GetProgramName : https://docs.python.org/3/c-api/init.html#c.Py_GetProgramName
func Py_GetProgramName() string {
	wcname := C.Py_GetProgramName()
	if wcname == nil {
		return ""
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname)
}

//Py_GetPrefix : https://docs.python.org/3/c-api/init.html#c.Py_GetPrefix
func Py_GetPrefix() string {
	wcname := C.Py_GetPrefix()
	if wcname == nil {
		return ""
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname)
}

//Py_GetExecPrefix : https://docs.python.org/3/c-api/init.html#c.Py_GetExecPrefix
func Py_GetExecPrefix() string {
	wcname := C.Py_GetExecPrefix()
	if wcname == nil {
		return ""
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname)
}

//Py_GetProgramFullPath : https://docs.python.org/3/c-api/init.html#c.Py_GetProgramFullPath
func Py_GetProgramFullPath() string {
	wcname := C.Py_GetProgramFullPath()
	if wcname == nil {
		return ""
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname)
}

//Py_GetPath : https://docs.python.org/3/c-api/init.html#c.Py_GetPath
func Py_GetPath() string {
	wcname := C.Py_GetPath()
	if wcname == nil {
		return ""
	}
	cname := C.Py_EncodeLocale(wcname, nil)
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname)
}

//Py_SetPath : https://docs.python.org/3/c-api/init.html#c.Py_SetPath
func Py_SetPath(path string) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	newPath := C.Py_DecodeLocale(cpath, nil)
	C.Py_SetPath(newPath)

	C.PyMem_RawFree(unsafe.Pointer(pythonPath))
	pythonHome = newPath
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
func PySys_SetArgvEx(args []string, updatepath bool) {
	argc := C.int(len(args))
	argv := make([]*C.wchar_t, argc, argc)
	for i, arg := range args {
		carg := C.CString(arg)
		defer C.free(unsafe.Pointer(carg))

		warg := C.Py_DecodeLocale(carg, nil)
		// Py_DecodeLocale requires a call to PyMem_RawFree to free the memory
		defer C.PyMem_RawFree(unsafe.Pointer(warg))

		argv[i] = warg
	}

	if updatepath {
		C.PySys_SetArgvEx(argc, (**C.wchar_t)(unsafe.Pointer(&argv[0])), 1)
	} else {
		C.PySys_SetArgvEx(argc, (**C.wchar_t)(unsafe.Pointer(&argv[0])), 0)
	}
}

//PySys_SetArgv : https://docs.python.org/3/c-api/init.html#c.PySys_SetArgv
func PySys_SetArgv(args []string) {
	argc := C.int(len(args))
	argv := make([]*C.wchar_t, argc, argc)
	for i, arg := range args {
		carg := C.CString(arg)
		defer C.free(unsafe.Pointer(carg))

		warg := C.Py_DecodeLocale(carg, nil)
		// Py_DecodeLocale requires a call to PyMem_RawFree to free the memory
		defer C.PyMem_RawFree(unsafe.Pointer(warg))

		argv[i] = warg
	}
	C.PySys_SetArgv(argc, (**C.wchar_t)(unsafe.Pointer(&argv[0])))
}

//Py_SetPythonHome : https://docs.python.org/3/c-api/init.html#c.Py_SetPythonHome
func Py_SetPythonHome(home string) {
	chome := C.CString(home)
	defer C.free(unsafe.Pointer(chome))

	newHome := C.Py_DecodeLocale(chome, nil)
	C.Py_SetPythonHome(newHome)

	C.PyMem_RawFree(unsafe.Pointer(pythonHome))
	pythonHome = newHome
}

//Py_GetPythonHome : https://docs.python.org/3/c-api/init.html#c.Py_GetPythonHome
func Py_GetPythonHome() string {
	wchome := C.Py_GetPythonHome()
	if wchome == nil {
		return ""
	}
	chome := C.Py_EncodeLocale(wchome, nil)
	defer C.PyMem_Free(unsafe.Pointer(chome))

	return C.GoString(chome)
}

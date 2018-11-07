package python3

/*
#include "Python.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
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

	wcname := C.Py_DecodeLocale(cname, nil)
	// wcname is copied, safe to free after calling C.Py_SetProgramName
	defer C.PyMem_RawFree(unsafe.Pointer(wcname))

	C.Py_SetProgramName(wcname)
}

//Py_GetProgramName : https://docs.python.org/3/c-api/init.html#c.Py_GetProgramName
func Py_GetProgramName() string {
	wcname := C.Py_GetProgramName()
	cname := C.Py_EncodeLocale(wcname, nil)
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname)
}

//Py_GetPrefix : https://docs.python.org/3/c-api/init.html#c.Py_GetPrefix
func Py_GetPrefix() string {
	wcname := C.Py_GetPrefix()
	cname := C.Py_EncodeLocale(wcname, nil)
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname)
}

//Py_GetExecPrefix : https://docs.python.org/3/c-api/init.html#c.Py_GetExecPrefix
func Py_GetExecPrefix() string {
	wcname := C.Py_GetExecPrefix()
	cname := C.Py_EncodeLocale(wcname, nil)
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname)
}

//Py_GetProgramFullPath : https://docs.python.org/3/c-api/init.html#c.Py_GetProgramFullPath
func Py_GetProgramFullPath() string {
	wcname := C.Py_GetProgramFullPath()
	cname := C.Py_EncodeLocale(wcname, nil)
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname)
}

//Py_GetPath : https://docs.python.org/3/c-api/init.html#c.Py_GetPath
func Py_GetPath() string {
	wcname := C.Py_GetPath()
	cname := C.Py_EncodeLocale(wcname, nil)
	defer C.PyMem_Free(unsafe.Pointer(cname))

	return C.GoString(cname)
}

//Py_SetPath : https://docs.python.org/3/c-api/init.html#c.Py_SetPath
func Py_SetPath(path string) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	wcpath := C.Py_DecodeLocale(cpath, nil)
	// wcpath is copied, safe to free after calling C.Py_SetPath
	defer C.PyMem_RawFree(unsafe.Pointer(wcpath))

	C.Py_SetPath(wcpath)
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

	wchome := C.Py_DecodeLocale(chome, nil)
	// wcpath is copied, safe to free after calling C.Py_SetPath
	defer C.PyMem_RawFree(unsafe.Pointer(wchome))

	C.Py_SetPythonHome(wchome)
}

//Py_GetPythonHome : https://docs.python.org/3/c-api/init.html#c.Py_GetPythonHome
func Py_GetPythonHome() string {
	wchome := C.Py_GetPythonHome()
	chome := C.Py_EncodeLocale(wchome, nil)
	defer C.PyMem_Free(unsafe.Pointer(chome))

	return C.GoString(chome)
}

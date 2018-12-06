package python3

/*
#cgo pkg-config: python3
#include "Python.h"
*/
import "C"
import (
	"unsafe"
)

//Py_Main : https://docs.python.org/3/c-api/veryhigh.html?highlight=pycompilerflags#c.Py_Main
func Py_Main(args []string) int {
	argc := C.int(len(args))
	argv := make([]*C.wchar_t, argc, argc)
	for i, arg := range args {
		carg := C.CString(arg)
		defer C.free(unsafe.Pointer(carg))

		warg := C.Py_DecodeLocale(carg, nil)
		if warg == nil {
			return -1
		}
		// Py_DecodeLocale requires a call to PyMem_RawFree to free the memory
		defer C.PyMem_RawFree(unsafe.Pointer(warg))
		argv[i] = warg
	}

	return int(C.Py_Main(argc, (**C.wchar_t)(unsafe.Pointer(&argv[0]))))
}

//PyRun_AnyFile : https://docs.python.org/3/c-api/veryhigh.html?highlight=pycompilerflags#c.PyRun_AnyFile
func PyRun_AnyFile(filename string) int {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	mode := C.CString("r")
	defer C.free(unsafe.Pointer(mode))

	cfile := C.fopen(cfilename, mode)
	if cfile == nil {
		return 1
	}
	defer C.fclose(cfile)

	// C.PyRun_AnyFile is a macro, using C.PyRun_AnyFileFlags instead
	return int(C.PyRun_AnyFileFlags(cfile, cfilename, nil))
}

//PyRun_SimpleString : https://docs.python.org/3/c-api/veryhigh.html?highlight=pycompilerflags#c.PyRun_SimpleString
func PyRun_SimpleString(command string) int {
	ccommand := C.CString(command)
	defer C.free(unsafe.Pointer(ccommand))

	// C.PyRun_SimpleString is a macro, using C.PyRun_SimpleStringFlags instead
	return int(C.PyRun_SimpleStringFlags(ccommand, nil))
}

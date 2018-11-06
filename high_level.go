package python3

/*
#cgo pkg-config: python3
#include "Python.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type PyCompilerFlags C.PyCompilerFlags

func Py_Main(args []string) error {
	argc := C.int(len(args))
	argv := make([]*C.wchar_t, argc, argc)
	for i, arg := range args {
		carg := C.CString(arg)
		defer C.free(unsafe.Pointer(carg))

		warg := C.Py_DecodeLocale(carg, nil)
		if warg == nil {
			return fmt.Errorf("Unable to convert command line to *wchar_t")
		}
		defer C.free(unsafe.Pointer(warg))

		argv[i] = warg
	}
	ret := C.Py_Main(argc, (**C.wchar_t)(unsafe.Pointer(&argv[0])))

	if ret == 1 {
		return fmt.Errorf("The interpreter exited due to an exception")
	}
	if ret == 2 {
		return fmt.Errorf("The parameter list does not represent a valid Python command line")
	}
	return nil
}

func PyRun_AnyFile(filename string) error {
	return PyRun_AnyFileFlags(filename, nil)
}

func PyRun_AnyFileFlags(filename string, flags *PyCompilerFlags) error {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	mode := C.CString("r")
	defer C.free(unsafe.Pointer(mode))

	cfile := C.fopen(cfilename, mode)
	if cfile == nil {
		return fmt.Errorf("File %s cannot be opened", filename)
	}

	ret := C.PyRun_AnyFileFlags(cfile, cfilename, (*C.PyCompilerFlags)(flags))

	if ret == -1 {
		return fmt.Errorf("An exception was raised during execution")
	}
	return nil
}

func PyRun_SimpleString(command string) error {
	return PyRun_SimpleStringFlags(command, nil)
}

func PyRun_SimpleStringFlags(command string, flags *PyCompilerFlags) error {
	ccommand := C.CString(command)
	defer C.free(unsafe.Pointer(ccommand))

	ret := C.PyRun_SimpleStringFlags(ccommand, (*C.PyCompilerFlags)(flags))

	if ret == -1 {
		return fmt.Errorf("An exception was raised during execution")
	}
	return nil
}

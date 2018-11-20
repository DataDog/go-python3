package python3

/*
#include "Python.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//PySys_GetObject : https://docs.python.org/3/c-api/sys.html#c.PySys_GetObject
func PySys_GetObject(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PySys_GetObject(cname))
}

//PySys_SetObject : https://docs.python.org/3/c-api/sys.html#c.PySys_SetObject
func PySys_SetObject(name string, v *PyObject) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	if C.PySys_SetObject(cname, toc(v)) == -1 {
		return fmt.Errorf("PySys_SetObject: an error occured")
	}
	return nil
}

//PySys_ResetWarnOptions : https://docs.python.org/3/c-api/sys.html#c.PySys_ResetWarnOptions
func PySys_ResetWarnOptions() {
	C.PySys_ResetWarnOptions()
}

//PySys_AddWarnOption : https://docs.python.org/3/c-api/sys.html#c.PySys_AddWarnOption
func PySys_AddWarnOption(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	ws := C.Py_DecodeLocale(cs, nil)
	defer C.PyMem_RawFree(unsafe.Pointer(ws))

	C.PySys_AddWarnOption(ws)
}

//PySys_AddWarnOptionUnicode : https://docs.python.org/3/c-api/sys.html#c.PySys_AddWarnOptionUnicode
func PySys_AddWarnOptionUnicode(unicode *PyObject) {
	C.PySys_AddWarnOptionUnicode(toc(unicode))
}

//PySys_SetPath : https://docs.python.org/3/c-api/sys.html#c.PySys_SetPath
func PySys_SetPath(path string) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	wpath := C.Py_DecodeLocale(cpath, nil)
	defer C.PyMem_RawFree(unsafe.Pointer(wpath))

	C.PySys_SetPath(wpath)
}

//PySys_AddXOption : https://docs.python.org/3/c-api/sys.html#c.PySys_AddXOption
func PySys_AddXOption(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	ws := C.Py_DecodeLocale(cs, nil)
	defer C.PyMem_RawFree(unsafe.Pointer(ws))

	C.PySys_AddXOption(ws)
}

//PySys_GetXOptions : https://docs.python.org/3/c-api/sys.html#c.PySys_GetXOptions
func PySys_GetXOptions() *PyObject {
	return togo(C.PySys_GetXOptions())
}

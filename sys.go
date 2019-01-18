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

//PySys_GetObject : https://docs.python.org/3/c-api/sys.html#c.PySys_GetObject
func PySys_GetObject(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PySys_GetObject(cname))
}

//PySys_SetObject : https://docs.python.org/3/c-api/sys.html#c.PySys_SetObject
func PySys_SetObject(name string, v *PyObject) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return int(C.PySys_SetObject(cname, toc(v)))
}

//PySys_ResetWarnOptions : https://docs.python.org/3/c-api/sys.html#c.PySys_ResetWarnOptions
func PySys_ResetWarnOptions() {
	C.PySys_ResetWarnOptions()
}

//PySys_AddWarnOption : https://docs.python.org/3/c-api/sys.html#c.PySys_AddWarnOption
func PySys_AddWarnOption(s string) error {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	ws := C.Py_DecodeLocale(cs, nil)
	if ws == nil {
		return fmt.Errorf("fail to call Py_DecodeLocale on '%s'", s)
	}
	defer C.PyMem_RawFree(unsafe.Pointer(ws))

	C.PySys_AddWarnOption(ws)

	return nil
}

//PySys_SetPath : https://docs.python.org/3/c-api/sys.html#c.PySys_SetPath
func PySys_SetPath(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	wpath := C.Py_DecodeLocale(cpath, nil)
	if wpath == nil {
		return fmt.Errorf("fail to call Py_DecodeLocale on '%s'", path)
	}
	defer C.PyMem_RawFree(unsafe.Pointer(wpath))

	C.PySys_SetPath(wpath)

	return nil
}

//PySys_AddXOption : https://docs.python.org/3/c-api/sys.html#c.PySys_AddXOption
func PySys_AddXOption(s string) error {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	ws := C.Py_DecodeLocale(cs, nil)
	if ws == nil {
		return fmt.Errorf("fail to call Py_DecodeLocale on '%s'", s)
	}
	defer C.PyMem_RawFree(unsafe.Pointer(ws))

	C.PySys_AddXOption(ws)

	return nil
}

//PySys_GetXOptions : https://docs.python.org/3/c-api/sys.html#c.PySys_GetXOptions
func PySys_GetXOptions() *PyObject {
	return togo(C.PySys_GetXOptions())
}

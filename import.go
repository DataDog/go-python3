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
	"unsafe"
)

//PyImport_ImportModule : https://docs.python.org/3/c-api/import.html#c.PyImport_ImportModule
func PyImport_ImportModule(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyImport_ImportModule(cname))
}

//PyImport_ImportModuleEx : https://docs.python.org/3/c-api/import.html#c.PyImport_ImportModuleEx
func PyImport_ImportModuleEx(name string, globals, locals, fromlist *PyObject) *PyObject {
	return PyImport_ImportModuleLevel(name, globals, locals, fromlist, 0)
}

//PyImport_ImportModuleLevelObject : https://docs.python.org/3/c-api/import.html#c.PyImport_ImportModuleLevelObject
func PyImport_ImportModuleLevelObject(name, globals, locals, fromlist *PyObject, level int) *PyObject {
	return togo(C.PyImport_ImportModuleLevelObject(toc(name), toc(globals), toc(locals), toc(fromlist), C.int(level)))
}

//PyImport_ImportModuleLevel : https://docs.python.org/3/c-api/import.html#c.PyImport_ImportModuleLevel
func PyImport_ImportModuleLevel(name string, globals, locals, fromlist *PyObject, level int) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyImport_ImportModuleLevel(cname, toc(globals), toc(locals), toc(fromlist), C.int(level)))
}

//PyImport_Import : https://docs.python.org/3/c-api/import.html#c.PyImport_Import
func PyImport_Import(name *PyObject) *PyObject {
	return togo(C.PyImport_Import(toc(name)))
}

//PyImport_ReloadModule : https://docs.python.org/3/c-api/import.html#c.PyImport_ReloadModule
func PyImport_ReloadModule(name *PyObject) *PyObject {
	return togo(C.PyImport_ReloadModule(toc(name)))
}

//PyImport_AddModuleObject : https://docs.python.org/3/c-api/import.html#c.PyImport_AddModuleObject
func PyImport_AddModuleObject(name *PyObject) *PyObject {
	return togo(C.PyImport_AddModuleObject(toc(name)))
}

//PyImport_AddModule : https://docs.python.org/3/c-api/import.html#c.PyImport_AddModule
func PyImport_AddModule(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyImport_AddModule(cname))
}

//PyImport_ExecCodeModule : https://docs.python.org/3/c-api/import.html#c.PyImport_ExecCodeModule
func PyImport_ExecCodeModule(name string, co *PyObject) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyImport_ExecCodeModule(cname, toc(co)))
}

//PyImport_ExecCodeModuleEx : https://docs.python.org/3/c-api/import.html#c.PyImport_ExecCodeModuleEx
func PyImport_ExecCodeModuleEx(name string, co *PyObject, pathname string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cpathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(cpathname))

	return togo(C.PyImport_ExecCodeModuleEx(cname, toc(co), cpathname))
}

//PyImport_ExecCodeModuleObject : https://docs.python.org/3/c-api/import.html#c.PyImport_ExecCodeModuleObject
func PyImport_ExecCodeModuleObject(name, co, pathname, cpathname *PyObject) *PyObject {
	return togo(C.PyImport_ExecCodeModuleObject(toc(name), toc(co), toc(pathname), toc(cpathname)))
}

//PyImport_ExecCodeModuleWithPathnames : https://docs.python.org/3/c-api/import.html#c.PyImport_ExecCodeModuleWithPathnames
func PyImport_ExecCodeModuleWithPathnames(name string, co *PyObject, pathname string, cpathname string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cspathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(cspathname))

	ccpathname := C.CString(cpathname)
	defer C.free(unsafe.Pointer(ccpathname))

	return togo(C.PyImport_ExecCodeModuleWithPathnames(cname, toc(co), cspathname, ccpathname))
}

//PyImport_GetMagicNumber : https://docs.python.org/3/c-api/import.html#c.PyImport_GetMagicNumber
func PyImport_GetMagicNumber() int {
	return int(C.PyImport_GetMagicNumber())
}

//PyImport_GetMagicTag : https://docs.python.org/3/c-api/import.html#c.PyImport_GetMagicTag
func PyImport_GetMagicTag() string {
	cmagicTag := C.PyImport_GetMagicTag()

	return C.GoString(cmagicTag)
}

//PyImport_GetModuleDict : https://docs.python.org/3/c-api/import.html#c.PyImport_GetModuleDict
func PyImport_GetModuleDict() *PyObject {
	return togo(C.PyImport_GetModuleDict())
}

//PyImport_GetModule : https://docs.python.org/3/c-api/import.html#c.PyImport_GetModule
func PyImport_GetModule(name *PyObject) *PyObject {
	return togo(C.PyImport_GetModule(toc(name)))

}

//PyImport_GetImporter : https://docs.python.org/3/c-api/import.html#c.PyImport_GetImporter
func PyImport_GetImporter(path *PyObject) *PyObject {
	return togo(C.PyImport_GetImporter(toc(path)))

}

//PyImport_ImportFrozenModuleObject : https://docs.python.org/3/c-api/import.html#c.PyImport_ImportFrozenModuleObject
func PyImport_ImportFrozenModuleObject(name *PyObject) int {
	return int(C.PyImport_ImportFrozenModuleObject(toc(name)))

}

//PyImport_ImportFrozenModule : https://docs.python.org/3/c-api/import.html#c.PyImport_ImportFrozenModule
func PyImport_ImportFrozenModule(name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return int(C.PyImport_ImportFrozenModule(cname))

}

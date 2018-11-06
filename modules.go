package python3

/*
#include "Python.h"
*/
import "C"

import (
	"unsafe"
)

func PyImport_ImportModule(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return (*PyObject)(C.PyImport_ImportModule(cname))
}

func PyImport_ImportModuleEx(name string, globals, locals, fromlist *PyObject) *PyObject {
	return PyImport_ImportModuleLevel(name, globals, locals, fromlist, 0)
}

func PyImport_ImportModuleLevelObject(name, globals, locals, fromlist *PyObject, level int) *PyObject {

	return (*PyObject)(C.PyImport_ImportModuleLevelObject((*C.PyObject)(name), (*C.PyObject)(globals), (*C.PyObject)(locals), (*C.PyObject)(fromlist), C.int(level)))
}

func PyImport_ImportModuleLevel(name string, globals, locals, fromlist *PyObject, level int) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return (*PyObject)(C.PyImport_ImportModuleLevel(cname, (*C.PyObject)(globals), (*C.PyObject)(locals), (*C.PyObject)(fromlist), C.int(level)))
}

func PyImport_Import(name *PyObject) *PyObject {
	return (*PyObject)(C.PyImport_Import((*C.PyObject)(name)))
}

func PyImport_ReloadModule(name *PyObject) *PyObject {
	return (*PyObject)(C.PyImport_ReloadModule((*C.PyObject)(name)))
}

func PyImport_AddModuleObject(name *PyObject) *PyObject {
	return (*PyObject)(C.PyImport_AddModuleObject((*C.PyObject)(name)))
}

func PyImport_AddModule(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return (*PyObject)(C.PyImport_AddModule(cname))
}

func PyImport_ExecCodeModule(name string, co *PyObject) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return (*PyObject)(C.PyImport_ExecCodeModule(cname, (*C.PyObject)(co)))
}

func PyImport_ExecCodeModuleEx(name string, co *PyObject, pathname string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cpathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(cpathname))

	return (*PyObject)(C.PyImport_ExecCodeModuleEx(cname, (*C.PyObject)(co), cpathname))
}

func PyImport_ExecCodeModuleObject(name, co, pathname, cpathname *PyObject) *PyObject {
	return (*PyObject)(C.PyImport_ExecCodeModuleObject((*C.PyObject)(name), (*C.PyObject)(co), (*C.PyObject)(pathname), (*C.PyObject)(cpathname)))
}

func PyImport_ExecCodeModuleWithPathnames(name string, co *PyObject, pathname string, cpathname string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cspathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(cspathname))

	ccpathname := C.CString(cpathname)
	defer C.free(unsafe.Pointer(ccpathname))

	return (*PyObject)(C.PyImport_ExecCodeModuleWithPathnames(cname, (*C.PyObject)(co), cspathname, ccpathname))
}

func PyImport_GetMagicNumber() int {
	return int(C.PyImport_GetMagicNumber())
}

func PyImport_GetMagicTag() string {
	cmagicTag := C.PyImport_GetMagicTag()
	defer C.free(unsafe.Pointer(cmagicTag))

	return C.GoString(cmagicTag)
}

func PyImport_GetModuleDict() *PyObject {
	return (*PyObject)(C.PyImport_GetModuleDict())
}

func PyImport_GetModule(name *PyObject) *PyObject {
	return (*PyObject)(C.PyImport_GetModule((*C.PyObject)(name)))

}

func PyImport_GetImporter(path *PyObject) *PyObject {
	return (*PyObject)(C.PyImport_GetImporter((*C.PyObject)(path)))

}

func PyImport_ImportFrozenModuleObject(name *PyObject) int {
	return int(C.PyImport_ImportFrozenModuleObject((*C.PyObject)(name)))

}

func PyImport_ImportFrozenModule(name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return int(C.PyImport_ImportFrozenModule(cname))

}

/*
Unless explicitly stated otherwise all files in this repository are licensed
under the $license_for_repo License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"
import "unsafe"

//Module : https://docs.python.org/3/c-api/module.html#c.PyModule_Type
var Module = togo((*C.PyObject)(unsafe.Pointer(&C.PyModule_Type)))

//PyModule_Check : https://docs.python.org/3/c-api/module.html#c.PyModule_Check
func PyModule_Check(p *PyObject) bool {
	return C._go_PyModule_Check(toc(p)) != 0
}

//PyModule_CheckExact : https://docs.python.org/3/c-api/module.html#c.PyModule_CheckExact
func PyModule_CheckExact(p *PyObject) bool {
	return C._go_PyModule_CheckExact(toc(p)) != 0
}

//PyModule_NewObject : https://docs.python.org/3/c-api/module.html#c.PyModule_NewObject
func PyModule_NewObject(name *PyObject) *PyObject {
	return togo(C.PyModule_NewObject(toc(name)))
}

//PyModule_New : https://docs.python.org/3/c-api/module.html#c.PyModule_New
func PyModule_New(name string) *PyObject {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return togo(C.PyModule_New(cname))
}

//PyModule_GetDict : https://docs.python.org/3/c-api/module.html#c.PyModule_GetDict
func PyModule_GetDict(module *PyObject) *PyObject {
	return togo(C.PyModule_GetDict(toc(module)))
}

//PyModule_GetNameObject : https://docs.python.org/3/c-api/module.html#c.PyModule_GetNameObject
func PyModule_GetNameObject(module *PyObject) *PyObject {
	return togo(C.PyModule_GetNameObject(toc(module)))
}

//PyModule_GetName : https://docs.python.org/3/c-api/module.html#c.PyModule_GetName
func PyModule_GetName(module *PyObject) string {
	cname := C.PyModule_GetName(toc(module))
	return C.GoString(cname)
}

//PyModule_GetState : https://docs.python.org/3/c-api/module.html#c.PyModule_GetState
func PyModule_GetState(module *PyObject) unsafe.Pointer {
	return unsafe.Pointer(C.PyModule_GetNameObject(toc(module)))
}

//PyModule_GetFilenameObject : https://docs.python.org/3/c-api/module.html#c.PyModule_GetFilenameObject
func PyModule_GetFilenameObject(module *PyObject) *PyObject {
	return togo(C.PyModule_GetFilenameObject(toc(module)))
}

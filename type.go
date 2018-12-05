package python3

/*
#include "Python.h"
#include "macro.h"
#include "type.h"
*/
import "C"

//Type : https://docs.python.org/3/c-api/type.html#c.PyType_Type
var Type = togo(C._go_PyType_Type)

//PyType_Check : https://docs.python.org/3/c-api/type.html#c.PyType_Check
func PyType_Check(o *PyObject) bool {
	return C._go_PyType_Check(toc(o)) != 0
}

//PyType_CheckExact : https://docs.python.org/3/c-api/type.html#c.PyType_CheckExact
func PyType_CheckExact(o *PyObject) bool {
	return C._go_PyType_CheckExact(toc(o)) != 0
}

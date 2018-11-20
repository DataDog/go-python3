package python3

/*
#include "Python.h"
#include "macro.h"
*/
import "C"

//PyType_Check : https://docs.python.org/3/c-api/type.html#c.PyType_Check
func PyType_Check(p *PyObject) bool {
	return C._PyType_Check(toc(p)) != 0
}

//PyType_CheckExact : https://docs.python.org/3/c-api/type.html#c.PyType_CheckExact
func PyType_CheckExact(p *PyObject) bool {
	return C._PyType_CheckExact(toc(p)) != 0
}

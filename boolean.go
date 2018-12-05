package python3

/*
#include "Python.h"
#include "macro.h"
#include "type.h"
*/
import "C"

//python boolean constants
var (
	False = togo(C.Py_False)
	True  = togo(C.Py_True)
)

//Bool : https://docs.python.org/3/c-api/bool.html#c.PyBool_Type
var Bool = togo(C._go_PyBool_Type)

//PyBool_Check : https://docs.python.org/3/c-api/bool.html#c.PyBool_Check
func PyBool_Check(o *PyObject) bool {
	return C._go_PyBool_Check(toc(o)) != 0
}

//PyBool_FromLong : https://docs.python.org/3/c-api/bool.html#c.PyBool_FromLong
func PyBool_FromLong(v int) *PyObject {
	return togo(C.PyBool_FromLong(C.long(v)))
}

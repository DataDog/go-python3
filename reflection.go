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

//PyEval_GetBuiltins : https://docs.python.org/3/c-api/reflection.html?highlight=reflection#c.PyEval_GetBuiltins
func PyEval_GetBuiltins() *PyObject {
	return togo(C.PyEval_GetBuiltins())
}

//PyEval_GetLocals : https://docs.python.org/3/c-api/reflection.html?highlight=reflection#c.PyEval_GetLocals
func PyEval_GetLocals() *PyObject {
	return togo(C.PyEval_GetLocals())
}

//PyEval_GetGlobals : https://docs.python.org/3/c-api/reflection.html?highlight=reflection#c.PyEval_GetGlobals
func PyEval_GetGlobals() *PyObject {
	return togo(C.PyEval_GetGlobals())
}

//PyEval_GetFuncName : https://docs.python.org/3/c-api/reflection.html?highlight=reflection#c.PyEval_GetFuncName
func PyEval_GetFuncName(pyFunc *PyObject) string {
	return C.GoString(C.PyEval_GetFuncName(toc(pyFunc)))
}

//PyEval_GetFuncDesc : https://docs.python.org/3/c-api/reflection.html?highlight=reflection#c.PyEval_GetFuncDesc
func PyEval_GetFuncDesc(pyFunc *PyObject) string {
	return C.GoString(C.PyEval_GetFuncDesc(toc(pyFunc)))
}

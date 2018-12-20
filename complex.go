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

//Complex : https://docs.python.org/3/c-api/complex.html#c.PyComplex_Type
var Complex = togo((*C.PyObject)(unsafe.Pointer(&C.PyComplex_Type)))

//PyComplex_Check : https://docs.python.org/3/c-api/complex.html#c.PyComplex_Check
func PyComplex_Check(p *PyObject) bool {
	return C._go_PyComplex_Check(toc(p)) != 0
}

//PyComplex_CheckExact : https://docs.python.org/3/c-api/complex.html#c.PyComplex_CheckExact
func PyComplex_CheckExact(p *PyObject) bool {
	return C._go_PyComplex_CheckExact(toc(p)) != 0
}

//PyComplex_FromDoubles : https://docs.python.org/3/c-api/complex.html#c.PyComplex_FromDoubles
func PyComplex_FromDoubles(real, imag float64) *PyObject {
	return togo(C.PyComplex_FromDoubles(C.double(real), C.double(imag)))
}

//PyComplex_RealAsDouble : https://docs.python.org/3/c-api/complex.html#c.PyComplex_RealAsDouble
func PyComplex_RealAsDouble(op *PyObject) float64 {
	return float64(C.PyComplex_RealAsDouble(toc(op)))
}

//PyComplex_ImagAsDouble : https://docs.python.org/3/c-api/complex.html#c.PyComplex_ImagAsDouble
func PyComplex_ImagAsDouble(op *PyObject) float64 {
	return float64(C.PyComplex_ImagAsDouble(toc(op)))
}

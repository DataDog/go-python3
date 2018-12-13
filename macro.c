/*
Unless explicitly stated otherwise all files in this repository are licensed
under the $license_for_repo License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

#include "macro.h"

int _go_Py_EnterRecursiveCall(const char *where) {
    return Py_EnterRecursiveCall(where);
}

void _go_Py_LeaveRecursiveCall() {
    Py_LeaveRecursiveCall();
}

int _go_PyType_Check(PyObject *o) {
    return PyType_Check(o);
}
int _go_PyType_CheckExact(PyObject *o) {
    return PyType_CheckExact(o);
}



int _go_PyLong_Check(PyObject *p) {
    return PyLong_Check(p);
}
int _go_PyLong_CheckExact(PyObject *p) {
    return PyLong_CheckExact(p);
}

int _go_PyBool_Check(PyObject *o) {
    return PyBool_Check(o);
}

int _go_PyFloat_Check(PyObject *p) {
    return PyFloat_Check(p);
}

int _go_PyFloat_CheckExact(PyObject *p) {
    return PyFloat_CheckExact(p);
}

int _go_PyComplex_Check(PyObject *p) {
    return PyComplex_Check(p);
}

int _go_PyComplex_CheckExact(PyObject *p) {
    return PyComplex_CheckExact(p);
}

int _go_PyBytes_Check(PyObject *o) {
    return PyBytes_Check(o);
}
int _go_PyBytes_CheckExact(PyObject *o) {
    return PyBytes_CheckExact(o);
}

int _go_PyByteArray_Check(PyObject *o) {
    return PyByteArray_Check(o);
}
int _go_PyByteArray_CheckExact(PyObject *o) {
    return PyByteArray_CheckExact(o);
}

int _go_PyUnicode_Check(PyObject *o) {
    return PyUnicode_Check(o);
}
int _go_PyUnicode_CheckExact(PyObject *o) {
    return PyUnicode_CheckExact(o);
}

int _go_PyTuple_Check(PyObject *p) {
    return PyTuple_Check(p);
}
int _go_PyTuple_CheckExact(PyObject *p) {
    return PyTuple_CheckExact(p);
}

int _go_PyList_Check(PyObject *p) {
    return PyList_Check(p);
}
int _go_PyList_CheckExact(PyObject *p) {
    return PyList_CheckExact(p);
}

int _go_PyDict_Check(PyObject *p) {
    return PyDict_Check(p);
}
int _go_PyDict_CheckExact(PyObject *p) {
    return PyDict_CheckExact(p);
}

int _go_PyModule_Check(PyObject *p) {
    return PyModule_Check(p);
}
int _go_PyModule_CheckExact(PyObject *p) {
    return PyModule_CheckExact(p);
}

int _go_PyObject_DelAttr(PyObject *o, PyObject *attr_name) {
    return PyObject_DelAttr(o, attr_name);
}
int _go_PyObject_DelAttrString(PyObject *o, const char *attr_name) {
    return PyObject_DelAttrString(o, attr_name);
}

int _go_PyObject_TypeCheck(PyObject *o, PyTypeObject *type) {
    return PyObject_TypeCheck(o, type);
}
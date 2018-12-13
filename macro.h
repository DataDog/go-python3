/*
Unless explicitly stated otherwise all files in this repository are licensed
under the $license_for_repo License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

#ifndef MACRO_H
#define MACRO_H

#include "Python.h"

int _go_Py_EnterRecursiveCall(const char *where);
void _go_Py_LeaveRecursiveCall();

int _go_PyType_Check(PyObject *o);
int _go_PyType_CheckExact(PyObject *o);

int _go_PyLong_Check(PyObject *p);
int _go_PyLong_CheckExact(PyObject *p);

int _go_PyBool_Check(PyObject *o);

int _go_PyFloat_Check(PyObject *p);
int _go_PyFloat_CheckExact(PyObject *p);

int _go_PyComplex_Check(PyObject *p);
int _go_PyComplex_CheckExact(PyObject *p);

int _go_PyBytes_Check(PyObject *o);
int _go_PyBytes_CheckExact(PyObject *o);

int _go_PyByteArray_Check(PyObject *o);
int _go_PyByteArray_CheckExact(PyObject *o);

int _go_PyUnicode_Check(PyObject *o);
int _go_PyUnicode_CheckExact(PyObject *o);

int _go_PyTuple_Check(PyObject *p);
int _go_PyTuple_CheckExact(PyObject *p);

int _go_PyList_Check(PyObject *p);
int _go_PyList_CheckExact(PyObject *p);

int _go_PyDict_Check(PyObject *p);
int _go_PyDict_CheckExact(PyObject *p);

int _go_PyModule_Check(PyObject *p);
int _go_PyModule_CheckExact(PyObject *p);

int _go_PyObject_DelAttr(PyObject *o, PyObject *attr_name);
int _go_PyObject_DelAttrString(PyObject *o, const char *attr_name);

int _go_PyObject_TypeCheck(PyObject *o, PyTypeObject *type);

#endif
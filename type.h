/*
Unless explicitly stated otherwise all files in this repository are licensed
under the $license_for_repo License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

#ifndef TYPE_H
#define TYPE_H

#include "Python.h"

extern PyObject *_go_PyType_Type;
extern PyObject *_go_PyLong_Type;
extern PyObject *_go_PyBool_Type;
extern PyObject *_go_PyFloat_Type;
extern PyObject *_go_PyComplex_Type;

extern PyObject *_go_PyBytes_Type;
extern PyObject *_go_PyByteArray_Type;
extern PyObject *_go_PyUnicode_Type;
extern PyObject *_go_PyTuple_Type;
extern PyObject *_go_PyList_Type;

extern PyObject *_go_PyDict_Type;

extern PyObject *_go_PyModule_Type;

#endif

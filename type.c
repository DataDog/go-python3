/*
Unless explicitly stated otherwise all files in this repository are licensed
under the $license_for_repo License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

#include "type.h"

PyObject *_go_PyType_Type = (PyObject *)&PyType_Type;
PyObject *_go_PyLong_Type = (PyObject *)&PyLong_Type;
PyObject *_go_PyBool_Type = (PyObject *)&PyBool_Type;
PyObject *_go_PyFloat_Type = (PyObject *)&PyFloat_Type;
PyObject *_go_PyComplex_Type = (PyObject *)&PyComplex_Type;

PyObject *_go_PyBytes_Type = (PyObject *)&PyBytes_Type;
PyObject *_go_PyByteArray_Type = (PyObject *)&PyByteArray_Type;
PyObject *_go_PyUnicode_Type = (PyObject *)&PyUnicode_Type;
PyObject *_go_PyTuple_Type = (PyObject *)&PyTuple_Type;
PyObject *_go_PyList_Type = (PyObject *)&PyList_Type;

PyObject *_go_PyDict_Type = (PyObject *)&PyDict_Type;

PyObject *_go_PyModule_Type = (PyObject *)&PyModule_Type;

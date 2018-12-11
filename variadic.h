#ifndef VARIADIC_H
#define VARIADIC_H

#include "Python.h"

PyObject* _go_PyObject_CallFunctionObjArgs(PyObject *callable, int argc, PyObject **args);
PyObject* _go_PyObject_CallMethodObjArgs(PyObject *obj, PyObject *name, int argc, PyObject **args);

#endif
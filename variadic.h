#ifndef VARIADIC_H
#define VARIADIC_H

#include "Python.h"

PyObject* _PyObject_CallFunctionObjArgs(PyObject *callable, int argc, PyObject **args);
PyObject* _PyObject_CallMethodObjArgs(PyObject *obj, PyObject *name, int argc, PyObject **args);

#endif
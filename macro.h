#include "Python.h"

int _Py_EnterRecursiveCall(const char *where);
void _Py_LeaveRecursiveCall();

int _PyLong_Check(PyObject *p);
int _PyLong_CheckExact(PyObject *p);

int _PyBool_Check(PyObject *o);

int _PyFloat_Check(PyObject *p);
int _PyFloat_CheckExact(PyObject *p);

int _PyComplex_Check(PyObject *p);
int _PyComplex_CheckExact(PyObject *p);
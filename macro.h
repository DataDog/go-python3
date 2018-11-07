#include "Python.h"

int _Py_EnterRecursiveCall(const char *where);
void _Py_LeaveRecursiveCall();

int _PyLong_Check(PyObject *p);
int _PyLong_CheckExact(PyObject *p);
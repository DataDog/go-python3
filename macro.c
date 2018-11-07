#include "macro.h"

int _Py_EnterRecursiveCall(const char *where) {
    return Py_EnterRecursiveCall(where);
}

void _Py_LeaveRecursiveCall() {
    Py_LeaveRecursiveCall();
}

int _PyLong_Check(PyObject *p) {
    return PyLong_Check(p);
}
int _PyLong_CheckExact(PyObject *p) {
    return PyLong_CheckExact(p);
}

int _PyBool_Check(PyObject *o) {
    return PyBool_Check(o);
}

int _PyFloat_Check(PyObject *p) {
    return PyFloat_Check(p);
}

int _PyFloat_CheckExact(PyObject *p) {
    return PyFloat_CheckExact(p);
}

int _PyComplex_Check(PyObject *p) {
    return PyComplex_Check(p);
}

int _PyComplex_CheckExact(PyObject *p) {
    return PyComplex_CheckExact(p);
}
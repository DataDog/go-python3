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
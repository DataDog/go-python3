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

int _PyBytes_Check(PyObject *o) {
    return PyBytes_Check(o);
}
int _PyBytes_CheckExact(PyObject *o) {
    return PyBytes_CheckExact(o);
}

int _PyByteArray_Check(PyObject *o) {
    return PyByteArray_Check(o);
}
int _PyByteArray_CheckExact(PyObject *o) {
    return PyByteArray_CheckExact(o);
}

int _PyUnicode_Check(PyObject *o) {
    return PyUnicode_Check(o);
}
int _PyUnicode_CheckExact(PyObject *o) {
    return PyUnicode_CheckExact(o);
}

int _PyTuple_Check(PyObject *p) {
    return _PyTuple_Check(p);
}
int _PyTuple_CheckExact(PyObject *p) {
    return _PyTuple_CheckExact(p);
}

int _PyObject_DelAttr(PyObject *o, PyObject *attr_name) {
    return PyObject_DelAttr(o, attr_name);
}
int _PyObject_DelAttrString(PyObject *o, const char *attr_name) {
    return PyObject_DelAttrString(o, attr_name);
}

int _PyObject_TypeCheck(PyObject *o, PyTypeObject *type) {
    return PyObject_TypeCheck(o, type);
}
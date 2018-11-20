#include "Python.h"

int _Py_EnterRecursiveCall(const char *where);
void _Py_LeaveRecursiveCall();

int _PyType_Check(PyObject *o);
int _PyType_CheckExact(PyObject *o);

int _PyLong_Check(PyObject *p);
int _PyLong_CheckExact(PyObject *p);

int _PyBool_Check(PyObject *o);

int _PyFloat_Check(PyObject *p);
int _PyFloat_CheckExact(PyObject *p);

int _PyComplex_Check(PyObject *p);
int _PyComplex_CheckExact(PyObject *p);

int _PyBytes_Check(PyObject *o);
int _PyBytes_CheckExact(PyObject *o);

int _PyByteArray_Check(PyObject *o);
int _PyByteArray_CheckExact(PyObject *o);

int _PyUnicode_Check(PyObject *o);
int _PyUnicode_CheckExact(PyObject *o);

int _PyTuple_Check(PyObject *p);
int _PyTuple_CheckExact(PyObject *p);

int _PyList_Check(PyObject *p);
int _PyList_CheckExact(PyObject *p);

int _PyDict_Check(PyObject *p);
int _PyDict_CheckExact(PyObject *p);

int _PyModule_Check(PyObject *p);
int _PyModule_CheckExact(PyObject *p);

int _PyObject_DelAttr(PyObject *o, PyObject *attr_name);
int _PyObject_DelAttrString(PyObject *o, const char *attr_name);

int _PyObject_TypeCheck(PyObject *o, PyTypeObject *type);
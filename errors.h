#include "Python.h"

typedef struct GoPyErr {
    PyObject *pyType;
    PyObject *value;
    PyObject *traceback;
} GoPyErr;


GoPyErr GoPyErr_Fetch();
GoPyErr GoPyErr_NormalizeException(PyObject *exc, PyObject *val, PyObject *tb);
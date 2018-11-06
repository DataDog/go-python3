#include "errors.h"

GoPyErr GoPyErr_Fetch() {
    GoPyErr goPyErr;
    
    PyErr_Fetch(&(goPyErr.pyType), &(goPyErr.value), &(goPyErr.traceback));

    return goPyErr;
}

GoPyErr GoPyErr_NormalizeException(PyObject *exc, PyObject *val, PyObject *tb) {
    PyErr_NormalizeException(&exc, &val, &tb);
    GoPyErr goPyErr = {exc, val, tb};
    return goPyErr;
}
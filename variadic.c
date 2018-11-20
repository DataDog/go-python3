#include "Python.h"

PyObject* _PyObject_CallFunctionObjArgs(PyObject *callable, int argc, PyObject **argv) {

    PyObject *result = NULL;
    switch (argc) {
        case 0:
            result = PyObject_CallFunctionObjArgs(callable);
            break;
        case 1:
            result = PyObject_CallFunctionObjArgs(callable, argv[0]);
            break;
        case 2:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1]);
            break;
        case 3:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2]);
            break;
        case 4:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3]);
            break;
        case 5:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4]);
            break;
        case 6:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5]);
            break;
        case 7:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6]);
            break;
        case 8:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7]);
            break;
        case 9:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8]);
            break;
        case 10:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9]);
            break;
        case 11:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10]);
            break;
        case 12:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11]);
            break;
        case 13:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12]);
            break;
        case 14:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13]);
            break;
        case 15:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14]);
            break;
        case 16:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15]);
            break;
        case 17:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16]);
            break;
        case 18:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17]);
            break;
        case 19:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17], argv[18]);
            break;
        case 20:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17], argv[18], argv[19]);
            break;
    }
    return result;
}
PyObject* _PyObject_CallMethodObjArgs(PyObject *obj, PyObject *name, int argc, PyObject **argv) {

    PyObject *result = NULL;
    switch (argc) {
        case 0:
            result = PyObject_CallMethodObjArgs(obj, name);
            break;
        case 1:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0]);
            break;
        case 2:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1]);
            break;
        case 3:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2]);
            break;
        case 4:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3]);
            break;
        case 5:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4]);
            break;
        case 6:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5]);
            break;
        case 7:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6]);
            break;
        case 8:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7]);
            break;
        case 9:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8]);
            break;
        case 10:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9]);
            break;
        case 11:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10]);
            break;
        case 12:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11]);
            break;
        case 13:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12]);
            break;
        case 14:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13]);
            break;
        case 15:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14]);
            break;
        case 16:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15]);
            break;
        case 17:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16]);
            break;
        case 18:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17]);
            break;
        case 19:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17], argv[18]);
            break;
        case 20:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17], argv[18], argv[19]);
            break;
    }
    return result;
}

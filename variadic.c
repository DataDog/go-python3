#include "Python.h"

PyObject* _PyObject_CallFunctionObjArgs(PyObject *callable, int argc, PyObject **argv) {

    PyObject *result = NULL;
    switch (argc) {
        case 0:
            result = PyObject_CallFunctionObjArgs(callable, NULL);
            break;
        case 1:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], NULL);
            break;
        case 2:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], NULL);
            break;
        case 3:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], NULL);
            break;
        case 4:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], NULL);
            break;
        case 5:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], NULL);
            break;
        case 6:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], NULL);
            break;
        case 7:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], NULL);
            break;
        case 8:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], NULL);
            break;
        case 9:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], NULL);
            break;
        case 10:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], NULL);
            break;
        case 11:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], NULL);
            break;
        case 12:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], NULL);
            break;
        case 13:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], NULL);
            break;
        case 14:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], NULL);
            break;
        case 15:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], NULL);
            break;
        case 16:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], NULL);
            break;
        case 17:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], NULL);
            break;
        case 18:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17], NULL);
            break;
        case 19:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17], argv[18], NULL);
            break;
        case 20:
            result = PyObject_CallFunctionObjArgs(callable, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17], argv[18], argv[19], NULL);
            break;
    }
    return result;
}
PyObject* _PyObject_CallMethodObjArgs(PyObject *obj, PyObject *name, int argc, PyObject **argv) {

    PyObject *result = NULL;
    switch (argc) {
        case 0:
            result = PyObject_CallMethodObjArgs(obj, name, NULL);
            break;
        case 1:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], NULL);
            break;
        case 2:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], NULL);
            break;
        case 3:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], NULL);
            break;
        case 4:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], NULL);
            break;
        case 5:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], NULL);
            break;
        case 6:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], NULL);
            break;
        case 7:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], NULL);
            break;
        case 8:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], NULL);
            break;
        case 9:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], NULL);
            break;
        case 10:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], NULL);
            break;
        case 11:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], NULL);
            break;
        case 12:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], NULL);
            break;
        case 13:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], NULL);
            break;
        case 14:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], NULL);
            break;
        case 15:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], NULL);
            break;
        case 16:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], NULL);
            break;
        case 17:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], NULL);
            break;
        case 18:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17], NULL);
            break;
        case 19:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17], argv[18], NULL);
            break;
        case 20:
            result = PyObject_CallMethodObjArgs(obj, name, argv[0], argv[1], argv[2], argv[3], argv[4], argv[5], argv[6], argv[7], argv[8], argv[9], argv[10], argv[11], argv[12], argv[13], argv[14], argv[15], argv[16], argv[17], argv[18], argv[19], NULL);
            break;
    }
    return result;
}

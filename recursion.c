#include "recursion.h"

int _Py_EnterRecursiveCall(const char *where) {
    return Py_EnterRecursiveCall(where);
}

void _Py_LeaveRecursiveCall() {
    Py_LeaveRecursiveCall();
}
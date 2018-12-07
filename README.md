go-python3
==========

Golang bindings for the C-API of CPython-3.

This package provides a ``go`` package named "python" under which most of the
``PyXYZ`` functions and macros of the public C-API of CPython have been
exposed. Theoretically, you should be able use https://docs.python.org/3/c-api
and know what to type in your ``go`` program.

API
---

Some functions mix go code and call to Python function. Those functions will
return and `int` and `error` type. The `int` represent the Python result code
and the `error` represent any issue from the Go layer.

Example:

`func PyRun_AnyFile(filename string)` open `filename` and then call CPython API
function `int PyRun_AnyFile(FILE *fp, const char *filename)`.

Therefore its signature is `func PyRun_AnyFile(filename string) (int, error)`,
the `int` represent the error code from the CPython `PyRun_AnyFile` function
and error will be set if we failed to open `filename`.

If an error is raise before calling th CPython function `int` default to `-1`.

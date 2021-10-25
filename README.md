# Go bindings for the CPython-3 C-API

**Currently supports python-3.7 only.**

This package provides a ``go`` package named "python" under which most of the
``PyXYZ`` functions and macros of the public C-API of CPython have been
exposed. Theoretically, you should be able use https://docs.python.org/3/c-api
and know what to type in your ``go`` program.

## relation to `DataDog/go-python3`

This project is a community maintained successor to [`DataDog/go-python3`](https://github.com/DataDog/go-python3), which will get archived in December 2021.

- If you use the Go package `github.com/DataDog/go-python3` in your code, you can use `github.com/go-python/cpy3` as a drop-in replacement. We intend to not introduce breaking changes.
- If you have unmerged PRs or open issues on `DataDog/go-python3`, please re-submit them here.

## relation to `sbinet/go-python`

This project was inspired by [`sbinet/go-python`](https://github.com/sbinet/go-python) (Go bindings for the CPython-2 C-API).

# Install

## Deps

We will need `pkg-config` and a working `python3.7` environment to build these
bindings. Make sure you have Python libraries and header files installed as
well (`python3.7-dev` on Debian or `python3-devel` on Centos for example)..

By default `pkg-config` will look at the `python3` library so if you want to
choose a specific version just symlink `python-X.Y.pc` to `python3.pc` or use
the `PKG_CONFIG_PATH` environment variable.

## Go get

Then simply `go get github.com/go-python/cpy3`

# API

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

Take a look at some [examples](examples) and this [tutorial blogpost](https://poweruser.blog/embedding-python-in-go-338c0399f3d5).

# Contributing

Contributions are welcome! See [details](CONTRIBUTING.md).  


# Community
Find us in [`#go-python`](https://gophers.slack.com/archives/C4FDJLLET) on [Gophers Slack](https://gophers.slack.com). ([infos](https://blog.gopheracademy.com/gophers-slack-community/) | [invite](https://invite.slack.golangbridge.org/))  
  
This project follows the [Go Community Code of Conduct](https://golang.org/conduct).
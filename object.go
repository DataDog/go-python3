package python3

//go:generate go run script/variadic.go

/*
#include "Python.h"
#include "macro.h"
#include "variadic.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

var (
	Py_LT = C.Py_LT
	Py_LE = C.Py_LE
	Py_EQ = C.Py_EQ
	Py_NE = C.Py_NE
	Py_GT = C.Py_GT
	Py_GE = C.Py_GE
)

//PyObject : https://docs.python.org/3/c-api/structures.html?highlight=pyobject#c.PyObject
type PyObject C.PyObject

//togo converts a *C.PyObject to a *PyObject
func togo(cobject *C.PyObject) *PyObject {
	return (*PyObject)(cobject)
}

func toc(object *PyObject) *C.PyObject {
	return (*C.PyObject)(object)
}

//IncRef : https://docs.python.org/3/c-api/refcounting.html#c.Py_INCREF
func (pyObject *PyObject) IncRef() {
	C.Py_IncRef(toc(pyObject))
}

//DecRef : https://docs.python.org/3/c-api/refcounting.html#c.Py_DECREF
func (pyObject *PyObject) DecRef() {
	C.Py_DecRef(toc(pyObject))
}

//ReprEnter : https://docs.python.org/3/c-api/exceptions.html#c.Py_ReprEnter
func (pyObject *PyObject) ReprEnter() (bool, error) {
	if ret := C.Py_ReprEnter(toc(pyObject)); ret < 0 {
		return false, fmt.Errorf("recursion limit is reached")
	} else if ret > 0 {
		return true, nil
	}
	return false, nil
}

//ReprLeave : https://docs.python.org/3/c-api/exceptions.html#c.Py_ReprLeave
func (pyObject *PyObject) ReprLeave() {
	C.Py_ReprLeave(toc(pyObject))
}

//HasAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_HasAttr
func (pyObject *PyObject) HasAttr(attr_name *PyObject) bool {
	return C.PyObject_HasAttr(toc(pyObject), toc(attr_name)) == 1
}

//HasAttrString : https://docs.python.org/3/c-api/object.html#c.PyObject_HasAttrString
func (pyObject *PyObject) HasAttrString(attr_name string) bool {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return C.PyObject_HasAttrString(toc(pyObject), cattr_name) == 1
}

//GetAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_GetAttr
func (pyObject *PyObject) GetAttr(attr_name *PyObject) *PyObject {
	return togo(C.PyObject_GetAttr(toc(pyObject), toc(attr_name)))
}

//GetAttrString : https://docs.python.org/3/c-api/object.html#c.PyObject_GetAttrString
func (pyObject *PyObject) GetAttrString(attr_name string) *PyObject {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	return togo(C.PyObject_GetAttrString(toc(pyObject), cattr_name))
}

//SetAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_SetAttr
func (pyObject *PyObject) SetAttr(attr_name *PyObject, v *PyObject) {

	C.PyObject_SetAttr(toc(pyObject), toc(attr_name), toc(v))
}

//SetAttrString : https://docs.python.org/3/c-api/object.html#c.PyObject_SetAttrString
func (pyObject *PyObject) SetAttrString(attr_name string, v *PyObject) error {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	if C.PyObject_SetAttrString(toc(pyObject), cattr_name, toc(v)) == -1 {
		return fmt.Errorf("PyObject_SetAttrString: an error occured")
	}
	return nil
}

//DelAttr : https://docs.python.org/3/c-api/object.html#c.PyObject_DelAttr
func (pyObject *PyObject) DelAttr(attr_name *PyObject) {
	C._PyObject_DelAttr(toc(pyObject), toc(attr_name))
}

//DelAttrString : https://docs.python.org/3/c-api/object.html#c.PyObject_DelAttrString
func (pyObject *PyObject) DelAttrString(attr_name string) {
	cattr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(cattr_name))

	C._PyObject_DelAttrString(toc(pyObject), cattr_name)
}

//RichCompare : https://docs.python.org/3/c-api/object.html#c.PyObject_RichCompare
func (pyObject *PyObject) RichCompare(o *PyObject, opid int) *PyObject {
	return togo(C.PyObject_RichCompare(toc(pyObject), toc(o), C.int(opid)))
}

//RichCompareBool : https://docs.python.org/3/c-api/object.html#c.PyObject_RichCompareBool
func (pyObject *PyObject) RichCompareBool(o *PyObject, opid int) (bool, error) {
	ret := C.PyObject_RichCompareBool(toc(pyObject), toc(o), C.int(opid))
	if ret == -1 {
		return false, fmt.Errorf("RichCompareBool: an error occured")
	}
	return ret == 1, nil
}

//Repr : https://docs.python.org/3/c-api/object.html#c.PyObject_Repr
func (pyObject *PyObject) Repr() *PyObject {
	return togo(C.PyObject_Repr(toc(pyObject)))
}

//ASCII : https://docs.python.org/3/c-api/object.html#c.PyObject_ASCII
func (pyObject *PyObject) ASCII() *PyObject {
	return togo(C.PyObject_ASCII(toc(pyObject)))
}

//Str : https://docs.python.org/3/c-api/object.html#c.PyObject_Str
func (pyObject *PyObject) Str() *PyObject {
	return togo(C.PyObject_Str(toc(pyObject)))
}

//Bytes : https://docs.python.org/3/c-api/object.html#c.PyObject_Bytes
func (pyObject *PyObject) Bytes() *PyObject {
	return togo(C.PyObject_Bytes(toc(pyObject)))
}

//IsSubclass : https://docs.python.org/3/c-api/object.html#c.PyObject_IsSubclass
func (pyObject *PyObject) IsSubclass(cls *PyObject) (bool, error) {
	ret := C.PyObject_IsSubclass(toc(pyObject), toc(cls))
	if ret == -1 {
		return false, fmt.Errorf("IsSubclass: an error occured")
	}
	return ret == 1, nil
}

//IsInstance : https://docs.python.org/3/c-api/object.html#c.PyObject_IsInstance
func (pyObject *PyObject) IsInstance(cls *PyObject) (bool, error) {
	ret := C.PyObject_IsInstance(toc(pyObject), toc(cls))
	if ret == -1 {
		return false, fmt.Errorf("IsInstance: an error occured")
	}
	return ret == 1, nil
}

// PyCallable_Check : https://docs.python.org/3/c-api/object.html#c.PyCallable_Check
func PyCallable_Check(o *PyObject) bool {
	return C.PyCallable_Check(toc(o)) == 1
}

//Call : https://docs.python.org/3/c-api/object.html#c.PyObject_Call
func (pyObject *PyObject) Call(args *PyObject, kwargs *PyObject) *PyObject {
	return togo(C.PyObject_Call(toc(pyObject), toc(args), toc(kwargs)))
}

//CallObject : https://docs.python.org/3/c-api/object.html#c.PyObject_CallObject
func (pyObject *PyObject) CallObject(args *PyObject) *PyObject {
	return togo(C.PyObject_CallObject(toc(pyObject), toc(args)))
}

//CallFunctionObjArgs : https://docs.python.org/3/c-api/object.html#c.PyObject_CallFunctionObjArgs
func (pyObject *PyObject) CallFunctionObjArgs(args ...*PyObject) (*PyObject, error) {

	if len(args) > 20 {
		return nil, fmt.Errorf("CallFunctionObjArgs: too many arguments")
	}
	cargs := make([]*C.PyObject, len(args), len(args))
	for i, arg := range args {
		cargs[i] = toc(arg)
	}
	return togo(C._PyObject_CallFunctionObjArgs(toc(pyObject), C.int(len(args)), (**C.PyObject)(unsafe.Pointer(&cargs[0])))), nil
}

//CallMethodObjArgs : https://docs.python.org/3/c-api/object.html#c.PyObject_CallMethodObjArgs
func (pyObject *PyObject) CallMethodObjArgs(name *PyObject, args ...*PyObject) (*PyObject, error) {
	if len(args) > 20 {
		return nil, fmt.Errorf("CallMethodObjArgs: too many arguments")
	}
	cargs := make([]*C.PyObject, len(args), len(args))
	for i, arg := range args {
		cargs[i] = toc(arg)
	}
	return togo(C._PyObject_CallMethodObjArgs(toc(pyObject), toc(name), C.int(len(args)), (**C.PyObject)(unsafe.Pointer(&cargs[0])))), nil
}

//Hash : https://docs.python.org/3/c-api/object.html#c.PyObject_Hash
func (pyObject *PyObject) Hash() (int, error) {
	ret := C.PyObject_Hash(toc(pyObject))
	if ret == -1 {
		return 0, fmt.Errorf("PyObject_Hash: an error occured")
	}
	return int(ret), nil
}

//HashNotImplemented : https://docs.python.org/3/c-api/object.html#c.PyObject_HashNotImplemented
func (pyObject *PyObject) HashNotImplemented() int {
	return int(C.PyObject_HashNotImplemented(toc(pyObject)))
}

//IsTrue : https://docs.python.org/3/c-api/object.html#c.PyObject_IsTrue
func (pyObject *PyObject) IsTrue() (bool, error) {
	ret := C.PyObject_IsTrue(toc(pyObject))
	if ret == -1 {
		return false, fmt.Errorf("PyObject_IsTrue: an error occured")
	}
	return ret == 1, nil
}

//Not : https://docs.python.org/3/c-api/object.html#c.PyObject_Not
func (pyObject *PyObject) Not() (bool, error) {
	ret := C.PyObject_Not(toc(pyObject))
	if ret == -1 {
		return false, fmt.Errorf("PyObject_Not: an error occured")
	}
	return ret == 1, nil
}

//Type : https://docs.python.org/3/c-api/object.html#c.PyObject_Type
func (pyObject *PyObject) Type() *PyObject {
	return togo(C.PyObject_Type(toc(pyObject)))
}

//Length : https://docs.python.org/3/c-api/object.html#c.PyObject_Length
func (pyObject *PyObject) Length() (int, error) {
	ret := C.PyObject_Length(toc(pyObject))
	if ret == -1 {
		return 0, fmt.Errorf("PyObject_Length: an error occured")
	}
	return int(ret), nil
}

//LengthHint : https://docs.python.org/3/c-api/object.html#c.PyObject_LengthHint
func (pyObject *PyObject) LengthHint(pyDefault int) (int, error) {
	ret := C.PyObject_LengthHint(toc(pyObject), C.Py_ssize_t(pyDefault))
	if ret == -1 {
		return 0, fmt.Errorf("PyObject_Length: an error occured")
	}
	return int(ret), nil
}

//GetItem : https://docs.python.org/3/c-api/object.html#c.PyObject_GetItem
func (pyObject *PyObject) GetItem(key *PyObject) *PyObject {
	return togo(C.PyObject_GetItem(toc(pyObject), toc(key)))
}

//SetItem : https://docs.python.org/3/c-api/object.html#c.PyObject_SetItem
func (pyObject *PyObject) SetItem(key, v *PyObject) error {
	if ret := C.PyObject_SetItem(toc(pyObject), toc(key), toc(v)); ret == 0 {
		return nil
	}
	return fmt.Errorf("PyObject_SetItem: an error occured")
}

//DelItem : https://docs.python.org/3/c-api/object.html#c.PyObject_DelItem
func (pyObject *PyObject) DelItem(key *PyObject) error {
	if ret := C.PyObject_DelItem(toc(pyObject), toc(key)); ret == 0 {
		return nil
	}
	return fmt.Errorf("PyObject_DelItem: an error occured")
}

//Dir : https://docs.python.org/3/c-api/object.html#c.PyObject_Dir
func (pyObject *PyObject) Dir() *PyObject {
	return togo(C.PyObject_Dir(toc(pyObject)))
}

//GetIter : https://docs.python.org/3/c-api/object.html#c.PyObject_GetIter
func (pyObject *PyObject) GetIter() *PyObject {
	return togo(C.PyObject_GetIter(toc(pyObject)))
}

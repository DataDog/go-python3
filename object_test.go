/*
Unless explicitly stated otherwise all files in this repository are licensed
under the MIT License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttrString(t *testing.T) {
	Py_Initialize()

	sys := PyImport_ImportModule("sys")
	defer sys.DecRef()

	assert.True(t, sys.HasAttrString("stdout"))
	stdout := sys.GetAttrString("stdout")
	assert.NotNil(t, stdout)

	assert.Zero(t, sys.DelAttrString("stdout"))

	assert.Nil(t, sys.GetAttrString("stdout"))
	PyErr_Clear()

	assert.Zero(t, sys.SetAttrString("stdout", stdout))

}

func TestAttr(t *testing.T) {
	Py_Initialize()

	name := PyUnicode_FromString("stdout")
	defer name.DecRef()

	sys := PyImport_ImportModule("sys")
	defer sys.DecRef()

	assert.True(t, sys.HasAttr(name))
	stdout := sys.GetAttr(name)
	assert.NotNil(t, stdout)

	assert.Zero(t, sys.DelAttr(name))

	assert.Nil(t, sys.GetAttr(name))
	PyErr_Clear()

	assert.Zero(t, sys.SetAttr(name, stdout))
}

func TestRichCompareBool(t *testing.T) {
	Py_Initialize()

	s1 := PyUnicode_FromString("test1")
	s2 := PyUnicode_FromString("test2")

	assert.Zero(t, s1.RichCompareBool(s2, Py_EQ))

	assert.NotZero(t, s1.RichCompareBool(s1, Py_EQ))

}

func TestRichCompare(t *testing.T) {
	Py_Initialize()

	s1 := PyUnicode_FromString("test1")
	s2 := PyUnicode_FromString("test2")

	b1 := s1.RichCompare(s2, Py_EQ)
	defer b1.DecRef()
	assert.Equal(t, Py_False, b1)

	b2 := s1.RichCompare(s1, Py_EQ)
	assert.Equal(t, Py_True, b2)
	defer b2.DecRef()

}

func TestRepr(t *testing.T) {
	Py_Initialize()

	list := PyList_New(0)
	defer list.DecRef()

	repr := list.Repr()

	assert.Equal(t, "[]", PyUnicode_AsUTF8(repr))
}

func TestStr(t *testing.T) {
	Py_Initialize()

	list := PyList_New(0)
	defer list.DecRef()

	str := list.Str()

	assert.Equal(t, "[]", PyUnicode_AsUTF8(str))
}

func TestASCII(t *testing.T) {
	Py_Initialize()

	list := PyList_New(0)
	defer list.DecRef()

	ascii := list.ASCII()

	assert.Equal(t, "[]", PyUnicode_AsUTF8(ascii))
}

func TestCallable(t *testing.T) {
	Py_Initialize()

	builtins := PyEval_GetBuiltins()
	assert.True(t, PyDict_Check(builtins))

	len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(len))

	emptyList := PyList_New(0)
	assert.True(t, PyList_Check(emptyList))

	args := PyTuple_New(1)
	defer args.DecRef()
	assert.True(t, PyTuple_Check(args))

	PyTuple_SetItem(args, 0, emptyList)

	length := len.Call(args, nil)
	assert.True(t, PyLong_Check(length))
	assert.Equal(t, 0, PyLong_AsLong(length))
	length.DecRef()

	length = len.CallObject(args)
	assert.True(t, PyLong_Check(length))
	assert.Equal(t, 0, PyLong_AsLong(length))
	length.DecRef()

	length = len.CallFunctionObjArgs(emptyList)
	assert.True(t, PyLong_Check(length))
	assert.Equal(t, 0, PyLong_AsLong(length))
	length.DecRef()

}

func TestCallMethod(t *testing.T) {
	Py_Initialize()

	s := PyUnicode_FromString("hello world")
	assert.True(t, PyUnicode_Check(s))
	defer s.DecRef()

	sep := PyUnicode_FromString(" ")
	assert.True(t, PyUnicode_Check(sep))
	defer sep.DecRef()

	split := PyUnicode_FromString("split")
	assert.True(t, PyUnicode_Check(split))
	defer split.DecRef()

	words := s.CallMethodObjArgs(split, sep)
	assert.True(t, PyList_Check(words))
	defer words.DecRef()
	assert.Equal(t, 2, PyList_Size(words))

	hello := PyList_GetItem(words, 0)
	assert.True(t, PyUnicode_Check(hello))
	world := PyList_GetItem(words, 1)
	assert.True(t, PyUnicode_Check(world))

	assert.Equal(t, "hello", PyUnicode_AsUTF8(hello))
	assert.Equal(t, "world", PyUnicode_AsUTF8(world))

	words.DecRef()

	words = s.CallMethodArgs("split", sep)
	assert.True(t, PyList_Check(words))
	defer words.DecRef()
	assert.Equal(t, 2, PyList_Size(words))

	hello = PyList_GetItem(words, 0)
	assert.True(t, PyUnicode_Check(hello))
	world = PyList_GetItem(words, 1)
	assert.True(t, PyUnicode_Check(world))

	assert.Equal(t, "hello", PyUnicode_AsUTF8(hello))
	assert.Equal(t, "world", PyUnicode_AsUTF8(world))

	words.DecRef()

}

func TestIsTrue(t *testing.T) {
	Py_Initialize()

	b := Py_True.IsTrue() != 0
	assert.True(t, b)

	b = Py_False.IsTrue() != 0
	assert.False(t, b)
}

func TestNot(t *testing.T) {
	Py_Initialize()

	b := Py_True.Not() != 0
	assert.False(t, b)

	b = Py_False.Not() != 0
	assert.True(t, b)
}

func TestLength(t *testing.T) {
	Py_Initialize()
	length := 6

	list := PyList_New(length)
	defer list.DecRef()
	listLength := list.Length()

	assert.Equal(t, length, listLength)

}

func TestLengthHint(t *testing.T) {
	Py_Initialize()
	length := 6

	list := PyList_New(length)
	defer list.DecRef()
	listLength := list.LengthHint(0)

	assert.Equal(t, length, listLength)

}

func TestObjectItem(t *testing.T) {
	Py_Initialize()

	key := PyUnicode_FromString("key")
	defer key.DecRef()
	value := PyUnicode_FromString("value")
	defer value.DecRef()

	dict := PyDict_New()
	err := dict.SetItem(key, value)
	assert.Zero(t, err)

	dictValue := dict.GetItem(key)
	assert.Equal(t, value, dictValue)

	err = dict.DelItem(key)
	assert.Zero(t, err)

}

func TestDir(t *testing.T) {
	Py_Initialize()

	list := PyList_New(0)
	defer list.DecRef()

	dir := list.Dir()
	defer dir.DecRef()

	repr := dir.Repr()
	defer repr.DecRef()

	assert.Equal(t, "['__add__', '__class__', '__contains__', '__delattr__', '__delitem__', '__dir__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__gt__', '__hash__', '__iadd__', '__imul__', '__init__', '__init_subclass__', '__iter__', '__le__', '__len__', '__lt__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__reversed__', '__rmul__', '__setattr__', '__setitem__', '__sizeof__', '__str__', '__subclasshook__', 'append', 'clear', 'copy', 'count', 'extend', 'index', 'insert', 'pop', 'remove', 'reverse', 'sort']", PyUnicode_AsUTF8(repr))

}

func TestReprEnterLeave(t *testing.T) {
	Py_Initialize()

	s := PyUnicode_FromString("hello world")
	defer s.DecRef()

	assert.Zero(t, s.ReprEnter())

	assert.True(t, s.ReprEnter() > 0)

	s.ReprLeave()
	s.ReprLeave()
}

func TestIsSubclass(t *testing.T) {
	Py_Initialize()

	assert.Equal(t, 1, PyExc_Warning.IsSubclass(PyExc_Exception))
	assert.Equal(t, 0, Bool.IsSubclass(Float))
}

func TestHash(t *testing.T) {
	Py_Initialize()

	s := PyUnicode_FromString("test string")
	defer s.DecRef()

	assert.NotEqual(t, -1, s.Hash())
}

func TestObjectType(t *testing.T) {
	Py_Initialize()

	i := PyLong_FromGoInt(23543)
	defer i.DecRef()

	assert.Equal(t, Long, i.Type())
}

func TestHashNotImplemented(t *testing.T) {
	Py_Initialize()

	s := PyUnicode_FromString("test string")
	defer s.DecRef()

	assert.Equal(t, -1, s.HashNotImplemented())

	assert.True(t, PyErr_ExceptionMatches(PyExc_TypeError))

	PyErr_Clear()
}

func TestObjectIter(t *testing.T) {
	Py_Initialize()

	i := PyLong_FromGoInt(23)
	defer i.DecRef()

	assert.Nil(t, i.GetIter())

	assert.True(t, PyErr_ExceptionMatches(PyExc_TypeError))
	PyErr_Clear()

	list := PyList_New(23)
	defer list.DecRef()

	iter := list.GetIter()
	assert.NotNil(t, iter)
	defer iter.DecRef()
}

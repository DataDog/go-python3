package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttrString(t *testing.T) {
	Py_Initialize()

	s := PyUnicode_FromString("test")

	assert.True(t, s.HasAttrString("split"))
	split := s.GetAttrString("split")
	assert.NotNil(t, split)
}

func TestAttr(t *testing.T) {
	Py_Initialize()

	s := PyUnicode_FromString("test")
	name := PyUnicode_FromString("split")

	assert.True(t, s.HasAttr(name))
	split := s.GetAttr(name)
	assert.NotNil(t, split)
}

func TestRichCompareBool(t *testing.T) {
	Py_Initialize()

	s1 := PyUnicode_FromString("test1")
	s2 := PyUnicode_FromString("test2")

	b, _ := s1.RichCompareBool(s2, Py_EQ)
	assert.False(t, b)

	b, _ = s1.RichCompareBool(s1, Py_EQ)
	assert.True(t, b)

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

	length, _ = len.CallFunctionObjArgs(emptyList)
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

	words, _ := s.CallMethodObjArgs(split, sep)
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

	words, _ = s.CallMethodArgs("split", sep)
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

	b, _ := Py_True.IsTrue()
	assert.True(t, b)

	b, _ = Py_False.IsTrue()
	assert.False(t, b)
}

func TestNot(t *testing.T) {
	Py_Initialize()

	b, _ := Py_True.Not()
	assert.False(t, b)

	b, _ = Py_False.Not()
	assert.True(t, b)
}

func TestLength(t *testing.T) {
	Py_Initialize()
	length := 6

	list := PyList_New(length)
	defer list.DecRef()
	listLength, _ := list.Length()

	assert.Equal(t, length, listLength)

}

func TestLengthHint(t *testing.T) {
	Py_Initialize()
	length := 6

	list := PyList_New(length)
	defer list.DecRef()
	listLength, _ := list.LengthHint(0)

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
	assert.Nil(t, err, err)

	dictValue := dict.GetItem(key)
	assert.Equal(t, value, dictValue)

	err = dict.DelItem(key)
	assert.Nil(t, err, err)

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

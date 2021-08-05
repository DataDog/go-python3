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

func TestDict(t *testing.T) {
	Py_Initialize()

	dict := PyDict_New()
	assert.True(t, PyDict_Check(dict))
	assert.True(t, PyDict_CheckExact(dict))
	defer dict.DecRef()

	proxy := PyDictProxy_New(dict)
	assert.NotNil(t, proxy)
	proxy.DecRef()

	key1 := "key1"
	value1 := PyUnicode_FromString("value1")
	assert.NotNil(t, value1)
	defer value1.DecRef()

	key2 := PyUnicode_FromString("key2")
	assert.NotNil(t, key2)
	defer key2.DecRef()
	value2 := PyUnicode_FromString("value2")
	assert.NotNil(t, value2)
	defer value2.DecRef()

	key3 := PyUnicode_FromString("key3")
	assert.NotNil(t, key3)
	defer key3.DecRef()
	value3 := PyUnicode_FromString("value3")
	assert.NotNil(t, value3)
	defer value3.DecRef()

	err := PyDict_SetItem(dict, key2, value2)
	assert.Zero(t, err)

	err = PyDict_SetItemString(dict, key1, value1)
	assert.Zero(t, err)

	assert.Equal(t, value2, PyDict_GetItem(dict, key2))
	assert.Equal(t, value2, PyDict_SetDefault(dict, key2, Py_None))
	assert.Equal(t, value1, PyDict_GetItemString(dict, key1))
	assert.Nil(t, PyDict_GetItemWithError(dict, key3))
	b := PyDict_Contains(dict, key2) != 0
	assert.True(t, b)

	assert.Equal(t, 2, PyDict_Size(dict))

	keys := PyDict_Keys(dict)
	assert.True(t, PyList_Check(keys))
	keys.DecRef()

	values := PyDict_Values(dict)
	assert.True(t, PyList_Check(values))
	values.DecRef()

	items := PyDict_Items(dict)
	assert.True(t, PyList_Check(items))
	items.DecRef()

	err = PyDict_SetItem(dict, key3, value3)
	assert.Zero(t, err)

	newDict := PyDict_Copy(dict)
	assert.Equal(t, 3, PyDict_Size(newDict))
	defer newDict.DecRef()

	err = PyDict_DelItem(dict, key2)
	assert.Zero(t, err)

	err = PyDict_DelItemString(dict, key1)
	assert.Zero(t, err)

	assert.Equal(t, 1, PyDict_Size(dict))

	PyDict_Clear(dict)
	assert.Equal(t, 0, PyDict_Size(dict))

	dict.DecRef()

	PyDict_ClearFreeList()

}

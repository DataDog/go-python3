/*
Unless explicitly stated otherwise all files in this repository are licensed
under the $license_for_repo License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package main

import (
	"fmt"
	"os"

	"github.com/DataDog/go-python3"
)

func main() {
	python3.Py_Initialize()

	if !python3.Py_IsInitialized() {
		fmt.Println("Error initializing the python interpreter")
		os.Exit(1)
	}

	err := printList()
	if err != nil {
		fmt.Printf("failed to print the python list: %s\n", err)
	}

	python3.Py_Finalize()
}

func printList() error {
	list := python3.PyList_New(5)

	if exc := python3.PyErr_Occurred(); list == nil && exc != nil {
		return fmt.Errorf("Fail to create python list object")
	}
	defer list.DecRef()

	repr, err := pythonRepr(list)
	if err != nil {
		return fmt.Errorf("fail to get representation of object list")
	}
	fmt.Printf("python list: %s\n", repr)

	return nil
}

func pythonRepr(o *python3.PyObject) (string, error) {
	if o == nil {
		return "", fmt.Errorf("object is nil")
	}

	s := o.Repr()
	if s == nil {
		python3.PyErr_Clear()
		return "", fmt.Errorf("failed to call Repr object method")
	}
	defer s.DecRef()

	return python3.PyUnicode_AsUTF8(s), nil
}

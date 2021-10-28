/*
Unless explicitly stated otherwise all files in this repository are licensed
under the MIT License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package main

import (
	"fmt"
	"os"

	python3 "github.com/go-python/cpy3"
)

func main() {
	i, err := python3.Py_Main(os.Args)
	if err != nil {
		fmt.Printf("error launching the python interpreter: %s\n", err)
		os.Exit(1)
	}
	if i == 1 {
		fmt.Println("The interpreter exited due to an exception")
		os.Exit(1)
	}
	if i == 2 {
		fmt.Println("The parameter list does not represent a valid Python command line")
		os.Exit(1)
	}
}

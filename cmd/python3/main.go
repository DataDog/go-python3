package main

import (
	"fmt"
	"os"

	python3 "github.com/DataDog/go-python3"
)

func main() {
	err := python3.Py_Main(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	output     = flag.String("o", "variadic.c", "output file")
	caseNumber = flag.Int("n", 20, "number of case in the switch statement")
)

func main() {
	flag.Parse()

	out, err := os.Create(*output)
	if err != nil {
		fmt.Printf("Error opening file %s: %s", *output, err)
		os.Exit(1)
	}
	defer out.Close()

	_, err = out.WriteString("#include \"Python.h\"\n\n")
	if err != nil {
		fmt.Printf("Error writing to file %s: %s", *output, err)
		os.Exit(1)
	}
	out.WriteString(renderTemplate(*caseNumber, "PyObject_CallFunctionObjArgs", "callable"))
	out.WriteString(renderTemplate(*caseNumber, "PyObject_CallMethodObjArgs", "obj", "name"))

}

func renderTemplate(n int, functionName string, pyArgs ...string) string {
	template :=
		`PyObject* _%s(%sint argc, PyObject **argv) {

    PyObject *result = NULL;
    switch (argc) {
%s    }
    return result;
}
`
	args := ""
	for _, arg := range pyArgs {
		args += fmt.Sprintf("PyObject *%s, ", arg)
	}
	switchTemplate := ""
	for i := 0; i < n+1; i++ {
		switchTemplate += fmt.Sprintf("        case %d:\n", i)
		args := ""
		for _, arg := range pyArgs {
			args += fmt.Sprintf("%s, ", arg)
		}
		for j := 0; j < i; j++ {
			args += fmt.Sprintf("argv[%d], ", j)
		}
		args += "NULL"
		switchTemplate += fmt.Sprintf("            result = %s(%s);\n", functionName, args)
		switchTemplate += fmt.Sprintf("            break;\n")
	}
	return fmt.Sprintf(template, functionName, args, switchTemplate)
}

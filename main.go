package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type ClassFile struct {
	Namespaces []string
	ClassName  string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid argument")
		os.Exit(1)
	}

	className := os.Args[1]

	if className == "" {
		fmt.Println("Invalid argument")
		os.Exit(2)
	}

	createClass(strings.Split(className, "/"))
}

var hppCode = `#pragma once

{{ range $v := .Namespaces }}
namespace {{$v}} {
{{ end }}
class {{.ClassName}} {
	public:
	{{.ClassName}}();
};
{{ range .Namespaces }}
}
{{ end }}
`

var cppCode = `#include "{{.ClassName}}.hpp"

{{ range $v := .Namespaces }}
namespace {{$v}} {
{{ end }}
{{.ClassName}}::{{.ClassName}} ()
{
}
{{ range .Namespaces }}
}
{{ end }}
`

func createClass(config []string) {
	cf := ClassFile{
		Namespaces: config[:len(config)-1],
		ClassName:  config[len(config)-1],
	}

	hppTmpl, err := template.New("hpp.tmpl").Parse(hppCode)
	if err != nil {
		panic(err)
	}

	cppTmpl, err := template.New("cpp.tmpl").Parse(cppCode)
	if err != nil {
		panic(err)
	}

	filePath := fmt.Sprintf("%s/%s", strings.Join(cf.Namespaces, "/"), cf.ClassName)
	hppName := fmt.Sprintf("%s/%s.hpp", filePath, cf.ClassName)
	cppName := fmt.Sprintf("%s/%s.cpp", filePath, cf.ClassName)

	err = os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	hppFile, err := os.Create(hppName)
	if err != nil {
		panic(err)
	}

	cppFile, err := os.Create(cppName)
	if err != nil {
		panic(err)
	}

	err = hppTmpl.Execute(hppFile, cf)
	if err != nil {
		panic(err)
	}

	err = cppTmpl.Execute(cppFile, cf)
	if err != nil {
		panic(err)
	}
}

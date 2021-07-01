package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	var (
		err error
		tpl *template.Template
		out string
		in  string
	)
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <template> [input]")
		os.Exit(1)
	}
	in = os.Args[len(os.Args)-1]
	tpl, err = template.New("").Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error parsing template:", err)
		os.Exit(1)
	}
	out, err = filepath.Abs(in)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		os.Exit(1)
	}
	out = strings.TrimSuffix(out, filepath.Ext(out))
	out = out + ".go"
	err = tpl.Execute(os.Stdout, map[string]string{
		"input":  in,
		"output": out,
	})
	if err != nil {
		fmt.Println("Error executing template:", err)
		os.Exit(1)
	}
}

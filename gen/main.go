package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Update to a newer version of Go by changing this constant.
const goTag = "go1.13.8"

type Status struct {
	Name   string
	Status string
}

func main() {
	// Download and parse the statuses from the Go source code file
	statuses, err := parseFileForStatuses()
	if err != nil {
		log.Fatalln("unable to download and parse file:", err)
	}

	// Generate and output the statuses for erkhttp
	file := generateFile(statuses)
	err = ioutil.WriteFile("http_statuses.go", []byte(file), 0655)
	if err != nil {
		log.Fatalln("unable to write generated file:", err)
	}
}

func parseFileForStatuses() ([]Status, error) {
	file, err := downloadStatusFile()
	if err != nil {
		return nil, err
	}

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "net/http/status.go", file, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	statuses := []Status{}

	log.Println("Parsing file. Most warnings can be ignored.")
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			log.Println("WARNING: decl not *ast.GenDecl, so skipping")
			continue
		}

		if genDecl.Tok.String() != "const" {
			log.Println("WARNING: tok not const, so skipping")
			continue
		}

		for _, spec := range genDecl.Specs {
			valSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				log.Println("WARNING: spec not *ast.ValueSpec, so skipping")
				continue
			}

			if len(valSpec.Names) != 1 {
				log.Println("WARNING: names group not len 1, so skipping")
				continue
			}
			name := valSpec.Names[0].String()

			if !strings.HasPrefix(name, "Status") {
				log.Printf("WARNING: name '%s' does not start with 'Status', so skipping\n", name)
				continue
			}

			if len(valSpec.Values) != 1 {
				log.Println("WARNING: values group not len 1, so skipping")
				continue
			}
			valueDetails := valSpec.Values[0]

			valueBasicLit, ok := valueDetails.(*ast.BasicLit)
			if !ok {
				log.Println("WARNING: value not basic literal, so skipping")
				continue
			}
			value := valueBasicLit.Value

			statuses = append(statuses, Status{
				Name:   name,
				Status: value,
			})
		}
	}

	return statuses, nil
}

func downloadStatusFile() (string, error) {
	res, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/golang/go/%s/src/net/http/status.go", goTag))
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func generateFile(statuses []Status) string {
	file := "// WARNING: THIS FILE IS GENERATED USING gen/main.go. Edit that script and regenerate this file.\n\n"
	file += "package erkhttp\n\n"

	for _, status := range statuses {
		file += fmt.Sprintf("// %s = %s\n", status.Name, status.Status)
		file += fmt.Sprintf("type %s struct{}\n\n", status.Name)
		file += fmt.Sprintf("func (%s) HTTPStatus() int { return %s }\n\n", status.Name, status.Status)
		file += fmt.Sprintf("var _ HTTPStatusable = %s{}\n\n", status.Name)
	}

	return file
}

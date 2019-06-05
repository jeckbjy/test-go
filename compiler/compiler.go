package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"log"
)

// https://getstream.io/blog/how-a-go-program-compiles-down-to-machine-code/
// https://github.com/golang/go/blob/master/src/cmd/compile/README.md
// go compile source: go/src/cmd/compile
func main() {
	src := `
		package main

		import "fmt"
		func main() {
		fmt.Println("Hello, world!")
		}
	`
	buildScanner(src)
	buildParser(src)
}

func buildScanner(src string) {
	var s scanner.Scanner

	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, []byte(src), nil, 0)

	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}
}

func buildParser(src string) {
	log.Printf("\n\n\n")

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(fset, file)
}
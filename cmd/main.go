package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("../README.md", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("error opening log file:", err)
		return
	}

	defer logFile.Close()

	fset := token.NewFileSet()
	d, err := parser.ParseDir(fset, "../", nil, parser.ParseComments)
	if err != nil {
		log.Println(err)
		return
	}

	for _, f := range d {
		ast.Inspect(f, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.FuncDecl:
				logFile.WriteString("## Function: " + x.Name.Name)
				logFile.WriteString("\n\n")

				logFile.WriteString("### Comment: \n\n")

				logFile.WriteString(x.Doc.Text())

				logFile.WriteString("\n")

				if x.Type.Params != nil {
					logFile.WriteString("### parameter: ")
					for _, v := range x.Type.Params.List {
						for _, name := range v.Names {
							logFile.WriteString(name.Name)
							logFile.WriteString(" ")
						}
						logFile.WriteString(fmt.Sprintf("(%s) ", v.Type))
					}
				}

				logFile.WriteString("\n\n")
				if x.Type.Results != nil {
					logFile.WriteString("### return: ")
					for _, v := range x.Type.Results.List {
						for _, name := range v.Names {
							logFile.WriteString(name.Name)
						}

						logFile.WriteString(fmt.Sprintf("(%s) ", v.Type))
					}
				}
				logFile.WriteString("\n\n")
				logFile.WriteString("\n\n")

			}
			return true
		})
	}

}

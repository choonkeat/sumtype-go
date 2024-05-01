package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"strings"
)

type ParsedField struct {
	Name string // name of the argument
	Type string // type of the argument
}

type ParsedGeneric struct {
	Name       string // name of the generic type
	Constraint string // type of the generic type
}

type ParsedData struct {
	Name     string          // name of the struct
	Fields   []ParsedField   // fields of the struct
	Generics []ParsedGeneric // generics of the struct
}

type ParsedFile struct {
	PackageName string
	Imports     []string
	Data        map[string][]ParsedData
}

func parseFile(flags Flags) (ParsedFile, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, flags.inputFile, nil, parser.ParseComments)
	if err != nil {
		return ParsedFile{}, err
	}

	src, err := ioutil.ReadFile(flags.inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var parsedFile ParsedFile
	parsedFile.PackageName = node.Name.Name // Extracting the package name
	parsedFile.Data = make(map[string][]ParsedData)

	for _, imp := range node.Imports {
		// Extracting the import paths
		path := strings.Trim(imp.Path.Value, `"`)
		if imp.Name != nil {
			parsedFile.Imports = append(parsedFile.Imports, imp.Name.Name+" "+path)
		} else {
			parsedFile.Imports = append(parsedFile.Imports, path)
		}
	}

	for _, f := range node.Decls {
		// Check if the declaration is a general declaration (var, const, type, or func)
		genDecl, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range genDecl.Specs {
			// Check if the spec is a type specification (type name Type = ...)
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok || !strings.HasSuffix(typeSpec.Name.Name, flags.structSuffix) {
				continue
			}
			parsedFile.Data[typeSpec.Name.Name] = []ParsedData{}

			var parsedGeneric []ParsedGeneric
			// Add this block to handle type parameters
			if typeSpec.TypeParams != nil {
				for _, param := range typeSpec.TypeParams.List {
					for _, name := range param.Names {
						start := fset.Position(param.Type.Pos()).Offset
						end := fset.Position(param.Type.End()).Offset
						parsedGeneric = append(parsedGeneric, ParsedGeneric{
							Name:       name.Name,
							Constraint: string(src[start:end]),
						})
					}
				}
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			for _, field := range structType.Fields.List {
				funcType, ok := field.Type.(*ast.FuncType)
				if !ok {
					continue
				}

				var fields []ParsedField
				if funcType.Params != nil {
					for _, param := range funcType.Params.List {
						// Handle different types of fields
						fieldType := exprToString(param.Type)
						for _, paramName := range param.Names {
							fields = append(fields, ParsedField{
								Name: paramName.Name,
								Type: fieldType,
							})
						}
					}
				}

				for _, fieldName := range field.Names {
					parsedFile.Data[typeSpec.Name.Name] = append(parsedFile.Data[typeSpec.Name.Name], ParsedData{
						Name:     fieldName.Name,
						Fields:   fields,
						Generics: parsedGeneric,
					})
				}
			}
		}
	}

	return parsedFile, nil
}

// exprToString converts an ast.Expr to a string representation of the type
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt)
	case *ast.IndexExpr:
		return exprToString(t.X) + "[" + exprToString(t.Index) + "]"
	// Add more cases as needed for other types
	default:
		return fmt.Sprintf("exprToString failed: %#v", expr)
	}
}

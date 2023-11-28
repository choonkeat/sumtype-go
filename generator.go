package main

import (
	"fmt"
	"strings"
	"unicode"
)

func writeGoCode(flags Flags, parsedFile ParsedFile, builder *strings.Builder) {
	builder.WriteString("// Generated code by github.com/choonkeat/sumtype-go\npackage ")
	builder.WriteString(parsedFile.PackageName)
	builder.WriteString("\n\n")
	builder.WriteString("import (\n")
	for _, imp := range parsedFile.Imports {
		builder.WriteString("\t\"")
		builder.WriteString(imp)
		builder.WriteString("\"\n")
	}
	builder.WriteString(")\n")

	for _, data := range parsedFile.Data {
		structName := unexported(data.Name)

		// Generate struct
		fmt.Fprintf(builder, "\n// %s\n", data.Name)
		fmt.Fprintf(builder, "type %s struct {\n", structName)
		for _, field := range data.Fields {
			fmt.Fprintf(builder, "\t%s %s\n", field.Name, field.Type)
		}
		fmt.Fprintf(builder, "}\n\n")

		// Generate method
		fmt.Fprintf(builder, "func (s %s) %s(scenarios %s) {\n", structName, flags.switchName, parsedFile.Name)
		if len(data.Fields) > 0 {
			fmt.Fprintf(builder, "\tscenarios.%s(s.%s)\n", data.Name, strings.Join(getFieldNames(data.Fields), ", s."))
		} else {
			fmt.Fprintf(builder, "\tscenarios.%s()\n", data.Name)
		}
		fmt.Fprintf(builder, "}\n\n")

		// Generate constructor function
		fmt.Fprintf(builder, "func %s(%s) %s {\n", data.Name, getParamList(data.Fields), strings.TrimSuffix(parsedFile.Name, flags.structSuffix))
		if len(data.Fields) > 0 {
			fmt.Fprintf(builder, "\treturn %s{%s}\n", structName, strings.Join(getFieldNames(data.Fields), ", "))
		} else {
			fmt.Fprintf(builder, "\treturn %s{}\n", structName)
		}
		fmt.Fprintf(builder, "}\n")
	}
}

// unexported returns an unexported (lowercase) version of the given string
func unexported(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

// getFieldNames returns a slice of field names from a slice of ParsedField
func getFieldNames(fields []ParsedField) []string {
	var names []string
	for _, field := range fields {
		names = append(names, field.Name)
	}
	return names
}

// getParamList returns a string representing the parameter list for a function
func getParamList(fields []ParsedField) string {
	var params []string
	for _, field := range fields {
		params = append(params, fmt.Sprintf("%s %s", field.Name, field.Type))
	}
	return strings.Join(params, ", ")
}

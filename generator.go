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
	if len(parsedFile.Imports) > 0 {
		builder.WriteString("import (\n")
		for _, imp := range parsedFile.Imports {
			builder.WriteString("\t\"")
			builder.WriteString(imp)
			builder.WriteString("\"\n")
		}
		builder.WriteString(")\n")
	}

	for _, data := range parsedFile.Data {
		structName := unexported(data.Name) + parsedFile.Name
		genericsDecl := buildGenericTypeDeclaration(data.Generics)
		paramList := buildParameterListDeclaration(data.Generics)

		// Generate struct
		fmt.Fprintf(builder, "\n// %s\n", data.Name)
		fmt.Fprintf(builder, "type %s%s struct {\n", structName, genericsDecl)
		for _, field := range data.Fields {
			fmt.Fprintf(builder, "\t%s %s\n", field.Name, field.Type)
		}
		fmt.Fprintf(builder, "}\n\n")

		// Generate method
		fmt.Fprintf(builder, "func (s %s%s) %s(scenarios %s%s) {\n", structName, paramList, flags.switchName, parsedFile.Name, paramList)
		if len(data.Fields) > 0 {
			fmt.Fprintf(builder,
				"\tscenarios.%s(s.%s)\n",
				data.Name,
				strings.Join(getFieldNames("", data.Fields), ", s."),
			)
		} else {
			fmt.Fprintf(builder, "\tscenarios.%s()\n", data.Name)
		}
		fmt.Fprintf(builder, "}\n\n")

		// Generate constructor function
		fmt.Fprintf(builder,
			"func %s%s(%s) %s%s {\n",
			data.Name,
			genericsDecl,
			getParamList("Arg", data.Fields),
			strings.TrimSuffix(parsedFile.Name, flags.structSuffix),
			paramList,
		)
		if len(data.Fields) > 0 {
			fmt.Fprintf(builder,
				"\treturn %s%s{%s}\n",
				structName,
				paramList,
				strings.Join(getFieldNames("Arg", data.Fields), ", "),
			)
		} else {
			fmt.Fprintf(builder, "\treturn %s%s{}\n", structName, paramList)
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
func getFieldNames(suffix string, fields []ParsedField) []string {
	var names []string
	for _, field := range fields {
		names = append(names, field.Name+suffix)
	}
	return names
}

// getParamList returns a string representing the parameter list for a function
func getParamList(suffix string, fields []ParsedField) string {
	var params []string
	for _, field := range fields {
		params = append(params, fmt.Sprintf("%s%s %s", field.Name, suffix, field.Type))
	}
	return strings.Join(params, ", ")
}

func buildGenericTypeDeclaration(data []ParsedGeneric) string {
	if len(data) == 0 {
		return ""
	}

	constraintMap := make(map[string][]string)
	var constraintsOrder []string

	for _, d := range data {
		if _, exists := constraintMap[d.Constraint]; !exists {
			constraintsOrder = append(constraintsOrder, d.Constraint)
		}
		constraintMap[d.Constraint] = append(constraintMap[d.Constraint], d.Name)
	}

	var result []string
	for _, constraint := range constraintsOrder {
		names := constraintMap[constraint]
		result = append(result, fmt.Sprintf("%s %s", strings.Join(names, ", "), constraint))
	}

	return "[" + strings.Join(result, ", ") + "]"
}

func buildParameterListDeclaration(data []ParsedGeneric) string {
	if len(data) == 0 {
		return ""
	}

	var names []string
	for _, d := range data {
		names = append(names, d.Name)
	}
	return "[" + strings.Join(names, ", ") + "]"
}

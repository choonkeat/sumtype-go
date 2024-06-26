package main

import (
	"fmt"
	"sort"
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

	// if output file is committed to git, sorting will prevent unnecessary diffs
	for _, row := range mapToSortedSlice(parsedFile.Data) {
		typeName := row.key
		dataList := row.value
		for _, data := range dataList {
			structName := unexported(data.Name) + typeName
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
			fmt.Fprintf(builder, "func (s %s%s) %s(variants %s%s) {\n", structName, paramList, flags.patternMatchFunction, typeName, paramList)
			if len(data.Fields) > 0 {
				fmt.Fprintf(builder,
					"\tvariants.%s(s.%s)\n",
					data.Name,
					strings.Join(getFieldNames("", data.Fields), ", s."),
				)
			} else {
				fmt.Fprintf(builder, "\tvariants.%s()\n", data.Name)
			}
			fmt.Fprintf(builder, "}\n\n")

			// Generate constructor function
			fmt.Fprintf(builder,
				"func %s%s(%s) %s%s {\n",
				data.Name,
				genericsDecl,
				getParamList("Arg", data.Fields),
				strings.TrimSuffix(typeName, flags.structSuffix),
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

// Define the tuple struct as a generic type
type tuple[K, V any] struct {
	key   K
	value V
}

// Function to convert a map[K]V to a sorted []tuple[K, V]
// K must be both comparable and ordered
func mapToSortedSlice[V any](m map[string]V) []tuple[string, V] {
	result := make([]tuple[string, V], 0, len(m))
	for k, v := range m {
		result = append(result, tuple[string, V]{key: k, value: v})
	}

	// Sort the slice of tuples by the key
	sort.Slice(result, func(i, j int) bool {
		return result[i].key < result[j].key
	})

	return result
}

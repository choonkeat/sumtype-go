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
		mainTypeName := strings.TrimSuffix(typeName, flags.structSuffix)

		generics := []ParsedGeneric{}
		for _, data := range dataList {
			generics = data.Generics // keep it for later
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
			fmt.Fprintf(builder, "\n// %s is a constructor function for %s; see %s for all constructor functions of %s\n",
				data.Name,
				mainTypeName,
				typeName,
				mainTypeName,
			)
			fmt.Fprintf(builder,
				"func %s%s(%s) %s%s {\n",
				data.Name,
				genericsDecl,
				getParamList("Arg", data.Fields),
				mainTypeName,
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

		// Generate generic variant
		fmt.Fprintf(builder, "\n// %sMap is parameter type of %sMap function,\n// like %s is parameter type of %s.Match method,\n// but with methods that returns a value of generic type\n",
			typeName,
			mainTypeName,
			typeName,
			unexported(mainTypeName),
		)
		genericT := newGenericType(generics, "any")
		combinedGenerics := append(generics, genericT)
		fmt.Fprintf(builder, "type %sMap%s struct{\n", typeName, buildGenericTypeDeclaration(combinedGenerics))
		for _, data := range dataList {
			fmt.Fprintf(builder, "\t%s func(%s) %s // when %s value pattern matches to %s, return different value\n",
				data.Name,
				getParamList("Arg", data.Fields),
				genericT.Name,
				mainTypeName,
				data.Name,
			)
		}
		fmt.Fprintf(builder, "}\n\n")

		// Generate Mapping function
		fmt.Fprintf(builder, "\n// %sMap is like %s.Match method except it returns a value of generic type\n// thus can transform a %s value into anything else\n",
			mainTypeName,
			unexported(mainTypeName),
			mainTypeName,
		)
		fmt.Fprintf(builder, "func %sMap%s(value %s%s, variants %sMap%s) %s {\n",
			mainTypeName,
			buildGenericTypeDeclaration(combinedGenerics),
			mainTypeName,
			buildParameterListDeclaration(generics),
			typeName,
			buildParameterListDeclaration(combinedGenerics),
			genericT.Name,
		)
		fmt.Fprintf(builder, "\tvar result %s\n", genericT.Name)
		fmt.Fprintf(builder, "\tvalue.Match(%s%s{\n", typeName, buildParameterListDeclaration(generics))
		for _, data := range dataList {
			fmt.Fprintf(builder, "\t\t%s: func(%s) {\n", data.Name, getParamList("Arg", data.Fields))
			fmt.Fprintf(builder, "\t\t\tresult = variants.%s(%s)\n", data.Name, strings.Join(getFieldNames("Arg", data.Fields), ", "))
			fmt.Fprintf(builder, "\t\t},\n")
		}
		fmt.Fprintf(builder, "\t})\n")
		fmt.Fprintf(builder, "\treturn result\n")
		fmt.Fprintf(builder, "}\n\n")
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

func newGenericType(generics []ParsedGeneric, constraint string) ParsedGeneric {
	// find a letter that is not already used
	used := make(map[rune]bool)
	for _, g := range generics {
		used[rune(g.Name[0])] = true
	}

	var name string
	for i := 'A'; i <= 'Z'; i++ {
		if !used[i] {
			name = string(i)
			break
		}
	}

	return ParsedGeneric{Name: name, Constraint: constraint}
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

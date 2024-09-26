package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"unicode"
)

func writeGoCode(flags Flags, parsedFile ParsedFile, builder *strings.Builder) {
	builder.WriteString("// Generated code by github.com/choonkeat/sumtype-go")
	if flags.actualVersion != "" {
		builder.WriteString("@")
		builder.WriteString(flags.actualVersion)
	}
	builder.WriteString("\npackage ")
	builder.WriteString(parsedFile.PackageName)
	builder.WriteString("\n\n")
	parsedFile.Imports = append(parsedFile.Imports, "encoding/json", "fmt")
	slices.Sort(parsedFile.Imports)
	parsedFile.Imports = slices.Compact(parsedFile.Imports)

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
		var genericsDecl, paramList string
		if len(dataList) > 0 {
			data := dataList[0]
			generics = data.Generics // keep it for later
			genericsDecl = buildGenericTypeDeclaration(data.Generics)
			paramList = buildParameterListDeclaration(data.Generics)
		}

		for _, data := range dataList {
			structName := unexported(data.Name) + typeName

			// Generate struct
			fmt.Fprintf(builder, "\n// %s\n", data.Name)
			fmt.Fprintf(builder, "type %s%s struct {\n", structName, genericsDecl)
			for _, field := range data.Fields {
				fmt.Fprintf(builder, "\t%s %s\n", exported(field.Name), field.Type)
			}
			fmt.Fprintf(builder, "}\n\n")

			// Generate method
			fmt.Fprintf(builder, "func (%sInstance %s%s) %s(%sVariants %s%s) {\n",
				unexported(data.Name),
				structName, paramList, flags.patternMatchFunction,
				unexported(data.Name), typeName, paramList)
			if len(data.Fields) > 0 {
				fmt.Fprintf(builder,
					"\t%sVariants.%s(%sInstance.%s)\n",
					unexported(data.Name),
					data.Name,
					unexported(data.Name),
					strings.Join(getFieldNames("", data.Fields), ", "+unexported(data.Name)+"Instance."),
				)
			} else {
				fmt.Fprintf(builder, "\t%sVariants.%s()\n", unexported(data.Name), data.Name)
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
					"\treturn %s%s{%s%s{%s}}\n",
					mainTypeName,
					paramList,
					structName,
					paramList,
					strings.Join(getFieldNames("Arg", data.Fields), ", "),
				)
			} else {
				fmt.Fprintf(builder, "\treturn %s%s{%s%s{}}\n", mainTypeName, paramList, structName, paramList)
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
		fmt.Fprintf(builder, "func %sMap%s(%sValue %s%s, %sVariants %sMap%s) %s {\n",
			mainTypeName,
			buildGenericTypeDeclaration(combinedGenerics),
			unexported(mainTypeName),
			mainTypeName,
			buildParameterListDeclaration(generics),
			unexported(mainTypeName),
			typeName,
			buildParameterListDeclaration(combinedGenerics),
			genericT.Name,
		)
		fmt.Fprintf(builder, "\tvar %sTemp %s\n", unexported(mainTypeName), genericT.Name)
		fmt.Fprintf(builder, "\t%sValue.Match(%s%s{\n", unexported(mainTypeName), typeName, buildParameterListDeclaration(generics))
		for _, data := range dataList {
			fmt.Fprintf(builder, "\t\t%s: func(%s) {\n", data.Name, getParamList("Arg", data.Fields))
			fmt.Fprintf(builder, "\t\t\t%sTemp = %sVariants.%s(%s)\n", unexported(mainTypeName), unexported(mainTypeName), data.Name, strings.Join(getFieldNames("Arg", data.Fields), ", "))
			fmt.Fprintf(builder, "\t\t},\n")
		}
		fmt.Fprintf(builder, "\t})\n")
		fmt.Fprintf(builder, "\treturn %sTemp\n", unexported(mainTypeName))
		fmt.Fprintf(builder, "}\n\n")

		variantNames := make([]string, len(dataList))
		for i, data := range dataList {
			variantNames[i] = data.Name
		}

		// Generate struct
		fmt.Fprintf(builder, "\n// %s = %s\n", mainTypeName, strings.Join(variantNames, " | "))
		fmt.Fprintf(builder, "type %s%s struct {\n", mainTypeName, genericsDecl)
		fmt.Fprintf(builder, "\t%s %s%s\n", unexported(mainTypeName), unexported(mainTypeName), paramList)
		fmt.Fprintf(builder, "}\n")

		// Generate interface
		fmt.Fprintf(builder, "\n// %s is the interface for %s\n", unexported(mainTypeName), typeName)
		fmt.Fprintf(builder, "type %s%s interface {\n", unexported(mainTypeName), genericsDecl)
		fmt.Fprintf(builder, "\tMatch(variants %s%s)\n", typeName, paramList)
		fmt.Fprintf(builder, "}\n")

		// Generate Match method of JSON struct
		fmt.Fprintf(builder, "func (%sInstance %s%s) Match(%sVariants %s%s) {\n", unexported(mainTypeName), mainTypeName, paramList, unexported(mainTypeName), typeName, paramList)
		fmt.Fprintf(builder, "\t%sInstance.%s.Match(%sVariants)\n", unexported(mainTypeName), unexported(mainTypeName), unexported(mainTypeName))
		fmt.Fprintf(builder, "}\n")

		// Generate MarshalJSON method for {mainTypeName}JSON
		fmt.Fprintf(builder, "func (%sInstance %s%s) MarshalJSON() (%sData []byte, %sErr error) {\n", unexported(mainTypeName), mainTypeName, paramList, unexported(mainTypeName), unexported(mainTypeName))
		fmt.Fprintf(builder, "\t%sInstance.%s.Match(%s%s{\n", unexported(mainTypeName), unexported(mainTypeName), typeName, paramList)
		for _, data := range dataList {
			structName := unexported(data.Name) + typeName
			fmt.Fprintf(builder, "\t\t%s: func(%s) {\n", data.Name, getParamList("Arg", data.Fields))
			fmt.Fprintf(builder, "\t\t\t%sData, %sErr = json.Marshal([]any{\n", unexported(mainTypeName), unexported(mainTypeName))
			fmt.Fprintf(builder, "\t\t\t\t%#v,\n", data.Name)
			fmt.Fprintf(builder, "\t\t\t\t%s%s{\n", structName, paramList)
			for _, field := range data.Fields {
				fmt.Fprintf(builder, "\t\t\t\t\t%s: %sArg,\n", exported(field.Name), field.Name)
			}
			fmt.Fprintf(builder, "\t\t\t}})\n")
			fmt.Fprintf(builder, "\t\t},\n")
		}
		fmt.Fprintf(builder, "\t})\n")
		fmt.Fprintf(builder, "\treturn %sData, %sErr\n", unexported(mainTypeName), unexported(mainTypeName))
		fmt.Fprintf(builder, "}\n")

		// Generate UnmarshalJSON method for {mainTypeName}JSON
		fmt.Fprintf(builder, "func (%sInstance *%s%s) UnmarshalJSON(%sData []byte) error {\n", unexported(mainTypeName), mainTypeName, paramList, unexported(mainTypeName))
		fmt.Fprintf(builder, "\t// The expected format is [\"TypeName\", { ... data... }]\n")
		fmt.Fprintf(builder, "\tvar %sRaw []json.RawMessage\n", unexported(mainTypeName))
		fmt.Fprintf(builder, "\tif err := json.Unmarshal(%sData, &%sRaw); err != nil {\n", unexported(mainTypeName), unexported(mainTypeName))
		fmt.Fprintf(builder, "\treturn fmt.Errorf(\"expected an array with type and data, got error: %%w\", err)\n")
		fmt.Fprintf(builder, "\t}\n")

		fmt.Fprintf(builder, "\tif len(%sRaw) != 2 {\n", unexported(mainTypeName))
		fmt.Fprintf(builder, "\treturn fmt.Errorf(\"expected array of two elements [type, data], got %%d elements\", len(%sRaw))\n", unexported(mainTypeName))
		fmt.Fprintf(builder, "\t}\n")

		fmt.Fprintf(builder, "\t// Unmarshal the first element to get the type\n")
		fmt.Fprintf(builder, "\tvar %sVariantName string\n", unexported(mainTypeName))
		fmt.Fprintf(builder, "\tif err := json.Unmarshal(%sRaw[0], &%sVariantName); err != nil {\n", unexported(mainTypeName), unexported(mainTypeName))
		fmt.Fprintf(builder, "\treturn fmt.Errorf(\"failed to unmarshal type name: %%w\", err)\n")
		fmt.Fprintf(builder, "\t}\n")

		fmt.Fprintf(builder, "\tswitch %sVariantName {\n", unexported(mainTypeName))
		for _, data := range dataList {
			structName := unexported(data.Name) + typeName
			fmt.Fprintf(builder, "\tcase %#v:\n", data.Name)
			fmt.Fprintf(builder, "\tvar %sTemp %s%s\n", unexported(mainTypeName), structName, paramList)
			fmt.Fprintf(builder, "\tif err := json.Unmarshal(%sRaw[1], &%sTemp); err != nil {\n", unexported(mainTypeName), unexported(mainTypeName))
			fmt.Fprintf(builder, "\treturn fmt.Errorf(\"failed to unmarshal data: %%w\", err)\n")
			fmt.Fprintf(builder, "\t}\n")
			fmt.Fprintf(builder, "\t%sInstance.%s = %s%s{\n", unexported(mainTypeName), unexported(mainTypeName), structName, paramList)
			for _, field := range data.Fields {
				fmt.Fprintf(builder, "\t\t%s: %sTemp.%s,\n", exported(field.Name), unexported(mainTypeName), exported(field.Name))
			}
			fmt.Fprintf(builder, "\t}\n")
		}
		fmt.Fprintf(builder, "\tdefault:\n")
		fmt.Fprintf(builder, "\treturn fmt.Errorf(\"unknown type %%q\", %sVariantName)\n", unexported(mainTypeName))
		fmt.Fprintf(builder, "\t}\n")
		fmt.Fprintf(builder, "\treturn nil\n")
		fmt.Fprintf(builder, "\t}\n")
	}
}

// exported returns an exported (lowercase) version of the given string
func exported(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
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
		if suffix != "" {
			names = append(names, field.Name+suffix)
		} else {
			names = append(names, exported(field.Name))
		}
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

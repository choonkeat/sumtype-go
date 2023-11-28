package main

import (
	"flag"
	"fmt"
	"go/format"
	"os"
	"strings"
)

func main() {
	// Command-line flags
	var inputFile string
	flag.StringVar(&inputFile, "input", "", "Input file name")
	flag.Parse()
	if inputFile == "" {
		flag.Usage()
		return
	}

	// Read and parse the input file
	parsedFile, err := parseFile(inputFile)
	if err != nil {
		panic(err)
	}

	var builder strings.Builder
	writeGoCode(parsedFile, &builder)
	formattedCode, err := format.Source([]byte(builder.String()))
	if err != nil {
		fmt.Println("Error formatting source:", err)
		return
	}

	// Open the output file
	outputFile := strings.Replace(inputFile, ".go", ".boilerplate.go", 1)
	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the generated code to the output file
	_, err = file.Write(formattedCode)
	if err != nil {
		panic(err)
	}
}

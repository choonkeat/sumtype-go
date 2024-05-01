package main

import (
	"flag"
	"fmt"
	"go/format"
	"os"
	"strings"
)

type Flags struct {
	inputFile            string
	structSuffix         string
	patternMatchFunction string
}

func main() {
	// Command-line flags
	var flags Flags
	flag.StringVar(&flags.inputFile, "input", "", "Input file name")
	flag.StringVar(&flags.structSuffix, "suffix", "Variants", "Suffix of the struct defining variants")
	flag.StringVar(&flags.patternMatchFunction, "pattern-match", "Match", "Name of the pattern match method")
	flag.Parse()
	if flags.inputFile == "" {
		flag.Usage()
		return
	}

	// Read and parse the input file
	parsedFile, err := parseFile(flags)
	if err != nil {
		panic(err)
	}

	var builder strings.Builder
	writeGoCode(flags, parsedFile, &builder)
	formattedCode, err := format.Source([]byte(builder.String()))
	if err != nil {
		fmt.Println("Error formatting source:", err)
		// return
		formattedCode = []byte(builder.String())
	}

	// Open the output file
	outputFile := strings.Replace(flags.inputFile, ".go", ".boilerplate.go", 1)
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

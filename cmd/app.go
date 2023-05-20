package main

import (
	"YadroTestWork/internal/utils/functions/reader"
	"YadroTestWork/internal/utils/functions/sorting"
	"YadroTestWork/internal/utils/functions/writer"
	"YadroTestWork/internal/utils/structures"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		panic("Program must contains 2 arguments: <input-file> <output-file>.")
	}

	dataFile, jsonErr := os.ReadFile("cmd/resources/sort.json")
	if jsonErr != nil {
		panic(jsonErr)
	}

	array, err := reader.Read(arguments[0])
	if err != nil {
		panic(err)
	}

	result := sorting.Sort(array, structures.Build(dataFile).Parse())

	writeErr := writer.Write(arguments[1], result)
	if writeErr != nil {
		panic(writeErr)
	}
}

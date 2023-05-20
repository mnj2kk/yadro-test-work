package main

import (
	"YadroTestWork/internal/utils/functions/reader"
	"YadroTestWork/internal/utils/functions/sort"
	"YadroTestWork/internal/utils/functions/writer"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		panic("Program must contains 2 arguments: <input-file> <output-file>.")
	}

	array, err := reader.Read(arguments[0])
	if err != nil {
		panic(err)
	}

	result := sort.Sort(array)

	writeErr := writer.Write(arguments[1], result)
	if writeErr != nil {
		panic(writeErr)
	}
}

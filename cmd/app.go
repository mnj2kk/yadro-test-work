package main

import (
	"YadroTestWork/internal/utils/functions/handler"
	"YadroTestWork/internal/utils/functions/reader"
	"YadroTestWork/internal/utils/functions/sorting"
	"YadroTestWork/internal/utils/functions/writer"
	"YadroTestWork/internal/utils/structures"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		panic("Program must contains 2 arguments: <input-file> <output-file>.")
	}
	handler.Handle(godotenv.Load())

	dataFile, jsonErr := os.ReadFile(fmt.Sprintf("cmd/config/%s.json", os.Getenv("CONFIG")))
	handler.Handle(jsonErr)

	array := reader.Read(arguments[0])
	result := sorting.Sort(array, structures.Parse(structures.Build(dataFile)))

	writer.Write(arguments[1], result)
}

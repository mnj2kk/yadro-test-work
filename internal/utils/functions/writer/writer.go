package writer

import (
	"YadroTestWork/internal/utils/functions/handler"
	"YadroTestWork/internal/utils/structures"
	"bufio"
	"fmt"
	"os"
)

func Write(path string, output structures.Result) (err error) {
	file, createErr := os.Create(path)
	handler.Handle(createErr)
	defer func(file *os.File) {
		handler.Handle(file.Close())
	}(file)

	writer := bufio.NewWriter(file)
	for _, kv := range output {
		_, outErr := fmt.Fprintf(writer, "%s: %d\n", kv.First, kv.Second)
		handler.Handle(outErr)
	}

	flushErr := writer.Flush()
	handler.Handle(flushErr)
	return
}

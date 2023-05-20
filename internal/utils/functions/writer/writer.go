package writer

import (
	"YadroTestWork/internal/utils/structures"
	"bufio"
	"fmt"
	"os"
)

func Write(path string, output []structures.Pair[string, int]) (err error) {
	file, createErr := os.Create(path)
	if createErr != nil {
		return createErr
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			err = closeErr
			return
		}
	}(file)

	writer := bufio.NewWriter(file)

	for _, kv := range output {
		_, outErr := fmt.Fprintf(writer, "%s: %d\n", kv.First, kv.Second)
		if outErr != nil {
			return outErr
		}
	}

	flushErr := writer.Flush()
	if flushErr != nil {
		return flushErr
	}

	return
}

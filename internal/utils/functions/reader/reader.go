package reader

import (
	"YadroTestWork/internal/utils/functions/handler"
	"bufio"
	"os"
)

func Read(path string) (array []string, err error) {
	file, openErr := os.Open(path)
	handler.Handle(openErr)
	defer func(file *os.File) {
		handler.Handle(file.Close())
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	return
}

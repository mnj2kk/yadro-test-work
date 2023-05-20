package reader

import (
	"bufio"
	"os"
)

func Read(path string) (array []string, err error) {
	array = make([]string, 0)

	file, openErr := os.Open(path)
	if openErr != nil {
		err = openErr
		return
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			err = closeErr
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	return
}

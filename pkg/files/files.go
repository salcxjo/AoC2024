package files

import (
	"bufio"
	"os"
)

func ReadLines(path string, callback func(string)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		callback(line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

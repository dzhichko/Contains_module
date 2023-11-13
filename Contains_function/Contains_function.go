package Contains_function

import (
	"errors"
	"io"
	"os"
	"strings"
)

func Contains(directory string, match string) (bool, error) {

	file, err := os.Open(directory)
	if err != nil {
		return false, errors.New("Cannot to open file")
	}

	text, err := io.ReadAll(file)
	if err != nil {
		return false, errors.New("Cannot to read file")
	}

	return strings.Contains(string(text), match), nil

}

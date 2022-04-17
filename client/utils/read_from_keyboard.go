package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadFromKeyboard() (string, error) {
	var err error
	var input string

	reader := bufio.NewReader(os.Stdin)

	input, err = reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	input = strings.Replace(input, "\r\n", "", -1)
	input = strings.Replace(input, "\n", "", -1)

	return input, nil
}

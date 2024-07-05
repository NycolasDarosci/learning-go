package fileutils

import "os"

func ReadTextFile(file string) (string, error) {
	content, err := os.ReadFile(file)

	// errors desing pattern (error management)
	if err != nil {
		return "", err
	} else {
		return string(content), nil
	}
}

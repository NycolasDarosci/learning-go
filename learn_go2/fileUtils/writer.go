package fileutils

import "os"

func WriteToFile(filePath string, data string) error {
	return os.WriteFile(filePath, []byte(data), 0666)
}

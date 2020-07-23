package logger

import "os"

// ref : https://gist.github.com/ivanzoid/5040166bb3f0c82575b52c2ca5f5a60c
func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

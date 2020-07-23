package logger

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
	log.SetOutput(os.Stdout)
}

func initFile(path, name string) *os.File {
	err := makeDirectoryIfNotExists(path)
	if err != nil {
		log.Fatal()
	}

	fullPath := fmt.Sprintf("%s/%s.txt", path, name)

	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	log.SetOutput(writer)

	return file
}

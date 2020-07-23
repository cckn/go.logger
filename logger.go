package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Logger() *log.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}
	fileName := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcName := runtime.FuncForPC(pc).Name()
	fn := funcName[strings.LastIndex(funcName, ".")+1:]
	return log.WithField("file", fileName).WithField("function", fn)
}

func init() {
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
	log.SetOutput(os.Stdout)
}

func FileLogger(path, name string) *os.File {
	fullPath := fmt.Sprintf("%s/%s.txt", path, name)
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	log.SetOutput(writer)

	return file
}

// ref : https://gist.github.com/mattes/d13e273314c3b3ade33f
func hasDirectory(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

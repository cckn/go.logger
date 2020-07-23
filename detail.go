package logger

import (
	"runtime"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func detail() *log.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}
	fileName := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcName := runtime.FuncForPC(pc).Name()
	fn := funcName[strings.LastIndex(funcName, ".")+1:]
	return log.WithField("file", fileName).WithField("function", fn)
}

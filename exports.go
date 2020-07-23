package logger

import log "github.com/sirupsen/logrus"

var (
	Detail   *log.Entry = detail()
	InitFile            = initFile
)

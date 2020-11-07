package logs

import "log"

const (
	DEBUG   = 10
	INFO    = 20
	WARNING = 30
	ERROR   = 40
)

var LogLevel = WARNING

func Debug(format string, v ...interface{}) {

	if LogLevel <= DEBUG {
		log.Printf(format, v...)
	}

}

func Info(format string, v ...interface{}) {

	if LogLevel <= INFO {
		log.Printf(format, v...)
	}

}

func Warning(format string, v ...interface{}) {

	if LogLevel <= WARNING {
		log.Printf(format, v...)
	}

}

func Error(format string, v ...interface{}) {

	if LogLevel <= ERROR {
		log.Printf(format, v...)
	}

}

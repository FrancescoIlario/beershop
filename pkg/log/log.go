package log

import "log"

type Logger interface {
	Logf(format string, v ...interface{})
	Logln(v ...interface{})
}

type L struct{}

func (l *L) Logf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *L) Logln(v ...interface{}) {
	log.Println(v...)
}

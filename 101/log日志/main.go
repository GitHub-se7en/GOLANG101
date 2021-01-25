package main

import (
	log "github.com/sirupsen/logrus"
)

type CLog struct {
	clog *log.Logger
}

func (l *CLog) Info(msg string, keysAndValues ...interface{}) {
	l.clog.WithFields(log.Fields{
		"data": keysAndValues,
	}).Info(msg)
}

func (l *CLog) Error(err error, msg string, keysAndValues ...interface{}) {
	l.clog.WithFields(log.Fields{
		"msg":  msg,
		"data": keysAndValues,
	}).Warn(msg)
}

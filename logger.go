package logger

import (
    "log/syslog"
)

func Info(msg string) {
    w, err := syslog.Dial("", "", syslog.LOG_INFO, "a2n logger ")
    defer w.Close()
    cherr(err)
    w.Info(msg)
}

func Debug(msg string) {
    w, err := syslog.Dial("", "", syslog.LOG_DEBUG, "a2n logger ")
    defer w.Close()
    cherr(err)
    w.Debug(msg)
}

func Notice(msg string) {
    w, err := syslog.Dial("", "", syslog.LOG_NOTICE, "a2n logger ")
    defer w.Close()
    cherr(err)
    w.Notice(msg)
}

func Warning(msg string) {
    w, err := syslog.Dial("", "", syslog.LOG_WARNING, "a2n logger ")
    defer w.Close()
    cherr(err)
    w.Warning(msg)
}

func Err(msg string) {
    w, err := syslog.Dial("", "", syslog.LOG_ERR, "a2n logger ")
    defer w.Close()
    cherr(err)
    w.Err(msg)
}

func Crit(msg string) {
    w, err := syslog.Dial("", "", syslog.LOG_CRIT, "a2n logger ")
    defer w.Close()
    cherr(err)
    w.Crit(msg)
}

func Alert(msg string) {
    w, err := syslog.Dial("", "", syslog.LOG_ALERT, "a2n logger ")
    defer w.Close()
    cherr(err)
    w.Alert(msg)
}

func Emerg(msg string) {
    w, err := syslog.Dial("", "", syslog.LOG_EMERG, "a2n logger ")
    defer w.Close()
    cherr(err)
    w.Emerg(msg)
}

func cherr(err error) {
    if err != nil {
	panic(err)
    }
}

package logger

import (
    "log/syslog"
)

type Logger struct {
    prefix string
}

func New(prefix string) *Logger {
    return &Logger {prefix}
}

func (l *Logger) Info(msg string) {
    l.log(msg, syslog.LOG_INFO)
}

func (l *Logger) Debug(msg string) {
    l.log(msg, syslog.LOG_DEBUG)
}

func (l *Logger) Notice(msg string) {
    l.log(msg, syslog.LOG_NOTICE)
}

func (l *Logger) Warning(msg string) {
    l.log(msg, syslog.LOG_WARNING)
}

func (l *Logger) Err(msg string) {
    l.log(msg, syslog.LOG_ERR)
}

func (l *Logger) Crit(msg string) {
    l.log(msg, syslog.LOG_CRIT)
}

func (l *Logger) Alert(msg string) {
    l.log(msg, syslog.LOG_ALERT)
}

func (l *Logger) Emerg(msg string) {
    l.log(msg, syslog.LOG_EMERG)
}

func (l *Logger) log(text string, priority syslog.Priority) {
    w, err := syslog.New(priority, l.prefix)
    defer w.Close()
    if err != nil {
	panic(err)
    }

    switch priority {
	case syslog.LOG_DEBUG:
	w.Debug(text)

	case syslog.LOG_INFO:
	w.Info(text)

	case syslog.LOG_NOTICE:
	w.Notice(text)

	case syslog.LOG_WARNING:
	w.Warning(text)

	case syslog.LOG_ERR:
	w.Err(text)

	case syslog.LOG_CRIT:
	w.Crit(text)

	case syslog.LOG_ALERT:
	w.Alert(text)

	case syslog.LOG_EMERG:
	w.Emerg(text)
    }
}

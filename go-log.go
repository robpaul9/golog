package golog

import (
	"github.com/sirupsen/logrus"
)

type (
	Config struct {
		ServiceName  string
		CustomLogger *logrus.Logger
		LogLevel     *logrus.Level
	}

	Logger interface {
		logrus.FieldLogger
		InfoWithFields(string, ...LogFields)
		DebugWithFields(string, ...LogFields)
		WarnWithFields(string, ...LogFields)
		ErrorWithFields(error, ...LogFields)
	}

	LogFields map[string]interface{}

	logrusLogger struct {
		*logrus.Entry
	}
)

func New(c Config) Logger {
	var baseLogger *logrus.Logger

	if baseLogger = c.CustomLogger; baseLogger == nil {
		baseLogger = logrus.New()
	}

	baseLogger.SetFormatter(&logrus.JSONFormatter{})

	if c.LogLevel != nil {
		baseLogger.SetLevel(*c.LogLevel)
	}

	baseEntry := baseLogger.WithFields(logrus.Fields{
		"service": c.ServiceName,
	})

	return &logrusLogger{baseEntry}
}

func (l *logrusLogger) InfoWithFields(msg string, f ...LogFields) {
	l.withFields(f).Infoln(msg)
}

func (l *logrusLogger) DebugWithFields(msg string, f ...LogFields) {
	l.withFields(f).Debugln(msg)
}

func (l *logrusLogger) WarnWithFields(msg string, f ...LogFields) {
	l.withFields(f).Warnln(msg)
}

func (l *logrusLogger) ErrorWithFields(err error, f ...LogFields) {
	l.withFields(f).Errorln(err)
}

func (l *logrusLogger) withFields(fields []LogFields) *logrus.Entry {
	logFields := make(logrus.Fields)
	for _, f := range fields {
		for k, v := range f {
			logFields[k] = v
		}
	}
	return l.WithFields(logFields)
}

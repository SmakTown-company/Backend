package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"runtime"
)

type LogService struct {
	output io.Writer
	logger *logrus.Logger
}

var Logger *LogService

func NewLogService(output io.Writer, lvl string) {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		level = logrus.InfoLevel
		err = nil
	}
	logger := logrus.New()
	logger.SetOutput(output)
	logger.SetLevel(level)
	Logger = &LogService{
		output: output,
		logger: logger,
	}
}

// Print логирует сообщение с ID и контекстом
func (l *LogService) Info(msg string) {
	l.logger.WithFields(logrus.Fields{}).Info(msg)

}
func (l *LogService) Warn(msg string) {
	l.logger.WithFields(logrus.Fields{}).Warn(msg)
}

func (l *LogService) SetFormat(writer io.Writer) {
	Logger.output = writer
	Logger.logger.SetOutput(writer)
	l.output = writer
	l.logger.SetOutput(writer)
}

func (l *LogService) Debug(msg string) {
	pc, _, _, _ := runtime.Caller(1)
	callerName := runtime.FuncForPC(pc).Name()
	l.logger.WithFields(logrus.Fields{
		"fromFunc": callerName,
	}).Debug(msg)

}
func (l *LogService) SetLevel(lvl string) {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		level = logrus.InfoLevel
		err = nil
	}
	l.logger.SetLevel(level)

}

func MakeLog(msg string, err error) string {
	if err != nil {
		return fmt.Sprintf("%s:Ошибка: %s", msg, err.Error())
	} else {
		return msg
	}
}

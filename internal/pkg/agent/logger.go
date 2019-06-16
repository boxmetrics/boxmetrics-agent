package agent

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

// Log Global Log variable
var Log boxlogger

// Logger object struct
type Logger struct {
	Type   string `mapstructure:"type"`
	Format string `mapstructure:"format"`
	Level  string `mapstructure:"level"`
	Path   string `mapstructure:"path"`
}

type boxlogger struct {
	loggers []*logrus.Logger
	number  int
}

func newLogger() boxlogger {
	log := boxlogger{make([]*logrus.Logger, 0), 0}

	return log
}

// AddLogger add a new boxlogger
func (l *boxlogger) AddLogger(logger *logrus.Logger) {
	l.loggers = append(l.loggers, logger)
	l.number = len(l.loggers)
}

// AddLogger add a new boxlogger
func (l *boxlogger) HasLogger() bool {
	if l.number > 0 {
		return true
	}

	return false
}

func (l *boxlogger) WithField(key string, value interface{}) *entries {
	formatLoggers := newEntries()
	for _, logger := range l.loggers {
		formatLoggers.loggers = append(formatLoggers.loggers, logger.WithField(key, value))
	}

	return &formatLoggers
}

func (l *boxlogger) WithFields(fields logrus.Fields) *entries {
	formatLoggers := newEntries()
	for _, logger := range l.loggers {
		formatLoggers.loggers = append(formatLoggers.loggers, logger.WithFields(fields))
	}

	return &formatLoggers
}

func (l *boxlogger) Trace(args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Trace(args...)
	}
}
func (l *boxlogger) Debug(args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Debug(args...)
	}
}
func (l *boxlogger) Info(args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Info(args...)
	}
}
func (l *boxlogger) Warn(args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Warn(args...)
	}
}
func (l *boxlogger) Error(args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Error(args...)
	}
}
func (l *boxlogger) Fatal(args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Log(logrus.FatalLevel, args...)
	}
	os.Exit(1)
}
func (l *boxlogger) Panic(args ...interface{}) {
	for _, logger := range l.loggers {
		logger.Log(logrus.PanicLevel, args...)
	}
	panic(fmt.Sprint(args...))
}

type entries struct {
	loggers []*logrus.Entry
	number  int
}

func newEntries() entries {
	log := entries{make([]*logrus.Entry, 0), 0}

	return log
}

func (e *entries) Trace(args ...interface{}) {
	for _, logger := range e.loggers {
		logger.Trace(args...)
	}
}
func (e *entries) Debug(args ...interface{}) {
	for _, logger := range e.loggers {
		logger.Debug(args...)
	}
}
func (e *entries) Info(args ...interface{}) {
	for _, logger := range e.loggers {
		logger.Info(args...)
	}
}
func (e *entries) Warn(args ...interface{}) {
	for _, logger := range e.loggers {
		logger.Warn(args...)
	}
}
func (e *entries) Error(args ...interface{}) {
	for _, logger := range e.loggers {
		logger.Error(args...)
	}
}
func (e *entries) Fatal(args ...interface{}) {
	for _, logger := range e.loggers {
		logger.Log(logrus.FatalLevel, args...)
	}
	os.Exit(1)
}
func (e *entries) Panic(args ...interface{}) {
	for _, logger := range e.loggers {
		logger.Log(logrus.PanicLevel, args...)
	}
	panic(fmt.Sprint(args...))
}

func getLoggers() ([]Logger, error) {
	var L []Logger
	if Config.IsSet("loggers") {

		err := Config.UnmarshalKey("loggers", &L)

		if err != nil {
			return L, err
		}
		return L, nil
	}

	return L, errors.New("No Logger configuration")
}

func initLogger() {
	Log = newLogger()

	loggers, err := getLoggers()

	if err == nil {
		for _, logger := range loggers {
			l, err := CreateLogger(logger)

			if err == nil {
				Log.AddLogger(l)
			} else {
				if Log.HasLogger() {
					Log.WithField("error", err).Fatal("Unable to create boxlogger")
				} else {
					log.Fatalln("Unable to create boxlogger:", err)
				}
			}
		}
	}

	Log.Info("loggers initialized")
}

func parseFormatter(format string) (logrus.Formatter, error) {
	var formatter logrus.Formatter
	switch format {
	case "text":
		formatter = new(logrus.TextFormatter)
		formatter.(*logrus.TextFormatter).DisableLevelTruncation = true
	case "json":
		formatter = new(logrus.JSONFormatter)
	default:
		return formatter, errors.New("Logger Formatter not support")
	}

	return formatter, nil
}

func parseType(logger Logger) (io.Writer, error) {
	var writer io.Writer
	switch logger.Type {
	case "console":
		writer = os.Stdout
	case "file":
		file, err := os.OpenFile(logger.Path, os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			return nil, err
		}

		writer = file
	default:
		return nil, errors.New("Logger Type not support")
	}

	return writer, nil
}

// CreateLogger create a new logrus Logger object
func CreateLogger(l Logger) (*logrus.Logger, error) {
	logger := logrus.New()

	out, err := parseType(l)

	if err != nil {
		return nil, err
	}

	logger.Out = out

	format, err := parseFormatter(l.Format)

	if err != nil {
		return nil, err
	}

	logger.Formatter = format

	level, err := logrus.ParseLevel(l.Level)

	if err != nil {
		return nil, err
	}

	logger.Level = level

	return logger, nil
}

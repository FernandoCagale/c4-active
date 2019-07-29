package logger

import "github.com/sirupsen/logrus"

var logger = logrus.New()

//Fields to logs
type Fields map[string]interface{}

//Info message
func Info(args ...interface{}) {
	logger.Info(args...)
}

//Debug message
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

//Warn message
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

//Error message
func Error(args ...interface{}) {
	logger.Error(args...)
}

//Fatal message
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

//WithFields adds a field
//logger.WithFields(logger.Fields{
//  "ID":   "id",
//}).Info("info")
func WithFields(fields Fields) *logrus.Entry {
	fieldsMap := make(logrus.Fields)
	for index, property := range fields {
		fieldsMap[index] = property
	}
	return logrus.WithFields(fieldsMap)
}

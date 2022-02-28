package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

//Event is public function to create logging
func Event(traceId string, text ...string) {
	msgText := "[" + strings.Join(text, "][") + "]"
	logrus.SetLevel(logrus.InfoLevel)
	logrus.WithField("trace_id", traceId).Info(msgText)
}

//Message is public function to create logging
func Message(traceId string, text ...string) {
	msgText := "[" + strings.Join(text, "][") + "]"
	logrus.SetLevel(logrus.TraceLevel)
	logrus.WithField("trace_id", traceId).Trace(msgText)
}

//Warning is public function to create logging
func Warning(traceId string, text ...string) {
	msgText := "[" + strings.Join(text, "][") + "]"
	logrus.SetLevel(logrus.WarnLevel)
	logrus.WithField("trace_id", traceId).Warn(msgText)
}

//Error is public function to create logging
func Error(traceId string, err error, text ...string) {
	msgText := "[" + strings.Join(text, "][") + "]"
	errText := errors.Wrap(err, err.Error())
	stackTrace := "[" + strings.Replace(strings.Replace(fmt.Sprintf("%+v", errText), "\n\t", " ", -1),
		"\n", " | ", -1) + "]"
	logrus.SetLevel(logrus.ErrorLevel)
	logrus.WithField("trace_id", traceId).WithError(err).Error(msgText + stackTrace)
}

//Fatal is public function to create logging
func Fatal(traceId string, err error, text ...string) {
	msgText := "[" + strings.Join(text, "][") + "]"
	errText := errors.Wrap(err, err.Error())
	stackTrace := "[" + strings.Replace(strings.Replace(fmt.Sprintf("%+v", errText), "\n\t", " ", -1),
		"\n", " | ", -1) + "]"
	logrus.SetLevel(logrus.FatalLevel)
	logrus.WithField("trace_id", traceId).WithError(err).Fatal(msgText + stackTrace)
}

// SetupLogging is used to set up logging system
func SetupLogging(appModeDebug bool) {
	logrus.SetOutput(os.Stdout)
	if appModeDebug {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

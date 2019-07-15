package logger

import (
	"bytes"
	"fmt"
	"log/syslog"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	lg "github.com/sirupsen/logrus/hooks/syslog"

	"backend/libs/configuration"
	"backend/libs/util"
)

var log *logrus.Entry
var requestID string

func init() {
	l := logrus.New()
	config := configuration.Get()
	out := config.GetEnvConfString("logger.output")
	if out == "" || out == "stdout" {
		l.SetOutput(os.Stdout)
	} else {
		f, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			panic(err)
		}
		l.SetOutput(f)
	}

	var formatter logrus.Formatter
	format := config.GetEnvConfString("logger.format")
	if format == "" || format == "text" {
		formatter = &logrus.TextFormatter{}
	}

	if format == "json" {
		formatter = &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "time",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "message",
				logrus.FieldKeyFunc:  "caller",
			},
		}
		server := config.GetEnvConfString("logger.server")
		protocol := config.GetEnvConfString("logger.protocol")
		program := config.GetEnvConfString("logger.app_name")
		if program == "" {
			program = "backend"
		}
		hook, err := lg.NewSyslogHook(protocol, server, syslog.LOG_WARNING|syslog.LOG_DAEMON, program)
		if err == nil {
			l.Hooks.Add(hook)
		}
	}
	l.SetFormatter(formatter)
	log = logrus.NewEntry(l)
}

func SetRequestID(id string) {
	if id == "" {
		id = util.Random(10)
	}
	requestID = id
	log = log.WithField("request_id", id)
}

func GetRequestID() string {
	return requestID
}

func Error(v ...interface{}) {
	message := toString(v)
	log.WithField("caller", getCaller()).Error(message)
}

func Info(v ...interface{}) {
	message := toString(v)
	log.Info(message)
}

func Fatal(v ...interface{}) {
	message := toString(v)
	log.WithField("func", getCaller()).Fatal(message)
}

func getCaller() string {
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return fmt.Sprintf("%s", details.Name())
	}
	return ""
}

func toString(v ...interface{}) string {
	var buf bytes.Buffer
	for _, s := range v {
		buf.WriteString(fmt.Sprintf("%+v", s))
	}
	value := strings.Replace(buf.String(), "[", "", 1)
	return strings.Replace(value, "]", "", 1)
}

func getFromEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

var (
	tlog *logrus.Logger

	LevelList = map[string]logrus.Level{
		"PANIC": logrus.PanicLevel,
		"FATAL": logrus.FatalLevel,
		"ERROR": logrus.ErrorLevel,
		"WARN":  logrus.WarnLevel,
		"INFO":  logrus.InfoLevel,
		"DEBUG": logrus.DebugLevel,
		"TRACE": logrus.TraceLevel,
	}
)

// ***********************************************************************
func LogToFile(FileName string) (error, io.Writer) {
	var (
		f   io.Writer
		err error = nil
	)
	f, err = os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error opening log file: " + err.Error())
		return err, f
	}
	return err, f
}

// ***********************************************************************
func LogToStdout() io.Writer {
	return os.Stdout
}

// ***********************************************************************
func LogLevel(DebugLevel string) {
	tlog.SetLevel(str2lvl(DebugLevel))
	tlog.Infoln("Logging level: " + DebugLevel)
}

// ***********************************************************************
func InitLogger(DebugLevel string, FileName string) {
	var (
		LogStream io.Writer
		err       error
	)

	tlog = logrus.New()
	if strings.ToLower(FileName) == "stdout" {
		tlog.SetOutput(os.Stdout)
	} else {
		err, LogStream = LogToFile(FileName)
		if err == nil {
			tlog.SetOutput(LogStream)
		} else {
			tlog.SetOutput(os.Stdout)
		}
	}

	tlog.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	LogLevel(DebugLevel)
	tlog.Infoln("Logging initialized")
}

// ***********************************************************************
func str2lvl(StrLvl string) logrus.Level {
	if _, ok := LevelList[strings.ToUpper(StrLvl)]; !ok {
		fmt.Println("ERROR: Unknown logging level (" + StrLvl + "). Set default (DEBUG)")
		return logrus.DebugLevel
	}
	return LevelList[strings.ToUpper(StrLvl)]
}

// ***********************************************************************
// ***********************************************************************
// ***********************************************************************

func Tracef(format string, args ...interface{}) {
	tlog.Logf(logrus.TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	tlog.Logf(logrus.DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	tlog.Logf(logrus.InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	tlog.Logf(logrus.WarnLevel, format, args...)
}

func Warningf(format string, args ...interface{}) {
	tlog.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	tlog.Logf(logrus.ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	tlog.Logf(logrus.FatalLevel, format, args...)
	tlog.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	tlog.Logf(logrus.PanicLevel, format, args...)
}

func Trace(args ...interface{}) {
	tlog.Log(logrus.TraceLevel, args...)
}

func Debug(args ...interface{}) {
	tlog.Log(logrus.DebugLevel, args...)
}

func Info(args ...interface{}) {
	tlog.Log(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	tlog.Log(logrus.WarnLevel, args...)
}

func Error(args ...interface{}) {
	tlog.Log(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	tlog.Log(logrus.FatalLevel, args...)
	tlog.Exit(1)
}

func Panic(args ...interface{}) {
	tlog.Log(logrus.PanicLevel, args...)
}

func Traceln(args ...interface{}) {
	tlog.Logln(logrus.TraceLevel, args...)
}

func Debugln(args ...interface{}) {
	tlog.Logln(logrus.DebugLevel, args...)
}

func Infoln(args ...interface{}) {
	tlog.Logln(logrus.InfoLevel, args...)
}

func Warnln(args ...interface{}) {
	tlog.Logln(logrus.WarnLevel, args...)
}

func Warningln(args ...interface{}) {
	tlog.Warnln(args...)
}

func Errorln(args ...interface{}) {
	tlog.Logln(logrus.ErrorLevel, args...)
}

func Panicln(args ...interface{}) {
	tlog.Logln(logrus.PanicLevel, args...)
}

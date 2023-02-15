package fmtlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

var (
	fmtlog *logrus.Logger

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
	fmtlog.SetLevel(str2lvl(DebugLevel))
	fmtlog.Infoln("Logging level: " + DebugLevel)
}

// ***********************************************************************
func InitLogger(DebugLevel string, FileName string) {
	var (
		LogStream io.Writer
		err       error
	)

	fmtlog = logrus.New()
	if strings.ToLower(FileName) == "stdout" {
		fmtlog.SetOutput(os.Stdout)
	} else {
		err, LogStream = LogToFile(FileName)
		if err == nil {
			fmtlog.SetOutput(LogStream)
		} else {
			fmtlog.SetOutput(os.Stdout)
		}
	}

	fmtlog.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	LogLevel(DebugLevel)
	fmtlog.Infoln("Logging initialized")
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
	fmtlog.Logf(logrus.TraceLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	fmtlog.Logf(logrus.DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	fmtlog.Logf(logrus.InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	fmtlog.Logf(logrus.WarnLevel, format, args...)
}

func Warningf(format string, args ...interface{}) {
	fmtlog.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	fmtlog.Logf(logrus.ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	fmtlog.Logf(logrus.FatalLevel, format, args...)
	fmtlog.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	fmtlog.Logf(logrus.PanicLevel, format, args...)
}

func Trace(args ...interface{}) {
	fmtlog.Log(logrus.TraceLevel, args...)
}

func Debug(args ...interface{}) {
	fmtlog.Log(logrus.DebugLevel, args...)
}

func Info(args ...interface{}) {
	fmtlog.Log(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	fmtlog.Log(logrus.WarnLevel, args...)
}

func Error(args ...interface{}) {
	fmtlog.Log(logrus.ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	fmtlog.Log(logrus.FatalLevel, args...)
	fmtlog.Exit(1)
}

func Panic(args ...interface{}) {
	fmtlog.Log(logrus.PanicLevel, args...)
}

func Traceln(args ...interface{}) {
	fmtlog.Logln(logrus.TraceLevel, args...)
}

func Debugln(args ...interface{}) {
	fmtlog.Logln(logrus.DebugLevel, args...)
}

func Infoln(args ...interface{}) {
	fmtlog.Logln(logrus.InfoLevel, args...)
}

func Warnln(args ...interface{}) {
	fmtlog.Logln(logrus.WarnLevel, args...)
}

func Warningln(args ...interface{}) {
	fmtlog.Warnln(args...)
}

func Errorln(args ...interface{}) {
	fmtlog.Logln(logrus.ErrorLevel, args...)
}

func Panicln(args ...interface{}) {
	fmtlog.Logln(logrus.PanicLevel, args...)
}

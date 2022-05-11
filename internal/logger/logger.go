package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

const containerName = `logger`

func GetContainerName() string {
	return containerName
}

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.ReportCaller = false

	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		QuoteEmptyFields: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			file, line := frame.Func.FileLine(frame.PC)
			return frame.Function, fmt.Sprintf("%s:%d", filepath.Base(file), line)
		},
	})

	go handleSIGUSR2(logger)
	return logger
}

func handleSIGUSR2(logger *logrus.Logger) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR2)
	for range ch {
		level := logger.GetLevel()
		switch level {
		case logrus.DebugLevel:
			logger.Warn("switching log level to INFO")
			logger.SetLevel(logrus.InfoLevel)
		default:
			logger.Warn("switching log level to DEBUG")
			logger.SetLevel(logrus.DebugLevel)
		}
	}
}

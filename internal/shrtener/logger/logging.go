package logger

import (
	"io"
	"log"
	"os"

	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/sirupsen/logrus"
)

func WithLogLevel(cfg *configuration.Logging) logrus.Level {
	switch cfg.Level {
	case "info":
		return logrus.InfoLevel
	case "warning":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "debug":
		return logrus.DebugLevel
	default:
		log.Print("applying default log level ERROR")
		return logrus.ErrorLevel
	}
}

func WithOutput(cfg *configuration.Logging) io.Writer {
	switch cfg.Output {
	case "stdout":
		return os.Stdout
	case "file":
		f, err := os.OpenFile(cfg.DirPath+"/logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic("could not open log file. Path: " + cfg.DirPath)
		}
		return f
	default:
		log.Print("applying default output - stdout")
		return os.Stdout
	}
}

func WithFormatter(formatter string) logrus.Formatter {
	switch formatter {
	case "json":
		return &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: "message",
			}}
	default:
		return &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: "message",
			}}
	}
}

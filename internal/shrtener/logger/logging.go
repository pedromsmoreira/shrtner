package logger

import (
	"github.com/pedromsmoreira/shrtener/internal/shrtener/configuration"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func ConfigureLogrus(cfg *configuration.Settings) {
	setFormatter(cfg.Logging.Format)
	setOutput(cfg.Logging)
	setLogLevel(cfg.Logging)
}

func setLogLevel(cfg *configuration.Logging) {
	switch cfg.Level {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		log.Print("applying default log level ERROR")
		logrus.SetLevel(logrus.ErrorLevel)
	}
}

func setOutput(cfg *configuration.Logging) {
	switch cfg.Output {
	case "stdout":
		logrus.SetOutput(os.Stdout)
	case "file":
		f, err := os.OpenFile(cfg.DirPath+"/logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic("could not open log file. Path: " + cfg.DirPath)
		}
		logrus.SetOutput(f)
	default:
		log.Print("applying default output - stdout")
		logrus.SetOutput(os.Stdout)
	}
}

func setFormatter(formatter string) {
	switch formatter {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: "message",
			},
		})
	default:
		log.Print("applying default formatter - json")
		logrus.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg: "message",
			},
		})
	}
}

func Info(message string, args ...interface{}) {
	if args != nil {
		logrus.WithFields(
			logrus.Fields{
				"data": args,
			},
		).Info(message)
		return
	}
	logrus.Info(message)
}

func Warning(message string, args ...interface{}) {
	if args != nil {
		logrus.WithFields(
			logrus.Fields{
				"data": args,
			},
		).Warning(message)
		return
	}
	logrus.Warning(message)
}

func Error(message string, args ...interface{}) {
	if args != nil {
		logrus.WithFields(
			logrus.Fields{
				"data": args,
			},
		).Error(message)
		return
	}
	logrus.Error(message)
}

func Debug(message string, args ...interface{}) {
	if args != nil {
		logrus.WithFields(
			logrus.Fields{
				"data": args,
			},
		).Debug(message)
		return
	}
	logrus.Debug(message)
}

package pkg

import (
	"fiber-boilerplate/config"
	"log"
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitializeLogger sets up the zap logger with daily log rotation
func InitLogger() {
	STAGE_STATUS := config.GetEnvValue("STAGE_STATUS")

	// Create logs directory if it doesn't exist
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			log.Fatalf("Error creating log directory: %v", err)
		}
	}

	// Build log file path based on current date
	logFileName := path.Join(logDir, time.Now().Format("2006-01-02")+".log")

	// Create a file writer
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	// Set up zapcore to log to both console and file
	fileWriter := zapcore.AddSync(file)
	consoleWriter := zapcore.AddSync(os.Stdout)

	// Set log level based on environment
	var logLevel zapcore.Level
	if STAGE_STATUS == "prod" {
		logLevel = zapcore.InfoLevel
	} else {
		logLevel = zapcore.DebugLevel
	}

	// Define encoder configuration
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	// Set up core
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileWriter, logLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleWriter, logLevel),
	)

	// Build the logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

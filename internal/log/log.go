/*
Copyright © 2022 Vladislav Pavkin [phylocko@gmail.com]
*/
package log

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var L *zap.Logger

func FileLogger(filename string) *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(config)

	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("unable to open log file: %s\n", err)
	}

	logWriter := zapcore.AddSync(logFile)
	logLevel := zap.NewAtomicLevel()

	core := zapcore.NewCore(fileEncoder, logWriter, logLevel)

	stackOption := zap.AddStacktrace(logLevel)
	callerOption := zap.AddCaller()

	logger := zap.New(core, callerOption, stackOption)

	L = logger
	return logger
}

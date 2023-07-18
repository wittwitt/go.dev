package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Create log file writers
	file1, _ := os.Create("file1.log")
	file2, _ := os.Create("file2.log")
	defer file1.Close()
	defer file2.Close()

	// Configure encoder
	encoderConfig := zap.NewDevelopmentEncoderConfig()

	// Create core for log file 1 with Info level
	core1 := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(file1),
		zapcore.InfoLevel,
	)
	logger1 := zap.New(core1)

	// Create core for log file 2 with Debug level
	core2 := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(file2),
		zapcore.DebugLevel,
	)
	logger2 := zap.New(core2)

	// Merge logger1 and logger2 into logger3
	logger3 := zap.New(zapcore.NewTee(logger1.Core(), logger2.Core()))

	// Usage examples
	logger3.Debug("This message will be logged to both file2.log and console")

	// Flush and sync
	defer func() {
		_ = logger1.Sync()
		_ = logger2.Sync()
		_ = logger3.Sync()
	}()
}

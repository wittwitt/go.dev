package main

import (
	"log/slog"
	"os"
)

func main() {
	slog.Info("1", "GODEV_TEST_VAR1", os.Getenv("GODEV_TEST_VAR1"))
	slog.Info("2", "GODEV_TEST_VAR2", os.Getenv("GODEV_TEST_VAR2"))
}

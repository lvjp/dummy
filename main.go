package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))
	slog.Info("Dummy app...")
}

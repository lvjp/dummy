package main

import (
	"os"

	"github.com/lvjp/dummy/cmd"
	"golang.org/x/exp/slog"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))

	cmd.Execute()
}

package main

import (
	"context"
	"os"

	"github.com/xLeSHka/life/internal/application"
)

func main() {
	ctx := context.Background()
	// Exit приводит к завершению программы с заданным кодом
	os.Exit(mainWithExitCode(ctx))
}

func mainWithExitCode(ctx context.Context) int {
	cfg := application.Config{
		Width:  50,
		Height: 50,
	}
	app := application.New(cfg)

	return app.Run(ctx)
}

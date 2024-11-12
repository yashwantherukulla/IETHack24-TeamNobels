package signals

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func HandleCancellation(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		cancel()
	}()

	return ctx
}

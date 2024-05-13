package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func RegisterInterrupt() <-chan struct{} {
	sigChan := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sigChan
		done <- struct{}{}
	}()
	return done
}

func WaitForTerminate(ctx context.Context, teardown func(msg string)) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	waitSignal(ctx, c, teardown)
}

// waitSignal wait for signal - SIGINT and SIGTERM
func waitSignal(ctx context.Context, c chan os.Signal, teardown func(msg string)) {
	select {
	case s := <-c:
		teardown(fmt.Sprintf("terminating: got signal %v", s))
	case <-ctx.Done():
		teardown("terminating: context done")
	}
}

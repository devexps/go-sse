package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/devexps/go-micro/transport/sse/v2"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := context.Background()

	s := sse.NewServer(
		sse.WithAddress(":8080"),
	)
	defer s.Stop(ctx)

	s.HandleServeHTTP("/events")

	s.CreateStream("demo")

	go func() {
		s.Start(ctx)
	}()

	s.Publish("demo", &sse.Event{Data: []byte("message")})

	<-interrupt
}

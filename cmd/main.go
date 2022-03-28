//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Mirobidjon/websocket-service/config"
	"github.com/Mirobidjon/websocket-service/pkg/logger"
	"github.com/Mirobidjon/websocket-service/router"
	"github.com/Mirobidjon/websocket-service/socket"
	"github.com/valyala/fasthttp"
)

func main() {
	cfg := config.Load()
	log := logger.New(logger.LevelDebug, "websocket_service")
	hub := socket.NewHub(log)

	router := router.NewHandler(hub)

	server := fasthttp.Server{
		Handler: router.Handler,
	}

	go server.ListenAndServe(cfg.HttpPort)
	go hub.Run()

	fmt.Println("Websocket service started at " + cfg.HttpPort)
	fmt.Println("Visit http://localhost:8080")

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh
	signal.Stop(sigCh)
	signal.Reset(os.Interrupt)
	server.Shutdown()
}

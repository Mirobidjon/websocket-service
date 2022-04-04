//go:build ignore
// +build ignore

package main

import (
	"net"
	"os"
	"os/signal"

	"github.com/Mirobidjon/udevs_websocket_service/config"
	pb "github.com/Mirobidjon/udevs_websocket_service/genproto/websocket_service"
	"github.com/Mirobidjon/udevs_websocket_service/grpc/service"
	"github.com/Mirobidjon/udevs_websocket_service/pkg/logger"
	"github.com/Mirobidjon/udevs_websocket_service/router"
	"github.com/Mirobidjon/udevs_websocket_service/socket"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()
	log := logger.New(logger.LevelDebug, "websocket_service")
	hub := socket.NewHub(log)

	router := router.NewHandler(hub)

	server := fasthttp.Server{
		Handler: router.Handler,
	}

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return
	}

	websocketService := service.NewWebsocketService(log, hub)

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterWebSocketServiceServer(s, websocketService)

	go server.ListenAndServe(cfg.HttpPort)
	go hub.Run()

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Error("error while listening: %v", logger.Error(err))
		}
	}()

	log.Info("Websocket service started at " + cfg.HttpPort + "HTTP port")
	log.Info("Websocket service started at " + cfg.RPCPort + "GRPC port")

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh
	signal.Stop(sigCh)
	signal.Reset(os.Interrupt)
	server.Shutdown()
}

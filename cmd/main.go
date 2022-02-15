//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Mirobidjon/websocket-service/pkg/logger"
	"github.com/Mirobidjon/websocket-service/socket"
	"github.com/buaazp/fasthttprouter"
	"github.com/dgrr/fastws"
	"github.com/valyala/fasthttp"
)

func main() {
	log := logger.New(logger.LevelDebug, "websocket-service")
	hub := socket.NewHub(log)
	router := fasthttprouter.New()
	router.GET("/", rootHandler)
	router.GET("/ws", func(ctx *fasthttp.RequestCtx) {
		if !customMiddleware(ctx) {
			return
		}

		fastws.Upgrade(hub.WsHandler)(ctx)
	})

	server := fasthttp.Server{
		Handler: router.Handler,
	}
	go server.ListenAndServe(":8080")
	go hub.Run()

	fmt.Println("Visit http://localhost:8080")

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh
	signal.Stop(sigCh)
	signal.Reset(os.Interrupt)
	server.Shutdown()
}

func rootHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	ctx.Response.SendFile("index.html")
}

func customMiddleware(ctx *fasthttp.RequestCtx) bool {
	token := ctx.QueryArgs().Peek("token")
	if len(token) == 0 {
		ctx.Error("Unauthorized", fasthttp.StatusUnauthorized)
		return false
	}

	// TODO: check token

	ctx.SetUserValue("user_id", string(token))
	ctx.SetUserValue("session_id", string(token))
	ctx.SetUserValue("room_id", string(token))
	return true
}

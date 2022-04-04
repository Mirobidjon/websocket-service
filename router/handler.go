package router

import (
	"fmt"
	"time"

	"github.com/Mirobidjon/udevs_websocket_service/pkg/logger"
	"github.com/Mirobidjon/udevs_websocket_service/socket"
	"github.com/buaazp/fasthttprouter"
	"github.com/dgrr/fastws"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

func Middleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	log := logger.New(logger.LevelDebug, "websocket_service")

	return func(ctx *fasthttp.RequestCtx) {
		startTime := time.Now()
		h(ctx)

		if ctx.Response.StatusCode() > 400 {
			log.Warn("access", zap.Int("code", ctx.Response.StatusCode()), zap.Duration("time", time.Since(startTime)), zap.ByteString("method", ctx.Method()), zap.ByteString("path", ctx.Path()))
		} else {
			log.Info("access", zap.Int("code", ctx.Response.StatusCode()), zap.Duration("time", time.Since(startTime)), zap.ByteString("method", ctx.Method()), zap.ByteString("path", ctx.Path()))
		}

		log.Info(ctx.Response.String())
	}
}

func NewHandler(hub *socket.Hub) *fasthttprouter.Router {

	router := fasthttprouter.New()
	router.GET("/", rootHandler)
	router.GET("/ws", Middleware(func(ctx *fasthttp.RequestCtx) {
		if !customMiddleware(ctx) {
			return
		}

		fastws.Upgrade(hub.WsHandler)(ctx)
	}))

	return router
}

func rootHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	ctx.Response.SendFile("index.html")
}

func customMiddleware(ctx *fasthttp.RequestCtx) bool {
	token := ctx.QueryArgs().Peek("token")
	fmt.Println("token", string(token))
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

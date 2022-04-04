package service

import (
	"context"

	pb "github.com/Mirobidjon/udevs_websocket_service/genproto/websocket_service"
	"github.com/Mirobidjon/udevs_websocket_service/pkg/helper"
	"github.com/Mirobidjon/udevs_websocket_service/pkg/logger"
	"github.com/Mirobidjon/udevs_websocket_service/socket"
)

type websocketService struct {
	logger logger.Logger
	hub    *socket.Hub
	pb.UnimplementedWebSocketServiceServer
}

func NewWebsocketService(log logger.Logger, hub *socket.Hub) *websocketService {
	return &websocketService{
		logger: log,
		hub:    hub,
	}
}

func (s *websocketService) SendMessage(ctx context.Context, req *pb.Message) (*pb.SendMessageResponse, error) {
	var msg socket.Message

	err := helper.ProtoToStruct(&msg, req)
	if err != nil {
		s.logger.Error("error converting proto message to struct", logger.Error(err), logger.Any("message", req))
		return nil, err
	}

	count := s.hub.Send(msg)

	return &pb.SendMessageResponse{
		SuccessCount: count,
	}, nil
}

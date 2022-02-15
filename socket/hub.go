package socket

import (
	"github.com/Mirobidjon/websocket-service/pkg/logger"
	"github.com/dgrr/fastws"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	rooms      map[string]bool
	log        logger.Logger
}

func NewHub(log logger.Logger) *Hub {
	return &Hub{
		register:   make(chan *Client, 10),
		unregister: make(chan *Client, 10),
		clients:    make(map[*Client]bool),
		rooms:      make(map[string]bool),
		log:        log,
	}
}

func (h *Hub) Send(msg Message) {
	for c := range h.clients {
		switch msg.Type {
		case "broadcast":
			c.Send(msg)
		case "room":
			if c.roomId == msg.To {
				c.Send(msg)
			}
		case "user":
			if c.userId == msg.To {
				c.Send(msg)
			}
		default:
			if c.sessionId == msg.To {
				c.Send(msg)
			}
		}
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			delete(h.clients, client)
		}
	}
}

func (h *Hub) WsHandler(conn *fastws.Conn) {
	userId, ok := conn.UserValue("user_id").(string)
	if !ok {
		conn.Close()
		return
	}
	sessionId, ok := conn.UserValue("session_id").(string)
	if !ok {
		conn.Close()
		return
	}

	roomId, ok := conn.UserValue("room_id").(string)
	if !ok {
		conn.Close()
		return
	}

	c := NewClient(conn, h, sessionId, roomId, userId, h.log)
	h.register <- c
	c.Read()
}

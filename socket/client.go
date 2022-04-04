package socket

import (
	"encoding/json"
	"strings"

	"github.com/Mirobidjon/udevs_websocket_service/pkg/logger"
	"github.com/dgrr/fastws"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"data"`
	To      string `json:"to"`
}

type Client struct {
	conn      *fastws.Conn
	hub       *Hub
	sessionId string
	roomId    string
	userId    string
	closed    bool
	log       logger.Logger
}

func NewClient(conn *fastws.Conn, hub *Hub, sessionId, roomId, userId string, log logger.Logger) *Client {
	return &Client{
		conn:      conn,
		sessionId: sessionId,
		roomId:    roomId,
		userId:    userId,
		hub:       hub,
		closed:    false,
		log:       log,
	}
}

func (c *Client) Send(msg Message) int32 {
	if c.closed {
		return 0
	}

	js, err := json.Marshal(msg)
	if err != nil {
		c.log.Error("error marshaling message", logger.Error(err), logger.Any("message", msg))
		return 0
	}

	_, err = c.conn.Write(js)
	if err != nil {
		c.log.Error("error writing message or connection closed", logger.Error(err), logger.Any("message", msg))
		c.Close()
	}

	return 1
}

func (c *Client) Read() {
	c.log.Info("reading from client", logger.String("session_id", c.sessionId), logger.String("room_id", c.roomId), logger.String("user_id", c.userId))
	var msg []byte
	var err error

	c.Send(Message{
		Type:    "ping",
		Content: "ping",
		To:      c.userId,
	})

	for {
		if c.closed {
			return
		}

		_, msg, err = c.conn.ReadMessage(msg[:0])
		if err != nil {
			if err != fastws.EOF {
				c.log.Error("reading time: error reading message", logger.Error(err))
			}
			c.Close()
			break
		}

		var msgObj Message
		err = json.Unmarshal(msg, &msgObj)
		if err != nil {
			c.log.Error("reading time: error unmarshaling message", logger.Error(err), logger.Any("message", msg))
			continue
		}

		if strings.ToLower(msgObj.Type) == "ping" {
			c.Send(Message{"user", "pong", c.userId})
		} else {
			c.hub.Send(msgObj)
		}
	}

	c.log.Info("Closed connection user : ", logger.Any("user_id", c.userId), logger.Any("session_id", c.sessionId))
}

func (c *Client) Close() {
	c.hub.unregister <- c
	c.closed = true
	c.conn.Close()
}

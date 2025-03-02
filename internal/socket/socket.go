package socket

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//TODO
		return true
	},
}

type client struct {
	conn *websocket.Conn
	send chan []byte
}

const (
	//writeWait      = 10 * time.Second
	pongWait = 20 * time.Second
	//pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		zap.L().Error("failed to upgrade to websocket", zap.Error(err))
		return
	}

	client := &client{conn: conn}

	go client.readServe()
}

func (c *client) readServe() {
	c.conn.SetReadLimit(maxMessageSize)
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		zap.L().Error("failed to set read deadline", zap.Error(err))
	}
	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	go func() {
		for {
			time.Sleep(5 * time.Second)
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				zap.L().Error("failed to write ping", zap.Error(err))
				return
			}
		}
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				zap.L().Error("socket something went wrong", zap.Error(err))
			}
			break
		}
		zap.L().Info("Testing reading message from socket", zap.ByteString("message", message))
	}
}

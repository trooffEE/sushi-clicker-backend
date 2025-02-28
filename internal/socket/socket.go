package socket

import (
	"fmt"
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
	pongWait = 60 * time.Second
	//pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

func ServeWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
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
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				zap.L().Error("socket something went wrong", zap.Error(err))
			}
			break
		}
		c.send <- message
	}
}

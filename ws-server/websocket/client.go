package ws

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	t "github.com/pradeep-selva/arcadia-typerace/ws-server/pkg/types"
	"github.com/pradeep-selva/arcadia-typerace/ws-server/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r * http.Request) bool {
		return true
	},
}

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingWait = (9 * pongWait)/10
	maxMessageSize = 512
)

type connection t.Connection
type subscription t.Subscription

func (s subscription) readPump() {
	c := s.Conn

	defer func () {
		h.UnRegister <- t.Subscription(s)
		c.Ws.Close()
	}()

	c.Ws.SetReadLimit(maxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(pongWait))
	c.Ws.SetPongHandler(func (string) error {
		c.Ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := c.Ws.ReadMessage()
		if err != nil {
			utils.LogError(err.Error())
		}

		m := t.Message{
			Data: msg,
			Room: s.Room,
		}
		h.Broadcast <- m
	}
}
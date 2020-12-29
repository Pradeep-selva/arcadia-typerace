package ws

import (
	"time"

	"github.com/gorilla/websocket"
	t "github.com/pradeep-selva/arcadia-typerace/ws-server/pkg/types"
	"github.com/pradeep-selva/arcadia-typerace/ws-server/utils"
)

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
		H.UnRegister <- t.Subscription(s)
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
		H.Broadcast <- m
	}
}

func (c *connection) write(mType int, payload []byte) error {
	c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.Ws.WriteMessage(mType, payload)
}

func (s *subscription) writePump() {
	conn := s.Conn
	var c connection = connection(*conn)

	ticker := time.NewTicker(pingWait)

	defer func(){
		ticker.Stop()
		c.Ws.Close()
	}()

	for {
		select {
		case message, ok := <- c.Send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				utils.LogError("Websocket closed")
				return
			}

			if err := c.write(websocket.TextMessage, message); 
			err != nil {
				utils.LogError("An error occurred while sending message")
				return
			}

		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{});
			err != nil {
				utils.LogError("An error occured while sending ping message")
				return
			}
		}
	}
}
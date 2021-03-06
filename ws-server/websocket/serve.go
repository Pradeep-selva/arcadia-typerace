package ws

import (
	"net/http"

	t "github.com/pradeep-selva/arcadia-typerace/ws-server/pkg/types"
	"github.com/pradeep-selva/arcadia-typerace/ws-server/utils"
)

var upgrader = utils.GetUpgrader()

func ServeWs(w http.ResponseWriter, r *http.Request, roomId string, userName string) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.LogError("Error occurred while connecting to room: " + roomId)
		return
	}
	userCount := len(H.Rooms[roomId])


	c := &t.Connection{
		Ws:   socket,
		Send: make(chan []byte, 256),
	}
	s := t.Subscription{
		Conn: c,
		Room: roomId,
		UserName: userName,
	}

	H.Register <- s

	if userCount == 1 {
		message := t.JoinMessage{
			UserName: userName,
		}

		m := t.Message{
			Data: message.ConvertToBytes(),
			Room: s.Room,
		}
		H.Broadcast <- m
	}

	_s := subscription(s)
	go _s.readPump()
	go _s.writePump()
}

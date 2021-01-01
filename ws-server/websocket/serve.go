package ws

import (
	"bytes"
	"encoding/json"
	"net/http"

	t "github.com/pradeep-selva/arcadia-typerace/ws-server/pkg/types"
	"github.com/pradeep-selva/arcadia-typerace/ws-server/utils"
)

var upgrader = utils.GetUpgrader()

func ServeWs(w http.ResponseWriter, r *http.Request, roomId string, userName string) {
	utils.LogSuccess("ROOM --> " + roomId)

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
		message.Fill()
		msgBytes := new(bytes.Buffer)
		json.NewEncoder(msgBytes).Encode(message)

		m := t.Message{
			Data: msgBytes.Bytes(),
			Room: s.Room,
		}
		H.Broadcast <- m
	}

	_s := subscription(s)
	utils.LogSuccess(_s.Room)
	go _s.readPump()
	go _s.writePump()
}

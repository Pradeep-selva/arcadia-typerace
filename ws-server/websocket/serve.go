package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
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

	socket.WriteMessage(websocket.TextMessage, []byte(userName+"joined"))

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

	_s := subscription(s)
	utils.LogSuccess(_s.Room)
	go _s.readPump()
	go _s.writePump()
}

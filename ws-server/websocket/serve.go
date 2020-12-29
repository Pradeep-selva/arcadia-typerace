package ws

import (
	"net/http"

	t "github.com/pradeep-selva/arcadia-typerace/ws-server/pkg/types"
	"github.com/pradeep-selva/arcadia-typerace/ws-server/utils"
)

func serveWs(w http.ResponseWriter, r *http.Request, roomId string) {
	utils.LogSuccess("ROOM --> " + roomId)

	socket, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		utils.LogError("Error occurred while connecting to room: "+roomId)
		return
	}

	c := &t.Connection{
		Ws: socket,
		Send: make(chan []byte, 256),
	}
	s := t.Subscription{
		Conn: c,
		Room: roomId,
	}

	h.Register <- s

	_s := subscription(s)
	go _s.readPump()
	go _s.writePump()
}
package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/pradeep-selva/arcadia-typerace/ws-server/utils"
	ws "github.com/pradeep-selva/arcadia-typerace/ws-server/websocket"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("--Welcome to Arcadia TypeRacer API--"))
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	roomId := strings.Split(r.URL.Path, "/")[2]
	userCount := len(ws.H.Rooms[roomId])

	utils.LogSuccess(
		"Number of users in room "+
		roomId+": "+
		strconv.Itoa(userCount),
	)

	if(userCount <= 1) {
		ws.ServeWs(w,r,roomId)
	}
}
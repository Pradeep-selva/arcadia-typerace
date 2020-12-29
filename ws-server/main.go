package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	utils "github.com/pradeep-selva/arcadia-typerace/ws-server/utils"
	ws "github.com/pradeep-selva/arcadia-typerace/ws-server/websocket"
)

func socketHandler(w http.ResponseWriter, r *http.Request) {
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

func main(){
	go ws.H.Run()

	var PORT string = ":5500"

	http.Handle("/", http.FileServer(http.Dir("../client")))
	http.HandleFunc("/ws/", socketHandler)

	utils.LogSuccess("Listening on port"+PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
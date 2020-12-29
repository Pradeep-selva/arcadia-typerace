package main

import (
	"log"
	"net/http"
	"strings"

	utils "github.com/pradeep-selva/arcadia-typerace/ws-server/utils"
	ws "github.com/pradeep-selva/arcadia-typerace/ws-server/websocket"
)


func main(){
	go ws.H.Run()

	var PORT string = ":5500"

	http.Handle("/", http.FileServer(http.Dir("../client-ui")))
	http.HandleFunc("/ws/race/", func(w http.ResponseWriter, r *http.Request){
		roomId := strings.Split(r.URL.Path, "/")[3]
		ws.ServeWs(w,r,roomId)
	})

	utils.LogSuccess("Listening on port"+PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
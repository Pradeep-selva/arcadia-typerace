package main

import (
	"log"
	"net/http"

	"github.com/pradeep-selva/arcadia-typerace/ws-server/controllers"
	utils "github.com/pradeep-selva/arcadia-typerace/ws-server/utils"
	ws "github.com/pradeep-selva/arcadia-typerace/ws-server/websocket"
)

func main() {
	go ws.H.Run()

	var PORT string = ":5500"

	http.HandleFunc(utils.GetApiPath("/"), controllers.HomeHandler)
	http.HandleFunc(utils.GetApiPath("/ws/"), controllers.SocketHandler)
	http.HandleFunc(utils.GetApiPath("/validate/"), controllers.JoinRoomValidationHandler)
	http.Handle("/test/", http.StripPrefix("/test/",
		http.FileServer(http.Dir("../client"))),
	)

	utils.LogSuccess("Listening on port" + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

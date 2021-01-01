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

	//home
	http.HandleFunc("/", controllers.HomeHandler)
	//socket
	http.Handle(utils.GetApiPath("/ws/"), 
		utils.Middleware(controllers.SocketHandler))
	//validation
	http.Handle(utils.GetApiPath("/validate/"), 
		utils.Middleware(controllers.RoomValidationHandler))
	//test
	http.Handle("/test/", http.StripPrefix("/test/",
		http.FileServer(http.Dir("../client"))),
	)

	utils.LogSuccess("Listening on port" + PORT)
	log.Println(http.ListenAndServe(PORT, nil))
}

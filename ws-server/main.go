package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pradeep-selva/arcadia-typerace/ws-server/controllers"
	utils "github.com/pradeep-selva/arcadia-typerace/ws-server/utils"
	ws "github.com/pradeep-selva/arcadia-typerace/ws-server/websocket"
)

func main() {
	go ws.H.Run()

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
	http.HandleFunc("/.well-known/pki-validation/869F50E3127FCA731973303536208C2D.txt", controllers.VerifyDNS)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("An error occurred while reading .env")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "80"
	}
	utils.LogSuccess("Listening on port :" + PORT)
	log.Println(http.ListenAndServe(":"+PORT, nil))
}

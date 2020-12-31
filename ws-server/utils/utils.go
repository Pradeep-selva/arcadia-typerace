package utils

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func LogSuccess(message string) {
	log.Println("[SERVER]: "+message)
}

func LogError(message string){
	log.Println("[ERROR]: "+message)
}

func LogReceived(message string) {
	log.Println("[MESSAGE RECEIVED]: "+message)
}

func GetUpgrader() websocket.Upgrader {
	return  websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r * http.Request) bool {
			return true
		},
	}
}

func GetApiPath(path string) string {
	return "/api"+path
}
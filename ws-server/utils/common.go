package utils

import (
	"log"
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
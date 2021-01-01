package types

import "github.com/gorilla/websocket"

type Message struct {
	Data []byte
	Room string
}

type Connection struct {
	Ws *websocket.Conn
	Send chan []byte
}

type Subscription struct {
	Conn *Connection
	Room string
	UserName string
}

type Hub struct {
	Rooms map[string]map[*Connection]bool
	Broadcast chan Message
	Register chan Subscription
	UnRegister chan Subscription
}

type ValidationResponse struct {
	Data string `json:"data"`
	Ok bool `json:"ok"`
}

type TypingMessage struct {
	Event string `json:"event"`
	UserName string `json:"userName"`
	Data string `json:"data"`
}

type JoinMessage struct {
	Event string `json:"event"`
	UserName string `json:"userName"`
}

type WonMessage struct {
	Event string `json:"event"`
	UserName string `json:"userName"`
}




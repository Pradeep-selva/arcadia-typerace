package types

import "github.com/gorilla/websocket"

type message struct {
	data []byte
	race string
}

type connection struct {
	ws *websocket.Conn
	send chan []byte
}

type subscription struct {
	conn *connection
	race string
}

type hub struct {
	races map[string]map[*connection]bool
	broadcast chan message
	register chan subscription
	unregister chan subscription
}




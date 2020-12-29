package ws

import (
	t "github.com/pradeep-selva/arcadia-typerace/ws-server/pkg/types"
)

type hub t.Hub

var H = hub {
	Rooms: make(map[string]map[*t.Connection]bool),
	Broadcast: make(chan t.Message),
	Register: make(chan t.Subscription),
	UnRegister: make(chan t.Subscription),
}

func (H *hub) Run() {
	for {
		select {
		case s := <-H.Register:
			if connection := H.Rooms[s.Room]; connection == nil {
				connection = make(map[*t.Connection]bool)
				H.Rooms[s.Room] = connection
			}
			H.Rooms[s.Room][s.Conn] = true
		case s := <-H.UnRegister:
			connections := H.Rooms[s.Room]
			if connections != nil {
				if _, ok := connections[s.Conn]; ok {
					delete(connections, s.Conn)
					close(s.Conn.Send)

					if len(connections) == 0 {
						delete(H.Rooms, s.Room)
					}
				}
			}
		case m := <-H.Broadcast:
			connections := H.Rooms[m.Room]
			for c := range connections {
				select {
				case c.Send <- m.Data:
				default:
					close(c.Send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(H.Rooms, m.Room)
					}
				}
			}
		}
	}
}
package ws

import t "github.com/pradeep-selva/arcadia-typerace/ws-server/pkg/types"

type hub t.Hub

var h = hub {
	Rooms: make(map[string]map[*t.Connection]bool),
	Broadcast: make(chan t.Message),
	Register: make(chan t.Subscription),
	UnRegister: make(chan t.Subscription),
}

func (h *hub) Run() {
	for {
		select {
		case s := <-h.Register:
			if connection := h.Rooms[s.Room]; connection == nil {
				connection = make(map[*t.Connection]bool)
				h.Rooms[s.Room] = connection
			}
			h.Rooms[s.Room][s.Conn] = true
		case s := <-h.UnRegister:
			connections := h.Rooms[s.Room]
			if connections != nil {
				if _, ok := connections[s.Conn]; ok {
					delete(connections, s.Conn)
					close(s.Conn.Send)

					if len(connections) == 0 {
						delete(h.Rooms, s.Room)
					}
				}
			}
		case m := <-h.Broadcast:
			connections := h.Rooms[m.Room]
			for c := range connections {
				select {
				case c.Send <- m.Data:
				default:
					close(c.Send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.Rooms, m.Room)
					}
				}
			}
		}
	}
}
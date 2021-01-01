package types

import "github.com/pradeep-selva/arcadia-typerace/ws-server/configs"

func (t *TypingMessage) Fill() {
	if t.Event == "" {
		t.Event = configs.Events.TYPING
	}
}

func (t *JoinMessage) Fill() {
	if t.Event == "" {
		t.Event = configs.Events.JOIN
	}
}

func (t *WonMessage) Fill() {
	if t.Event == "" {
		t.Event = configs.Events.WON
	}
}
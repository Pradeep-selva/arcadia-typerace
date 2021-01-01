package types

func (t *TypingMessage) Fill() {
	if t.Event == "" {
		t.Event = "typing"
	}
}

func (t *JoinMessage) Fill() {
	if t.Event == "" {
		t.Event = "join"
	}
}

func (t *WonMessage) Fill() {
	if t.Event == "" {
		t.Event = "won"
	}
}
package types

func (t *TypingMessage) fill() {
	if t.Event == "" {
		t.Event = "typing"
	}
}

func (t *JoinMessage) fill() {
	if t.Event == "" {
		t.Event = "join"
	}
}

func (t *WonMessage) fill() {
	if t.Event == "" {
		t.Event = "won"
	}
}
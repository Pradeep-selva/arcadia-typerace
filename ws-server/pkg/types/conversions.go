package types

import (
	"bytes"
	"encoding/json"
)

func (t *TypingMessage) ConvertToBytes() []byte {
	t.Fill()
	bytePayload := new(bytes.Buffer)

	json.NewEncoder(bytePayload).Encode(t)
	return bytePayload.Bytes()
}

func (j *JoinMessage) ConvertToBytes() []byte {
	j.Fill()
	bytePayload := new(bytes.Buffer)

	json.NewEncoder(bytePayload).Encode(j)
	return bytePayload.Bytes()
}

func (w *WonMessage) ConvertToBytes() []byte {
	w.Fill()
	bytePayload := new(bytes.Buffer)

	json.NewEncoder(bytePayload).Encode(w)
	return bytePayload.Bytes()
}
package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pradeep-selva/arcadia-typerace/ws-server/configs"
	t "github.com/pradeep-selva/arcadia-typerace/ws-server/pkg/types"
	"github.com/pradeep-selva/arcadia-typerace/ws-server/utils"
	ws "github.com/pradeep-selva/arcadia-typerace/ws-server/websocket"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("--Welcome to Arcadia TypeRacer API--"))
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	roomId := strings.Split(r.URL.Path, "/")[4]
	userName := strings.Split(r.URL.Path, "/")[5]
	userCount := len(ws.H.Rooms[roomId])

	if userCount <= 1 {
		ws.ServeWs(w,r,roomId, userName)
	}
}

func RoomValidationHandler(w http.ResponseWriter, r *http.Request) {
	valType := strings.Split(r.URL.Path, "/")[3]
	roomId := strings.Split(r.URL.Path, "/")[4]
	userCount := len(ws.H.Rooms[roomId])

	w.Header().Set("Content-Type", "application/json")

	err := configs.GetIfRequiredUsersPresent(valType, userCount)
	if err != nil {
		utils.LogError("Invalid operation blocked.")
		json.NewEncoder(w).Encode(t.ValidationResponse{
			Data: err.Error(), 
			Ok: false,
		})
		return
	}

	json.NewEncoder(w).Encode(t.ValidationResponse{
		Data: "Validation passed.",
		Ok: true,
	})
}
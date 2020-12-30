package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	roomId := strings.Split(r.URL.Path, "/")[2]
	userCount := len(ws.H.Rooms[roomId])

	utils.LogSuccess(
		"Number of users in room "+
		roomId+": "+
		strconv.Itoa(userCount),
	)

	if(userCount <= 1) {
		ws.ServeWs(w,r,roomId)
	}
}

func JoinRoomValidationHandler(w http.ResponseWriter, r *http.Request) {
	valType := strings.Split(r.URL.Path, "/")[2]
	roomId := strings.Split(r.URL.Path, "/")[3]
	userCount := len(ws.H.Rooms[roomId])

	w.Header().Set("Content-Type", "application/json")

	err := configs.GetIfRequiredUsersPresent(valType, userCount)
	if err != nil  {
		utils.LogError("Invalid operation blocked.")
		json.NewEncoder(w).Encode(t.ValidationResponse{
			Data: err.Error(), 
			Ok: false,
		})
		return
	}

	utils.LogSuccess("Validation passed. -- >"+valType+"< -- >"+roomId+"< --")
	json.NewEncoder(w).Encode(t.ValidationResponse{
		Data: "Validation passed.",
		Ok: true,
	})
}
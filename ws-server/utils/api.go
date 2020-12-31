package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	t "github.com/pradeep-selva/arcadia-typerace/ws-server/pkg/types"
)

func CORSMiddleware(next http.HandlerFunc) http.Handler {
    return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, x-api-key, Authorization, accept, origin, Cache-Control, X-Requested-With")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

			next.ServeHTTP(w,r)
    },
	)
}

func Middleware(next http.HandlerFunc) http.Handler {
	return CORSMiddleware(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			err := godotenv.Load(".env")
			if err != nil {
				LogError("An error occured while initializing .env")
				json.NewEncoder(w).Encode(t.ValidationResponse{
					Data: "An error occurred",
					Ok: false,
				})
				return
			}

			if os.Getenv("X_API_KEY") == r.Header.Get("x-api-key") {
				next.ServeHTTP(w,r)
			} else {
				LogError("Unauthorized user attempted to access")
				json.NewEncoder(w).Encode(t.ValidationResponse{
					Data: "You are not authorized to be here",
					Ok: false,
				})
			}
	},
))
}

func GetUpgrader() websocket.Upgrader {
	return  websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r * http.Request) bool {
			return true
		},
	}
}

func GetApiPath(path string) string {
	return "/api"+path
}

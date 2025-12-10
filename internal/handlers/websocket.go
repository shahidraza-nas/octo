// Package handlers provides HTTP handlers for WebSocket upgrades and related endpoints.
package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/shahid/octo-workspace/internal/client"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWs(hub client.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := client.NewClient(hub, conn)
	hub.Register() <- client

	go client.WritePump()
	go client.ReadPump()
}

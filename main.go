package main

import (
	"log"
	"net/http"

	"github.com/shahid/octo-workspace/internal/handlers"
	"github.com/shahid/octo-workspace/internal/hub"
)

func main() {

	h := hub.NewHub()
	go h.Run()

	http.HandleFunc("/", handlers.ServeHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWs(h, w, r)
	})

	addr := ":8080"
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

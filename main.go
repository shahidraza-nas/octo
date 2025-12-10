package main

import (
"log"
"net/http"

"github.com/shahid/octo-workspace/internal/handlers"
"github.com/shahid/octo-workspace/internal/hub"
)

func main() {
// Create a new hub
h := hub.NewHub()
go h.Run()

// Setup routes
http.HandleFunc("/", handlers.ServeHome)
http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
handlers.ServeWs(h, w, r)
})

// Start server
addr := ":8080"
log.Printf("Server starting on %s", addr)
if err := http.ListenAndServe(addr, nil); err != nil {
log.Fatal("ListenAndServe: ", err)
}
}

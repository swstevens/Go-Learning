package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", webSocket)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func webSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WebSocket Endpoint")
	// Upgrade the HTTP connection to a WebSocket connection

	// Handle WebSocket communication
	// Send and receive messages
	// Handle disconnections
}

func main() {
	fmt.Println("Go Webscokets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

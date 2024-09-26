package main

import (
	"fmt"
	"log"
	"math/rand"
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

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		runes := []rune(string(p))
		zeroIndices := []int{}
		for i, r := range runes {
			if r == '0' {
				zeroIndices = append(zeroIndices, i)
			}
		}

		randomIndex := zeroIndices[rand.Intn(len(zeroIndices))]
		runes[randomIndex] = '2'

		if err := conn.WriteMessage(messageType, []byte(string(runes))); err != nil {
			log.Println(err)
			return
		}
	}
}

func webSocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer ws.Close()
	log.Println("Client Successfully Connected")

	reader(ws)
}

func main() {
	fmt.Println("Go Webscokets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

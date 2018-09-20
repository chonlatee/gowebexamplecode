package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("connection error", err)
	}
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read Message error", err)
			return
		}

		msgBack := []byte(string(msg) + " :send from server")

		fmt.Printf("%s send: %s\n", conn.RemoteAddr(), string(msg))

		if err = conn.WriteMessage(msgType, msgBack); err != nil {
			log.Println("Write message error", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websocket.html")
	})

	http.ListenAndServe(":8888", nil)
}

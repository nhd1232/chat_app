package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func ws_handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Клиент подключился")
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	msg := string(p)
	log.Println(msg)
	defer conn.Close()
}

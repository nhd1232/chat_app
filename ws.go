package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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
	defer conn.Close()
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	msg := Message{}
	if err := json.Unmarshal(p, &msg); err != nil {
		log.Println(err)
		return
	}
	now := time.Now()
	msg.Date = now.Format("02.01.2006")
	msg.Time = now.Format("15:04")
	add_message(&messages, msg)
	log.Println(msg)
	b, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}
	if err := conn.WriteMessage(messageType, b); err != nil {
		log.Println(err)
		return
	}
}

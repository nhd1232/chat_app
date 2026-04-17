package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", chat_handler)
	http.HandleFunc("/send", send_handler)
	http.HandleFunc("/ws", ws_handler)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

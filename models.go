package main

import "time"

type Message struct {
	Author string
	Time   time.Time
	Text   string
}

type PageData struct {
	Title    string
	Messages []Message
}

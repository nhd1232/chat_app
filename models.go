package main

type Message struct {
	Author string `json:"author"`
	Time   string `json:"time"`
	Date   string `json:"date"`
	Text   string `json:"text"`
}

type PageData struct {
	Title    string
	Messages []Message
}

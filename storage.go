package main

var messages []Message

func add_message(messages *[]Message, m Message) {
	*messages = append(*messages, m)
}

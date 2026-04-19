package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func chat_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("./templates/chat.html")
	if err != nil {
		log.Printf("Template parse error: %v", err)
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:    "Мой чат",
		Messages: messages,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Template execute error: %v", err)
		return
	}
}

func send_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	author := r.FormValue("username")
	text := r.FormValue("message")
	author = strings.TrimSpace(author)
	text = strings.TrimSpace(text)

	if author == "" || text == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	msg := Message{
		Author: author,
		Time:   "",
		Text:   text,
	}

	add_message(&messages, msg)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

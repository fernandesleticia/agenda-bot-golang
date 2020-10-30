package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

// The conversation that message belongs
type Chat struct {
	Id int `json:"id"`
}

//received when an user interacts with the bot
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

//decode an incoming request
func parseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming message %s", err.Error())
		return nil, err
	}
	return &update, nil
}

func HandleMessage(w http.ResponseWriter, r *http.Request) {

	// parse incoming message
	var update, err = parseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}
}

func (u Update) String() string {
	return fmt.Sprintf("(update id: %d, message: %s)", u.UpdateId, u.Message)
}

func (m Message) String() string {
	return fmt.Sprintf("(text: %s, chat: %s)", m.Text, m.Chat)
}

func (c Chat) String() string {
	return fmt.Sprintf("(id: %d)", c.Id)
}

func main() {
	fmt.Println("Hello, World!")
}

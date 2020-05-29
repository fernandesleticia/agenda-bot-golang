package handler

import "fmt"

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

// The conversation that message belongs
type Chat struct {
	Id int `json:"id"`
}

// Telegram object received when an user interacts with the bot
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

func (u Update) String() string {
	return fmt.Sprintf("(update id: %d, message: %s)", u.UpdateId, u.Message)
}

func (m Message) String() string {
	return fmt.Sprintf("(text: %s, chat: %s)", m.Text, m.Chat, m.Audio)
}

func (c Chat) String() string {
	return fmt.Sprintf("(id: %d)", c.Id)
}

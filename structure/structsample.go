package main

import (
	"errors"
	"fmt"
	"time"
)

type Message struct {
	Data      []byte
	MimeType  string
	Timestamp time.Time
}

type TwitterSource struct {
	Username string
}

type SkypeSource struct {
	Login         string
	MSBackdoorKey string
}

// Finder represents any source for words lookup.
type Finder interface {
	Find(word string) ([]Message, error)
}

func (s TwitterSource) Find(word string) ([]Message, error) {
	return s.searchAPICall(s.Username, word)
}

func (s SkypeSource) Find(word string) ([]Message, error) {
	return s.searchSkypeServers(s.MSBackdoorKey, s.Login, word)
}

type Sources []Finder

func (s Sources) SearchWords(word string) []Message {
	var messages []Message
	for _, source := range s {
		msgs, err := source.Find(word)
		if err != nil {
			fmt.Println("WARNING:", err)
			continue
		}
		messages = append(messages, msgs...)
	}

	return messages
}

var (
	sources = Sources{
		TwitterSource{
			Username: "@rickhickey",
		},
		SkypeSource{
			Login:         "rich.hickey",
			MSBackdoorKey: "12345",
		},
	}

	person = Person{
		FullName: "Rick Hickey",
		Sources:  sources,
	}
)

type Person struct {
	FullName string
	Sources
}

func main() {
	msgs := person.SearchWords("если бы бабушка")
	fmt.Println(msgs)
}

func (s TwitterSource) searchAPICall(username, word string) ([]Message, error) {
	return []Message{
		Message{
			Data:      ([]byte)("Remember, remember, the fifth of November, если бы бабушка..."),
			MimeType:  "text/plain",
			Timestamp: time.Now(),
		},
	}, nil
}

func (s SkypeSource) searchSkypeServers(key, login, word string) ([]Message, error) {
	return []Message{}, errors.New("NSA can't read your skype messages ;)")
}

func (m Message) String() string {
	return string(m.Data) + " @ " + m.Timestamp.Format(time.RFC822)
}

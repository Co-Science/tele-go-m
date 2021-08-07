package main

import (
	"encoding/json"
	"log"
	"net/http"
	"errors"
	"fmt"
	"time"
)

/////  Struct for parsing incoming messages  ////

type Update struct {
	UpdateId int `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat `json:"chat"`
}

type Chat struct {
	Id int `json:"id"`
}

/////  Struct for parsing incoming messages  ////

////  Variables Required  /////

const Token string = "1815331593:AAGM_U2Dw5KxQo3rjTIajSesZvfcj9r_iYw" // invalid bot token, CHANGE!!
const baseUrl string = "https://api.telegram.org/bot"

////  Variables Required  /////


func tt() string {

	var response string = "TT is up and ready on a "

	if string(time.Now().Weekday()) == "Saturday" {
		response += "Monday"
	}
	return response
}

// Function to decode the incoming json and extract chat text from it
func parseConversation(r *http.Request) (*Update, error) {
	var update Update

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("Decoding error, %s", err.Error())
		return nil, err
	}

	if update.UpdateId == 0 {
		log.Printf("Invalid Update ID!")
		return nil, errors.New("Invalid Update ID")
	}

	return &update, nil
}

// function to receive and handle webhooks, calls parseConversation
// and directs the reply
func handleWebhooks(w http.ResponseWriter, r *http.Request) {
	update, err := parseConversation(r)
	if err != nil {
		log.Printf("Error in parsing update, %s", err.Error())
		return
	}

	if update.Message.Text == "/tt" {
		fmt.Println(tt())
	}
}

func main() {
	fmt.Println("Listening on port 8080 ....")
	http.ListenAndServe(":8080", http.HandlerFunc(handleWebhooks))
}
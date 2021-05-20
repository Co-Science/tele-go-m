package main

import (
	"fmt"
	"log"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type Response struct {
	update_id int 'json:"update_id"'
	message []Message 'json:"message"'
}

type Message struct {
	message_id int 'json:"message_id"'
	
}

type From struct {
	username string 'json:"username"'
}

func parseIncomingRequest(r *http.Request) (*Update, error) {

	var response Response

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		log.Printf("Error in decoding messages")

		return nil, err
	}

	if response.update_id == 0 {
		log.Printf("invalid update_id, got update id = 0")
		return nil, errors.New("Server Error")
	}

	return &update, nil
}

func main() {


}
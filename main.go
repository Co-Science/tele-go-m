package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Response Response `json:"result"`
}

type Response struct {
	update_id int `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	message_id int `json:"message_id"`
	Chat Chat `json:"chat"`
	text string `json:"text"`
}

type Chat struct {
	username string `json:"username"`
}

func parseIncomingRequest(r *http.Response) (*Result, error) {

	var result Result

	fmt.Println(r.Body)

	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		log.Printf("Error in decoding messages")

		return nil, err
	}

	if result.Response.update_id == 0 {
		log.Printf("invalid update_id, got update id = 0")
		return nil, errors.New("SERVER ERROR")
	}

	return &result, nil
}

func main() {

	httpreq, err := http.Get("https://api.telegram.org/bot1815331593:AAGM_U2Dw5KxQo3rjTIajSesZvfcj9r_iYw/getUpdates?limit=1")
	if err != nil {
		log.Printf("Error in rerieving request")
	}
	
	parsedData, err := parseIncomingRequest(httpreq)
	if err != nil {
		fmt.Println("Error in parsing retreived data!")
	}

	fmt.Println(parsedData.Response.Message.text)
}
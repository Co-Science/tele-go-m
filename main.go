package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
)

func sendMessages(message string) (error){

	url := fmt.Sprintf("https://api.telegram.org/bot<token>/sendMessage?chat_id=856051391&text=%s", message)

	httpSend, err := http.Get(url)
	if err != nil {
		return errors.New("request failed")
	}

	httpSendBody, _ := ioutil.ReadAll(httpSend.Body)

	success, _ := jsonparser.GetString(httpSendBody, "ok")
	if success != "true" {
		return errors.New("could not send messages")
	} else {
		return nil
	}
}

func parseIncomingRequest(httpResp *http.Response) (string, error) {

	bodyBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		panic(err)
	}

	body, err := jsonparser.GetString(bodyBytes, "result", "[0]", "message", "text")
	if err != nil {
		log.Fatal("Error in parsing data")
	}

	return body, nil
}

func main() {

	httpreq, err := http.Get("https://api.telegram.org/bot<token>/getUpdates?limit=1")
	if err != nil {
		log.Printf("Error in rerieving request")
	}
	
	parsedData, err := parseIncomingRequest(httpreq)
	if err != nil {
		fmt.Println("Error in parsing retreived data!")
	}

	fmt.Println(parsedData)

	sendMessages(parsedData)
}
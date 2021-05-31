package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/buger/jsonparser"
)

type JsonToGo struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result"`
}
type From struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
type Chat struct {
	ID                          int    `json:"id"`
	Title                       string `json:"title"`
	Type                        string `json:"type"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
}
type Message struct {
	MessageID    int    `json:"message_id"`
	From         From   `json:"from"`
	Chat         Chat   `json:"chat"`
	Date         int    `json:"date"`
	NewChatTitle string `json:"new_chat_title"`
}
type Result struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

func sendMessages(message string, token string, chat_id string) error {

	url := fmt.Sprintf("https://api.telegram.org/bot"+token+"/sendMessage?chat_id="+chat_id+"&text=%s", message)

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

func parseIncomingRequest(httpResp *http.Response) {

	bodyBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		panic(err)
	}

	sb := "[" + string(bodyBytes) + "]"

	fmt.Println(sb)
	bytes := []byte(sb)

	var json_in_go []JsonToGo
	json.Unmarshal(bytes, &json_in_go)

	fmt.Println(json_in_go)
	for _, v := range json_in_go {
		fmt.Printf("\nId = %v \n", v)
		for j, v := range v.Result {
			fmt.Printf("unk= %v,%v \n", j, v)
		}
		fmt.Println()
	}

}

func fileReader(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	secreats := make([]string, 2)
	for i, each_ln := range text {
		secreats[i] = strings.Split(each_ln, "=")[1]
	}
	return secreats
}

func main() {

	secreats := fileReader("env.txt")
	token, chat_id := secreats[0], secreats[1]
	fmt.Println(token, chat_id)

	httpreq, err := http.Get("https://api.telegram.org/bot" + token + "/getUpdates?limit=1")
	if err != nil {
		log.Printf("Error in rerieving request")
	}

	parseIncomingRequest(httpreq)

	// parsedData, err := parseIncomingRequest(httpreq)
	if err != nil {
		fmt.Println("Error in parsing retreived data!")
	}

	// fmt.Println(parsedData)

	sendMessages("hello", token, chat_id)
}

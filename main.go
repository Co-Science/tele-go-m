package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var token, chat_id string

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func sayHello(chatID string) error {
	res, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", token, chatID, "hello"))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}
	return nil
}

func Handler(res http.ResponseWriter, req *http.Request) {
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}
	if !strings.ContainsAny(strings.ToLower(body.Message.Text), "telegom") {
		return
	}
	if err := sayHello(fmt.Sprint(body.Message.Chat.ID)); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}
	fmt.Println("reply sent")
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	token = os.Getenv("TOKEN")

	fmt.Println(token)
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}

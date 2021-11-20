package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var TOKEN string

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID        int64  `json:"id"`
			Username  string `json:"username"`
			Firstname string `json:"first_name"`
		} `json:"chat"`
	} `json:"message"`
}

func sayHello(chatID, user string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", TOKEN, chatID, "hello "+user)
	res, err := http.Get(url)
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
	err := sayHello(fmt.Sprint(body.Message.Chat.ID), body.Message.Chat.Firstname)
	if err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}
	fmt.Println("reply sent")
	res.Write([]byte("Gorilla!\n"))
}

func main() {

	// // uncomment for local development
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	TOKEN = os.Getenv("TOKEN")

	fmt.Println(TOKEN)

	r := mux.NewRouter()
	r.HandleFunc("/", Handler)

	log.Fatal(http.ListenAndServe(":8000", r))

	// http.ListenAndServe(":8000", http.HandlerFunc(Handler))
}

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
}

// func fileReader(filename string) (err error) {

// 	data, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return err
// 	}
// 	TOKEN = strings.Trim(fmt.Sprint(strings.Split(string(data), "=")[1]), " ")
// 	return nil
// }

func main() {

	// uncomment for local development
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	TOKEN = os.Getenv("TOKEN")

	// err := fileReader(".env")
	// if err != nil {
	// 	fmt.Println("error in reading file", err)
	// 	return
	// }
	fmt.Println(TOKEN)
	http.ListenAndServe(":8080", http.HandlerFunc(Handler))
}

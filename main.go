package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var TOKEN string
var EASTEREGG []string = []string{
	"Did someone say 🙂 'telegom'?",
	"Howdie master...",
	"You have unlocked telegom+",
	"Come join our team... @[github](https://github.com/Co-Science)",
	"You can add more random things @[tele-go-m repo](https://github.com/Co-Science/tele-go-m)",
}

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

// reply to a specific user
func sayHello(chatID, user string) error {
	randomIndex := rand.Intn(len(EASTEREGG))
	url := ""
	if user == "" || randomIndex != 1 {
		url = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", TOKEN, chatID, EASTEREGG[randomIndex])
	} else {
		url = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", TOKEN, chatID, EASTEREGG[randomIndex]+user)
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	} else if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}
	return nil
}

// general reply
func sayCustomHelloWithoutName(chatID, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", TOKEN, chatID, text)
	res, err := http.Get(url)
	if err != nil {
		return err
	} else if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}
	return nil
}

// timetable
func tt() string {
	var response string = "TT is up and ready on a "
	if string(time.Now().Weekday()) == "Saturday" {
		response += "Monday"
	}
	return response
}

// Function to handle the request \o/
func Handler(res http.ResponseWriter, req *http.Request) {
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	// checks if the string contains 'telegom'
	if strings.Contains(strings.ToLower(body.Message.Text), "telegom") {
		err := sayHello(fmt.Sprint(body.Message.Chat.ID), body.Message.Chat.Firstname)
		if err != nil {
			fmt.Println("error in sending reply:", err)
		}
		return
	}

	// other commands
	switch body.Message.Text {
	case "/hello":
		sayCustomHelloWithoutName(fmt.Sprint(body.Message.Chat.ID), "Hello "+body.Message.Chat.Firstname+"!")
	case "/tt":
		sayCustomHelloWithoutName(fmt.Sprint(body.Message.Chat.ID), tt())
	case "/help":
		sayCustomHelloWithoutName(fmt.Sprint(body.Message.Chat.ID), "I can help you with the following commands:/hello, /tt")
	}
	// log.Println(body.Message.Text)
	log.Println("reply sent")
	res.Write([]byte("Gorilla!\n"))
}

func main() {

	// // uncomment for local development
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	TOKEN = os.Getenv("TOKEN")
	PORT := os.Getenv("PORT")

	fmt.Println(TOKEN)

	r := mux.NewRouter()
	r.HandleFunc("/", Handler)

	fmt.Println("use ngrok on http://localhost:" + PORT + " || http://0.0.0.0:" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}

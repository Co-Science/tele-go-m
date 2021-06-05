# tele-go-m
telegram bot for Gophers (the webhook style)

## How to make the best use of tele-go-m

- Well why dont u start with [botfather](https://t.me/botfather) . He is the one to meet if you need any bots around here 🙃

What you already did that ? Huh well i guess that all there is then... 

You now have a real class telegram bot on your hand. Now all you have to do is clone this repo and use it as ur own

```
git clone https://github.com/Co-Science/tele-go-m.git
cd tele-go-m
go run main.go
```

- Thats it we did all the work you need to get started but its not done yet. The program now listens on port 3000 for some request.

- There is no use of localhost:3000 unless it can communicate with the outside world so usse [ngrok](https://ngrok.com/) to do that.

- Now set your bots webhook to that url
```
https://api.telegram.org/bot<your_bot_token>/setWebhook?url=<your_https_url_ngrok_provides>
```

__Congratss__ Your bot is not ready to chat with you. Just type telegom or any letter in it and the bot replies with hello.

> You can edit or modify this as you wish

- Now to remove the webhook url just type 

```
https://api.telegram.org/bot<your_bot_token>/deleteWebhook
```

## Starts what are they?

![stars_in_github](./Resources/img/starts.png)

Starts are the way to determine the popularity of a repo.

Dont forget to ⭐ us if you like what u see 😉

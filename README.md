# tele-go-m
telegram bot in go (using webhook)


<details>
  <summary>Beginner's guide for telegram bot </summary>
<!--START_SECTION:waka-->

## How to make the best use of tele-go-m

- To start with this, meet [botfather](https://t.me/botfather) .

We have added the basic backend code to run your bot, you can clone this repo and use it as you'r own

```
git clone https://github.com/Co-Science/tele-go-m.git
cd tele-go-m
```

But before you begin create a `.env` file to the root directory and add your bot api key for testing:
```
TOKEN=<add your_bot_token here>
```
Now run the code:
```
go run main.go
```

- The program now listens on port 3000 for some request.

- Make use of [ngrok](https://ngrok.com/) to test your bot locally.
```
ngrok http 3000
```

- Now set your bots webhook to that url
```
https://api.telegram.org/bot<your_bot_token>/setWebhook?url=<your_https_url_ngrok_provides>
```
| this is essentially connecting your bot with the telegram server 

__Congratss__ Your bot is now ready to chat with you(locally). Just type telegom or any letter in it and see the bot reply.

- To remove the webhook url just type 

```
https://api.telegram.org/bot<your_bot_token>/deleteWebhook
```

<!--END_SECTION:waka-->
</details>	

Fork, Edit and Deploy it in [heroku](https://devcenter.heroku.com/articles/getting-started-with-go?singlepage=true) or with any cloud provider to use the bot.

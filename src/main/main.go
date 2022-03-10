package main

import (
	"github.com/yaroslavvasilenko/golang-telegram-bot/src/handlers"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

const myToken = "5027633442:AAH3phWo2Bc4fwE58aZDb0QlxEOnBYtt7E0"

func main() {
	pref := tele.Settings{
		Token:  myToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	b.Handle(
		tele.OnText,
		handlers.OnUserMessage,
	)
	b.Start()
}

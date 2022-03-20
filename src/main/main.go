package main

import (
	tele "gopkg.in/telebot.v3"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"wow/src/handlers"
)

func main() {
	myToken := getToken()
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

func getToken() string {
	f, e := ioutil.ReadFile("config.txt")
	if e != nil {
		panic(e)
	}
	return strings.Split(string(f), " ")[2]
}

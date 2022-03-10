package handlers

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"net/url"
	"strings"
)

// ToDo: Rework to accept a *.txt/*.json configuration file
var users = [2]string{"serebrennikov_oleg", "yaroslavvasilenko"}

func OnUserMessage(context tele.Context) error {
	var data []string
	user := context.Message().Sender.Username
	mess := context.Message().Text
	for _, acceptedUser := range users {
		if acceptedUser == user {
			data = strings.Split(mess, " ")
			if !validateData(data) {
				err := context.Send(`Write <address> <pass>`)
				if err != nil {
					log.Println("OnUserMessage ERROR:", err)
					return err
				}
				return nil
			}
			log.Printf("Address: %s, %s", data[0], data[1])
			return nil
		}
	}
	return nil
}

func validateData(a []string) bool {
	if len(a) != 2 {
		return false
	}
	if !isUrl(a[0]) {
		return false
	}
	return true
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	// ToDo: replace with some other, more efficient method (e.g. strings.contains())
	if err != nil || u.Scheme != "https" || len(strings.Split(u.Host, ".")) == 1 {
		return false
	}

	fmt.Println("All ok")
	return true
}

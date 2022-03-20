package handlers

import (
	"fmt"
	down "github.com/yaroslavvasilenko/videodw/src/start"
	"gopkg.in/telebot.v3"
	"log"
	"net/url"
	"strings"
)

// ToDo: Rework to accept a *.txt/*.json configuration file
var users = [2]string{"serebrennikov_oleg", "yaroslavvasilenko"}

func OnUserMessage(context telebot.Context) error {
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
			splittedUrl := strings.Split(data[0], "/")
			lessonId := splittedUrl[len(splittedUrl)-1]
			pass := data[1]
			urlstart := fmt.Sprintf("https://events.webinar.ru/api/eventsessions/%s/record?password=%s&withoutCuts=false", lessonId, pass) //
			down.Start(urlstart)
			context.Send("Загрузка всех фалов по ссылке", lessonId, "завершенна")
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
	if err != nil || u.Scheme != "https" || strings.Contains(u.Host, ".") == false {
		return false
	}

	fmt.Println("All ok")
	return true
}

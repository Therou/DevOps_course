package main

import (
	"io/ioutil"
	"log"
	"time"
	"strings"

	telegram "gopkg.in/tucnak/telebot.v2"
)

func main() {
	byte_token, _ := ioutil.ReadFile("token.txt")

	string_token_with_trailing_new_line := string(byte_token)
	string_token_without_trailing_new_line := strings.TrimSuffix(string_token_with_trailing_new_line, "\n")

	bot, err := telegram.NewBot(telegram.Settings{
		Token: string_token_without_trailing_new_line,
	        Poller: &telegram.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Starting GO bot\nSupported commands:\n/hello\n/Git\n/Tasks\n/TaskN, where N is 1-4\n")
	bot.Handle("/hello", func(m *telegram.Message) {
		bot.Send(m.Sender, "hello, this Go bot is online")
	})
	bot.Handle("/Git", func(m *telegram.Message) {
		bot.Send(m.Sender, "https://github.com/Therou/DevOps_course")
	})
        bot.Handle("/Tasks", func(m *telegram.Message) {
                bot.Send(m.Sender, "1 - ansible\n2 - scipt_first\n3 - script_second\n4 - Go Telegram bot")
	})
        bot.Handle("/Task1", func(m *telegram.Message) {
		bot.Send(m.Sender, "https://github.com/Therou/DevOps_course/tree/main/ansible")
	})
        bot.Handle("/Task2", func(m *telegram.Message) {
		bot.Send(m.Sender, "https://github.com/Therou/DevOps_course/tree/main/script_first")
	})
        bot.Handle("/Task3", func(m *telegram.Message) {
		bot.Send(m.Sender, "https://github.com/Therou/DevOps_course/tree/main/script_second")
	})
        bot.Handle("/Task4", func(m *telegram.Message) {
		bot.Send(m.Sender, "https://github.com/Therou/DevOps_course/tree/main/go_bot")
	})

	bot.Start()
}

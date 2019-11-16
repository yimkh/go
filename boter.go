package main

import (
	"database/sql"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//a bot to do things
func startBot(text string) {
	bot, err := tgbotapi.NewBotAPI("996518758:AAGU29DSWvGSqCO4yrR2pcOzcJGAZ9fn4JM")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	go autoGetMes(bot, text)
	// go inputGetMes(bot)
}

//input condition, you could get kinds of messages
//using name, date and so on
//but it is not used now
func inputGetMes(bot *tgbotapi.BotAPI) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/YQ")
	if err != nil {
		log.Fatalf("sql.Open error: %s\n", err)
	}

	rows, err := db.Query("SELECT * FROM user_info")
	if err != nil {
		log.Fatalf("db.Query error: %s\n", err)
	}

	var id int
	var title string
	var name string
	var path string

	for rows.Next() {
		err = rows.Scan(&id, &title, &name, &path)
		if err != nil {
			log.Fatalf("rows.Scan error: %s\n", err)
		}
	}

	u := tgbotapi.NewUpdate(0)

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		} else if update.Message.Text == "get" {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, title+"\n"+name+"\n"+path+"\n")
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		} else {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "please input “get” in order to getupdate messages")
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

//send update message continuously
//need identifier for the target chat or username firstly
func autoGetMes(bot *tgbotapi.BotAPI, text string) {
	//alse need date
	oldtext := ""

	for istext(text) && oldtext != text {
		msg := tgbotapi.NewMessage(-1001351288036, text)

		bot.Send(msg)

		oldtext = text
	}
}

func istext(text string) bool {
	if text != "" {
		return true
	}
	return false
}

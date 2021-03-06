package main

import (
  "os"
  "log"
  "gopkg.in/telegram-bot-api.v4"
)

func main() {
  args := os.Args[1:]
  bot, err := tgbotapi.NewBotAPI(args[0])

  if err != nil {
    log.Panic(err)
  }

  bot.Debug = true

  log.Printf("Authorized on account %s", bot.Self.UserName)

  u := tgbotapi.NewUpdate(0)
  u.Timeout = 60

  updates, err := bot.GetUpdatesChan(u)

  for update := range updates {
    if update.Message == nil {
      continue
    }

    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
    msg.ReplyToMessageID = update.Message.MessageID
    bot.Send(msg)
  }
}

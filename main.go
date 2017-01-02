package main

import (
	"flag"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
)

func main() {
	logCommand := flag.NewFlagSet("log", flag.ExitOnError)
	token1Flag := logCommand.String("token", "", "Bot token")

	sendCommand := flag.NewFlagSet("send", flag.ExitOnError)
	token2Flag := sendCommand.String("token", "", "Bot token")
	recipientFlag := sendCommand.Int64("chat_id", 0, "Chat ID to send")
	messageFlag := sendCommand.String("message", "", "Text message.")

	if len(os.Args) == 1 {
		fmt.Println("usage: gogram <command> [<args>]")
		fmt.Println("The most commonly commands are: ")
		fmt.Println(" log   Log messages to console")
		fmt.Println(" send  Send messages to group/person")
		return
	}

	switch os.Args[1] {
	case "log":
		logCommand.Parse(os.Args[2:])
	case "send":
		sendCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}

	if logCommand.Parsed() {
		if *token1Flag == "" {
			fmt.Println("Please supply the token using -token option.")
			return
		}
		fmt.Printf("You asked: %q\n", *token1Flag)

		bot, err := tgbotapi.NewBotAPI(*token1Flag)
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
		}
	}

	if sendCommand.Parsed() {
		if *token2Flag == "" {
			fmt.Println("Please supply the token using -token option.")
			return
		}
		if *recipientFlag == 0 {
			fmt.Println("Please supply the recipient using -chat_id option.")
			return
		}

		if *messageFlag == "" {
			fmt.Println("Please supply the message using -message option.")
			return
		}

		fmt.Printf("Your message is sent to %q.\n", *recipientFlag)
		fmt.Printf("Message: %q.\n", *messageFlag)

		bot, err := tgbotapi.NewBotAPI(*token2Flag)
		if err != nil {
			log.Panic(err)
		}

		bot.Debug = true

		log.Printf("Authorized on account %s", bot.Self.UserName)
		msg := tgbotapi.NewMessage(*recipientFlag, *messageFlag)
		bot.Send(msg)
	}
}

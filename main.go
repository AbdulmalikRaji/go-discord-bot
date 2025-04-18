package main

import (
	"fmt"

	"github.com/abdulmalikraji/go-discord-bot/bot"
	"github.com/abdulmalikraji/go-discord-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	bot.Start()

	<-make(chan struct{})
	fmt.Println("Bot started successfully. Press Ctrl+C to exit.")
	return

}

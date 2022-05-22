package main

import (
	"fmt"

	"github.com/jcksnvllxr80/go-tuts/discord-ping-bot/bot"
	"github.com/jcksnvllxr80/go-tuts/discord-ping-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}

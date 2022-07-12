package main

import (
	"github.com/ksusonic/relax-tg-bot/pkg/telegram"
	"log"
)

func main() {
	teleBot := telegram.NewBot(true)
	if err := teleBot.Start(); err != nil {
		log.Fatal(err)
	}
}

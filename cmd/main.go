package main

import (
	"github.com/ksusonic/relax-tg-bot/pkg/telegram"
)

func main() {
	teleBot := telegram.NewBot(true)
	teleBot.Start()
}

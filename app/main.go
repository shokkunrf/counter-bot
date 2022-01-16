package main

import (
	"app/discord"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bot, err := discord.MakeBot()
	if err != nil {
		log.Fatalln(err)
	}

	err = bot.Start()
	if err != nil {
		log.Fatalln(err)
	}
	defer bot.Stop()

	log.Println("--- Start ---")

	// 終了を待機
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)

	select {
	case <-signalChan:
		return
	}
}

package main

import (
	"flag"
	"log"
)

func main() {

	tgClient := telegram.New(tgBotHost, mustToken())

	//t := mustToken()
	// tgClient = telegram.New(token)

	//fetcher = fetcher.New(tgclient)

	//processor = processor.New(tgclient)

	//consumer.Start(fetcher, proccessor)

}

func mustToken() (string, error) {

	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("incorrect token")
	}

}

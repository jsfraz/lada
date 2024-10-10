package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jsfraz/lada/handlers"
)

func main() {
	log.SetPrefix("lada: ")
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds)

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	// Register handleres
	dg.AddHandler(handlers.MessageCreate)

	// Fix to receive message content https://github.com/bwmarrin/discordgo/issues/1270
	dg.Identify.Intents |= discordgo.IntentMessageContent

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatalln(err)
	}
	// Log
	log.Println("Bot started!")
	// Wait here until CTRL-C or other term signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

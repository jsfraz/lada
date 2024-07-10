package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	log.SetPrefix("lada: ")
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds)

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}
	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatalln(err)
	}
	// Log
	log.Println("Láďa is running...")
	// Wait here until CTRL-C or other term signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

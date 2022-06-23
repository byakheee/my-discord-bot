package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/byakheee/my-discord-bot/internal/handlers"
)

func main() {
	var token string

	flag.StringVar(&token, "t", "", "API Token")
	flag.Parse()

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("failed to create discordgo client. error: %s", err.Error())
	}

	// register handler.
	discord.AddHandler(handlers.OnConnect)
	discord.AddHandler(handlers.OnMessageCreate)
	discord.AddHandler(handlers.OnReady)
	discord.AddHandler(handlers.OnVoiceServerUpdate)
	discord.AddHandler(handlers.OnVoiceStateUpdate)

	if err := discord.Open(); err != nil {
		log.Fatalf("failed to open connection. error: %s", err.Error())
	}

	log.Println("Bot is now running.  Press CTRL-C to exit.")

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if err := discord.Close(); err != nil {
		log.Fatalf("failed to close connection. error: %s", err.Error())
	}

	log.Println("Bot successfully shutdown!")
}

package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/byakheee/my-discord-bot/internal/handlers"
)

func main() {
	token := flag.String("t", "", "API Token")
	isDebug := flag.Bool("debug", false, "If turn on debug mode, true")
	flag.Parse()
	initLogger(*isDebug)

	discord, err := discordgo.New("Bot " + *token)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed to create discordgo")
	}

	// Register handler.
	discord.AddHandler(handlers.OnConnect)
	discord.AddHandler(handlers.OnMessageCreate)
	discord.AddHandler(handlers.OnReady)
	discord.AddHandler(handlers.OnVoiceServerUpdate)
	discord.AddHandler(handlers.OnVoiceStateUpdate)

	if err := discord.Open(); err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed to open connection")
	}

	log.Info().Msg("Bot is now running.  Press CTRL-C to exit.")

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if err := discord.Close(); err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed to close connection. error: %s")
	}

	log.Info().Msg("Bot successfully shutdown!")
}

func initLogger(debug bool) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if !debug {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/byakheee/my-discord-bot/internal/extensions/logex"
)

// nolint:gochecknoglobals
var voiceConnection *discordgo.VoiceConnection

func OnMessageCreate(discord *discordgo.Session, message *discordgo.MessageCreate) {
	log.Info().Msg("On message create!")
	logex.DebugAnyStruct(message)

	if message.Author.Bot {
		log.Debug().Msg("Skip message because author is bot.")

		return
	}

	if message.Content == "" {
		log.Debug().Msg("Skip message because content is empty.")

		return
	}

	// メッセージの種類: https://discord.com/developers/docs/resources/channel#message-object-message-types
	if message.Type != discordgo.MessageTypeDefault {
		log.Debug().Msg("Skip message because type is not default.")

		return
	}

	if strings.HasPrefix(message.Content, "!!") {
		handleCommand(discord, message)

		return
	}

	if _, err := discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Echo (%s)", message.Content)); err != nil {
		log.Error().Stack().Err(err).
			Str("ChannelID", message.ChannelID).
			Msg("Failed to send message to channel")
	}
}

func handleCommand(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == "!!away" {
		onCommandAway()

		return
	}

	if message.Content == "!!comehere" {
		onCommandComehere(discord, message.GuildID, message.Author.ID)

		return
	}

	if message.Content == "!!play" {
		onCommandPlay()

		return
	}
}

func onCommandComehere(discord *discordgo.Session, guildID, authorID string) {
	voiceState, err := discord.State.VoiceState(guildID, authorID)
	if err != nil {
		log.Warn().
			Str("GuildID", guildID).
			Str("AuthorID", authorID).
			Msg("Author do not stay in voice channel.")

		return
	}

	vcon, err := discord.ChannelVoiceJoin(voiceState.GuildID, voiceState.ChannelID, false, true)
	if err != nil {
		log.Error().Stack().Err(err).
			Str("vsGuildID", voiceState.GuildID).
			Str("vsUserID", voiceState.UserID).
			Str("vsChannelID", voiceState.ChannelID).
			Msg("Failed to join voice channel.")
	}

	voiceConnection = vcon
}

func onCommandAway() {
	if voiceConnection == nil {
		log.Warn().
			Msg("Bot do not stay in voice channel.")

		return
	}

	if err := voiceConnection.Disconnect(); err != nil {
		log.Error().Stack().Err(err).
			Msg("Failed to disconnect voice connection.")

		return
	}

	voiceConnection.Close()
	voiceConnection = nil

	log.Info().Msg("voice connection is closed.")
}

func onCommandPlay() {
	if voiceConnection == nil {
		log.Warn().
			Msg("Bot do not stay in voice channel.")

		return
	}

	exec, err := os.Executable()
	if err != nil {
		log.Error().Stack().Err(err).
			Msg("Failed to get os.Executable.")

		return
	}

	dgvoice.PlayAudioFile(
		voiceConnection,
		filepath.Join(filepath.Dir(exec), "internal/assets/newtype.mp3"),
		make(chan bool))
}

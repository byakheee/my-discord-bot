package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/byakheee/my-discord-bot/internal/extensions/logex"
)

func OnMessageCreate(discord *discordgo.Session, message *discordgo.MessageCreate) {
	log.Info().Msg("On message create!")
	logex.DebugRawJSON(message)

	if message.Author.Bot {
		return
	}

	if message.Content == "" ||
		!(message.Type == discordgo.MessageTypeDefault ||
			message.Type == discordgo.MessageTypeReply) {
		// メッセージの種類: https://pkg.go.dev/github.com/bwmarrin/discordgo#MessageType
		log.Info().Msgf("Skip message type: %v, content: %s", message.Type, message.Content)

		return
	}

	if _, err := discord.ChannelMessageSend(message.ChannelID, "Hello! I'm bot!"); err != nil {
		log.Error().Stack().Err(err).
			Str("ChannelID", message.ChannelID).
			Msg("Failed to send message to channel")
	}
}

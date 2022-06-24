package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/byakheee/my-discord-bot/internal/extensions/logex"
)

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

	if _, err := discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Echo (%s)", message.Content)); err != nil {
		log.Error().Stack().Err(err).
			Str("ChannelID", message.ChannelID).
			Msg("Failed to send message to channel")
	}
}

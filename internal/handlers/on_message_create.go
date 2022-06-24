package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/byakheee/my-discord-bot/internal/extensions/logex"
	"github.com/byakheee/my-discord-bot/internal/gateways"
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
	targetUserID := ""

	if message.Content == "!!panic" {
		// TODO: panicしたらリカバリしてsessionをcloseする
		// そうじゃないと voice チャンネルに残ったままになる
		panic("user panic!")
	}

	if message.Content == "!!away" {
		// TODO: onCommandAwayとonCommandComehereにしよう
		targetUserID = discord.State.User.ID
	}

	if message.Content == "!!comehere" {
		targetUserID = message.Author.ID
	}

	channelID := gateways.GetStayngVoiceChannel(message.GuildID, targetUserID)
	if channelID == "" {
		log.Warn().
			Str("TargetUserID", targetUserID).
			Msg("target user do not stay in voice channel.")

		return
	}

	if message.Content == "!!away" {
		channelID = ""
	}

	if _, err := discord.ChannelVoiceJoin(message.GuildID, channelID, false, false); err != nil {
		log.Error().Stack().Err(err).
			Str("GuildID", message.GuildID).
			Str("TargetUserID", targetUserID).
			Str("ChannelID", channelID).
			Msg("Failed to join voice channel.")
	}
}

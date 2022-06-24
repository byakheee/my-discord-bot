package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/byakheee/my-discord-bot/internal/extensions/logex"
)

func OnConnect(discord *discordgo.Session, connect *discordgo.Connect) {
	log.Info().Msg("On connect!")
	logex.DebugAnyStruct(connect)
}

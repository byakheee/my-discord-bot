package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/byakheee/my-discord-bot/internal/extensions/logex"
)

func OnReady(discord *discordgo.Session, ready *discordgo.Ready) {
	log.Info().Msg("On ready!")
	logex.DebugAnyStruct(ready)
}

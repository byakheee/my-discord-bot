package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnReady(discord *discordgo.Session, ready *discordgo.Ready) {
	log.Printf("On ready!\n%#v", *ready)
}

package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnConnect(discord *discordgo.Session, connect *discordgo.Connect) {
	log.Printf("On connect!\n%#v", *connect)
}

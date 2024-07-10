package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called every time a new message is created
// on any channel that the authenticated bot has access to.
//
//	@param s
//	@param m
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// Check if message starts with bot mention
	if !strings.HasPrefix(m.Content, s.State.User.Mention()) {
		return
	}
	// Check if the message isn't empty
	content := strings.TrimPrefix(m.Content, s.State.User.Mention())
	// Remove leading spaces
	spaceCount := 0
	for spaceCount < len(content) && content[spaceCount] == ' ' {
		spaceCount++
	}
	content = content[spaceCount:]
	if content == "" {
		return
	}
	log.Println(content)
	// Check if message content is pong
	if strings.ToLower(content) != "ping" {
		return
	}
	// Set user mention
	mention := ""
	if m.GuildID != "" {
		mention = fmt.Sprintf("%s ", m.Author.Mention())
	}
	// Send message
	_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%sPong!", mention))
	if err != nil {
		log.Println(err)
	}
}

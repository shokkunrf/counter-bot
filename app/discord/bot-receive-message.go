package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) receiveMessage(session *discordgo.Session, event *discordgo.MessageCreate) {
	// mentionされたときのみ通す
	me, err := session.User("@me")
	if err != nil {
		log.Println(err)
		return
	}
	if !isMentioned(event.Mentions, me.ID) {
		return
	}
}

func isMentioned(mentions []*discordgo.User, myID string) bool {
	for _, user := range mentions {
		if user.ID == myID {
			return true
		}
	}
	return false
}

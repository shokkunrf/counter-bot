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

	msg, err := b.session.ChannelMessageSendEmbed(event.ChannelID, &discordgo.MessageEmbed{
		Title:       "Counter",
		Description: "Increment :arrow_up:, Decrement :arrow_down:, Reset :zero:",
	})
	if err != nil {
		log.Println(err)
		return
	}
	b.message = msg

	err = b.session.MessageReactionAdd(event.ChannelID, msg.ID, INCREMENT_EMOJI)
	if err != nil {
		log.Println(err)
		return
	}
	err = b.session.MessageReactionAdd(event.ChannelID, msg.ID, DECREMENT_EMOJI)
	if err != nil {
		log.Println(err)
		return
	}
	err = b.session.MessageReactionAdd(event.ChannelID, msg.ID, RESET_EMOJI)
	if err != nil {
		log.Println(err)
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

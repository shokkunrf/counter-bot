package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) receiveReaction(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	me, err := session.User("@me")
	if err != nil {
		log.Println(err)
		return
	}

	if event.UserID == me.ID {
		return
	}
	if event.MessageID != b.message.ID {
		return
	}

	emojiID := event.Emoji.Name

	_, err = b.session.ChannelMessageEditEmbed(event.ChannelID, event.MessageID, &discordgo.MessageEmbed{
		Title:       "Counter",
		Description: "Increment :arrow_up:, Decrement :arrow_down:, Reset :zero:",
		// Fields:      []*discordgo.MessageEmbedField{{}, {}},
	})
	if err != nil {
		log.Println(err)
		return
	}

	err = b.session.MessageReactionRemove(event.ChannelID, event.MessageID, emojiID, event.UserID)
	if err != nil {
		log.Println(err)
		return
	}
}

package discord

import (
	"app/store"
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

	storeClient, _ := store.GetClient()

	// 基本emojiにはIDがない
	emojiID := event.Emoji.ID
	if emojiID == "" {
		emojiID = event.Emoji.Name
	}

	switch emojiID {
	case INCREMENT_EMOJI:
		storeClient.IncrementCount(event.UserID)
	case DECREMENT_EMOJI:
		storeClient.DecrementCount(event.UserID)
	case RESET_EMOJI:
		storeClient.ResetCount(event.UserID)
	}

	counters, _ := storeClient.GetCounters()
	embeddedMessage, err := b.generateEmbeddedMessage(counters)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = b.session.ChannelMessageEditEmbed(event.ChannelID, event.MessageID, &embeddedMessage)
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

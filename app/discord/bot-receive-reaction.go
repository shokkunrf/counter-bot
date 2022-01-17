package discord

import (
	"app/store"
	"log"
	"strconv"

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

	messageField := []*discordgo.MessageEmbedField{}
	counters, _ := storeClient.GetCounters()
	for _, counter := range counters {
		user, err := session.User(counter.UserID)
		if err != nil {
			log.Println(err)
			return
		}

		messageField = append(messageField, &discordgo.MessageEmbedField{
			Name:   user.Username,
			Value:  strconv.Itoa(counter.Count) + " pt",
			Inline: true,
		})
	}

	_, err = b.session.ChannelMessageEditEmbed(event.ChannelID, event.MessageID, &discordgo.MessageEmbed{
		Title:       "Counter",
		Description: "Increment :arrow_up:, Decrement :arrow_down:, Reset :zero:",
		Fields:      messageField,
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

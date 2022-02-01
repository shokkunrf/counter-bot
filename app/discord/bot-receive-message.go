package discord

import (
	"app/store"
	"log"
	"regexp"
	"strings"

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

	storeClient, _ := store.GetClient()

	// メッセージを取得
	str := regexp.MustCompile(`<@\!\d*>`).ReplaceAllString(event.Content, "")
	cmd := strings.TrimSpace(str)
	switch cmd {
	case START_MESSAGE:
		counters, _ := storeClient.GetCounters()

		embeddedMessage, err := b.generateEmbeddedMessage(counters)
		if err != nil {
			log.Println(err)
			return
		}

		msg, err := b.session.ChannelMessageSendEmbed(event.ChannelID, &embeddedMessage)
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
	case CLEAR_MESSAGE:
		storeClient.ClearCounters()
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

package discord

import (
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) receiveMessage(session *discordgo.Session, event *discordgo.MessageCreate) {
	// mentionされた場合にチャンネルを設定する
	me, err := session.User("@me")
	if err != nil {
		log.Println(err)
		return
	}
	for _, user := range event.Mentions {
		if user.ID == me.ID {
			err := b.setChannelID(event.ChannelID)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}

	// 設定されたチャンネルのみ通す
	if b.channelID != event.ChannelID {
		return
	}

	// メッセージを取得
	str := regexp.MustCompile(`<@\!\d*>`).ReplaceAllString(event.Content, "")
	message := strings.TrimSpace(str)
	if message == "" {
		return
	}
	// push
}

func (b *Bot) setChannelID(channelID string) error {
	b.channelID = channelID

	channel, err := b.session.Channel(channelID)
	if err != nil {
		return err
	}
	err = b.session.UpdateGameStatus(0, "Watching "+channel.Name)
	if err != nil {
		return err
	}
	return nil
}

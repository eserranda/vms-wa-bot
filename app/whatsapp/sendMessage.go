package whatsapp

import (
	"context"
	"log"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
)

func (w *WhatsappmeowClient) SendMessage(ctx context.Context, infoChat, infoSender types.JID, message string) error {
	msg := &waE2E.Message{
		Conversation: &message,
	}

	_, err := w.client.SendMessage(ctx, infoChat, msg)
	if err != nil {
		return err
	}

	// end typing status
	err = w.client.SendChatPresence(infoSender, "paused", "")
	if err != nil {
		log.Println(err)
	}

	return nil
}

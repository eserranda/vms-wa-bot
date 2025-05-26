package model

import (
	"context"

	"go.mau.fi/whatsmeow/types"
)

type WhatsAppClient interface {
	Connect(ctx context.Context) error
	Disconnect() error
	SendMessage(ctx context.Context, infoChat, infoSender types.JID, message string) error
}

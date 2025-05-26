package whatsapp

import (
	"context"
	"fmt"
	"os"

	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type WhatsappmeowClient struct {
	client *whatsmeow.Client
}

func NewWhatsappmeowClient() (*WhatsappmeowClient, error) {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	ctx := context.Background()
	// Inisialisasi database store
	if os.Getenv("WHATSAPP_DB_NAME") == "" {
		return nil, fmt.Errorf("WHATSAPP_DB_NAME environment variable is not set")
	}

	sql, err := sqlstore.New(ctx, "sqlite3", fmt.Sprintf("file:%s.db?_foreign_keys=on", os.Getenv("WHATSAPP_DB_NAME")), dbLog)
	if err != nil {
		panic(err)
	}

	// Mendapatkan perangkat pertama dari store
	deviceStore, err := sql.GetFirstDevice(ctx)
	if err != nil {
		return nil, err
	}

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	// Membuat Whatsmeow client
	client := whatsmeow.NewClient(
		deviceStore,
		clientLog,
	)

	return &WhatsappmeowClient{
		client: client,
	}, nil
}

func (w *WhatsappmeowClient) Connect(ctx context.Context) error {
	if w.client.Store.ID == nil {
		qrChan, _ := w.client.GetQRChannel(ctx)
		err := w.client.Connect()
		if err != nil {
			return err
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("Scan the QR Code with your WhatsApp app:")
				fmt.Println("Code:", evt.Code)
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.M, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err := w.client.Connect()
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *WhatsappmeowClient) Disconnect() error {
	w.client.Disconnect()

	return nil
}

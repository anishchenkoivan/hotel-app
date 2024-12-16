package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/anishchenkoivan/hotel-app/notification-service/internal/model"
	"io"
	"log"
	"net/http"
)

type BotClient struct {
	TelegramBotUrl string
	client         *http.Client
}

func NewBotClient(url string) *BotClient {
	return &BotClient{url, http.DefaultClient}
}

func (c *BotClient) SendMessage(message model.Message) error {
	text := fmt.Sprintf(
		"You have booked room %s from %s to %s at a price of %d.\nYour reservation ID is %s",
		message.RoomId.String(),
		message.InTime.String(),
		message.OutTime.String(),
		message.Cost,
		message.ReservationId.String(),
	)

	messageDto := TelegramMessageDto{TelegramId: message.TelegramId, Message: text}

	jsonMessage, err := json.Marshal(messageDto)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.TelegramBotUrl+"/notify", bytes.NewBuffer(jsonMessage))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error while closing response body")
		}
	}(resp.Body)

	if !((resp.StatusCode >= 200) && (resp.StatusCode <= 299)) {
		log.Printf("Error while sending message: %s", resp.Status)
	}

	return nil
}

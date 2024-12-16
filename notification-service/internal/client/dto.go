package client

type TelegramMessageDto struct {
	TelegramId string `json:"tgId"`
	Message    string `json:"message"`
}

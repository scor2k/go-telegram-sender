package gotelegramsender

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var telegramToken = os.Getenv("TELEGRAM_TOKEN")
var telegramChatID = os.Getenv("TELEGRAM_CHAT")

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// SendMessage via Telegram
func SendMessage(message string) error {
	chatID, err := strconv.ParseInt(telegramChatID, 10, 64)

	if err != nil {
		fmt.Printf("Setup TELEGRAM_CHAT env")
		return errors.New("Setup TELEGRAM_CHAT env")
	}

	if telegramToken == "" {
		fmt.Printf("Setup TELEGRAM_TOKEN env")
		return errors.New("Setup TELEGRAM_TOKEN env")
	}

	// Creates an instance of our custom sendMessageReqBody Type
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   message,
	}

	// Convert our custom type into json format
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Make a request to send our message using the POST method to the telegram bot API
	resp, err := http.Post(
		"https://api.telegram.org/bot"+telegramToken+"/"+"sendMessage",
		"application/json",
		bytes.NewBuffer(reqBytes),
	)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + resp.Status)
	}

	return err
}

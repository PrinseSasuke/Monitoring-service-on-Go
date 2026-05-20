package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

func sendTelegram(botToken, chatID, message string) {
	message = url.QueryEscape(message)

	apiURL := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		botToken,
		chatID,
		message,
	)

	resp, err := http.Get(apiURL)

	if err != nil {
		fmt.Println("TELEGRAM ERROR:", err)
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func checkServer(address string) bool {
	conn, err := net.DialTimeout(
		"tcp",
		address,
		5*time.Second,
	)

	if err != nil {
		return false
	}

	conn.Close()

	return true
}

func main() {
	address := "YOUR_ADRESS:22"

	botToken := "YOUR_TOKEN"
	chatID := "CHAT_ID"

	serverWasUp := true

	for {
		isUp := checkServer(address)

		if !isUp && serverWasUp {
			fmt.Println("SERVER DOWN")

			sendTelegram(
				botToken,
				chatID,
				"SERVER DOWN",
			)

			serverWasUp = false
		}

		if isUp && !serverWasUp {
			fmt.Println("SERVER RECOVERED")

			sendTelegram(
				botToken,
				chatID,
				"SERVER RECOVERED",
			)

			serverWasUp = true
		}

		time.Sleep(10 * time.Second)
	}
}
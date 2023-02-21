package main

import (
	"antonovskie_apples_bot/handler"
	"antonovskie_apples_bot/tgclient"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	start()

	ch := make(chan any)
	bot := initBot()

	collectBotInfo(bot, ch)
	listenUpdates(bot, ch)
}

func start() {
	fmt.Println("Hello, 🍎🍏❤️")
	fmt.Printf("Starting bot...\n\n")

	initEnv()
}

func initEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
}

func initBot() tgclient.TelegramBot {
	token := os.Getenv("TG_BOT_TOKEN")

	return tgclient.TelegramBot{Token: token}
}

func collectBotInfo(bot tgclient.Client, ch chan any) {
	bot.GetMe(ch)
	handler.HandleGetMeResponse(ch)

	fmt.Printf("\nBot is ready to handle connections! :)\n\n")
}

func listenUpdates(bot tgclient.TelegramBot, ch chan any) {
	var wg sync.WaitGroup

	for {
		wg.Add(1)

		bot.GetUpdates(ch, &wg)
		handler.HandleUpdateResponse(&bot, ch)

		wg.Wait()
	}
}

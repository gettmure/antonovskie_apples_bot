package main

import (
	"fmt"
	"log"
	"os"

	"antonovskie_apples_bot/tgclient"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, 🍎🍏❤️")
	initEnv()

	ch := make(chan interface{})
	bot := initBot()

	bot.GetMe(ch)
	response := <-ch

	switch p := response.(type) {
	case *tgclient.Response[tgclient.GetMeResponse]:
		fmt.Println(p.Result.Firstname)
		fmt.Println(*p.Result.Username)
	}
}

func initEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading .env file %s", err))
	}
}

func initBot() tgclient.Client {
	token := os.Getenv("TG_BOT_TOKEN")

	return tgclient.TelegramBot{Token: token}
}

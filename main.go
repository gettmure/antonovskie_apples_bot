package main

import (
	"antonovskie_apples_bot/tgclient"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/recoilme/graceful"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	graceful.Unignore(quit, func() error {
		cancel()
		time.Sleep(1 * time.Second)

		return nil
	}, graceful.Terminate...)

	start()

	token := os.Getenv("TG_BOT_TOKEN")
	if len(token) < 1 {
		log.Fatalln("Empty token")
	}

	client := tgclient.InitClient()
	bot, err := tgclient.InitBot(token, client)
	if err != nil {
		log.Fatalln(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	bot.ListenUpdates(ctx, &wg)
	wg.Wait()
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

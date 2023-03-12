package main

import (
	"antonovskie_apples_bot/tgclient"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
		log.Fatalln("Empty tg bot token")
	}

	client := tgclient.InitClient()
	bot, err := tgclient.InitBot(token, client)
	if err != nil {
		log.Fatalln(err)
	}

	bot.ListenUpdates(ctx)
}

func start() {
	fmt.Println("Hello, 🍎🍏❤️")
	fmt.Printf("Starting bot...\n\n")

	initEnv()
}

func initEnv() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("Unable to identify current directory (needed to load .env)")
	}

	basepath := filepath.Dir(file)
	err := godotenv.Load(filepath.Join(basepath, ".env"))
	if err != nil {
		log.Fatalln(err)
	}
}

package tgclient

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func logGetMe(getMe *GetMeResponse) {
	fmt.Println("Bot info:")
	fmt.Println("- id:", getMe.Result.Id)
	fmt.Println("- public name:", getMe.Result.Firstname)
	fmt.Println("- username:", getMe.Result.Username)
	fmt.Print("\nBot is ready to handle updates!\n\n")
}

func logUpdate(update *UpdateResponse) {
	location, err := time.LoadLocation("Asia/Vladivostok")
	if err != nil {
		fmt.Println(err)

		return
	}

	now := time.Now().In(location)
	path := "logs"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	filename := fmt.Sprintf("%s.txt", now.Format("2006-01-02"))
	logFile, err := os.OpenFile(filepath.Join(path, filepath.Base(filename)), os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	defer logFile.Close()
	if err != nil {
		fmt.Println(filename, err)

		return
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	logMessage := fmt.Sprintf("\nReceived update: \n- update id: %d\n- chat id: %d\n- message id: %d\n- message text: %s\n- user first_name: %s\n- user last_name: %s\n- user username: %s\n",
		update.UpdateId,
		update.Message.Chat.Id,
		update.Message.MessageId,
		update.Message.Text,
		update.Message.From.FirstName,
		update.Message.From.LastName,
		update.Message.From.Username,
	)

	if update.Message.Audio != nil {
		logMessage += fmt.Sprintf("- audio file id:%s\n- audio file unique id:%s\n", update.Message.Audio.FileId, update.Message.Audio.FileUniqueId)
	}

	log.Println(logMessage)
}

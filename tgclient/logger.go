package tgclient

import "fmt"

func logGetMe(getMe *GetMeResponse) {
	fmt.Println("Bot info:")
	fmt.Println("- id:", getMe.Result.Id)
	fmt.Println("- public name:", getMe.Result.Firstname)
	fmt.Println("- username:", getMe.Result.Username)
}

func logUpdate(update *UpdateResponse) {
	fmt.Println("Received update:")
	fmt.Println("- update id:", update.UpdateId)
	fmt.Println("- chat id:", update.Message.Chat.Id)
	fmt.Println("- message id:", update.Message.MessageId)
	fmt.Println("- message text:", update.Message.Text)
}

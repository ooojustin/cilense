package chat

import "fmt"

func JoinRoom(data map[string]interface{}) {
	roomID := data["room_id"]
	fmt.Println("user joining room", roomID)
}

func SendMessage(data map[string]interface{}) {
	message := data["message"]
	fmt.Println("user sending message:", message)
}

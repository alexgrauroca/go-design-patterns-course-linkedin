package main

import "fmt"

func main() {
	bldr := newNotificationBuilder()

	bldr.SetTile("New Notification")
	bldr.SetIcon("icon.png")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(5)
	bldr.SetMessage("This is a basic Notification with some text in it")
	bldr.SetType("alert")
	bldr.SetSubTitle("This is a subtitle")

	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating the notification:", err)
	} else {
		fmt.Printf("Notification: %+v\n", notif)
	}
}

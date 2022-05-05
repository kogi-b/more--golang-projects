package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main(){

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3414052338899-3400128760247-sccWB7zF4PArgyIDf9T7kXwV")
	os.Setenv("CHANNEL_ID", "C03C619MS74")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Happiness.pdf"}

	for i := 0; i<len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil{
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
package main

import (
	"fmt"
	"os"

	"github.com/jcksnvllxr80/go-tuts/slack-file-bot/secrets"
	"github.com/slack-go/slack"
)

func main() {
	secrets.SetEnvVars()
	api := slack.New(os.Getenv("SLACK_FILE_BOT_TOKEN"))
	channelArr := []string{os.Getenv("SLACK_FILE_BOT_CHANNEL_ID")}
	fileArr := []string{"files/HelloSlack.txt", "files/ZIPL.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}

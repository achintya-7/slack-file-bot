package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)



func main()  {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3957073649254-3987523392976-m6Cwx5ixIkXjJLaiUDBpNznc")
	os.Setenv("CHANNEL_ID", "C03U8M3C65B")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	// channels and files
	channelArr := [] string {os.Getenv("CHANNEL_ID")}
	fileArr := [] string {"Resume3.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("Error : %s", err)
			return
		}
		fmt.Printf("Name : %s\nUrl : %s\n", file.Name, file.URL)
	}
}

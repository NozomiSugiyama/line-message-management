package main

import (
	"log"
	"net/http"
	"os"

	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	lineChannelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineChannelAccessToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	providerWebOrigin := os.Getenv("PROVIDER_WEB_ORIGIN")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/hook", func(c *gin.Context) {
		client := &http.Client{Timeout: time.Duration(15 * time.Second)}
		bot, err := linebot.New(lineChannelSecret, lineChannelAccessToken, linebot.WithHTTPClient(client))
		if err != nil {
			fmt.Println(err)
			return
		}
		received, err := bot.ParseRequest(c.Request)

		for _, event := range received {
			log.Print(linebot.EventTypeMessage)
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					source := event.Source
					if source.Type == linebot.EventSourceTypeUser {
						switch message.Text {
						case "連携":
							res, err := bot.IssueLinkToken(source.UserID).Do()
							if err != nil {
								log.Print(err)
								return
							}
							postMessage := linebot.NewTextMessage(providerWebOrigin + "/client-sign-in?link-token=" + res.LinkToken)
							if _, err = bot.ReplyMessage(event.ReplyToken, postMessage).Do(); err != nil {
								log.Print(err)
							}
						default:
							postMessage := linebot.NewTextMessage("reply :" + message.Text)
							if _, err = bot.ReplyMessage(event.ReplyToken, postMessage).Do(); err != nil {
								log.Print(err)
							}
						}
					}
				}
			}
		}
	})

	router.Run(":" + port)
}

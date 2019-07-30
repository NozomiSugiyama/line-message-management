package main

import (
	"log"
	"net/http"
	"os"

	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"math/rand"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	lineChannelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineChannelAccessToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")

	log.Print("=============")
	log.Print(lineChannelAccessToken)
	log.Print("=============")

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
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					source := event.Source
					if source.Type == linebot.EventSourceTypeRoom {
						if resMessage := getResMessage(message.Text); resMessage != "" {
							postMessage := linebot.NewTextMessage(resMessage)
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

func getResMessage(reqMessage string) (message string) {
	resMessages := [3]string{"testes", "test2", "test3"}

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(5) == 0 {
		if math := rand.Intn(4); math != 3 {
			message = resMessages[math]
		} else {
			message = "replay :" + reqMessage
		}
	}
	return message
}

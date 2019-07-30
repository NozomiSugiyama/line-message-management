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
	log.Print("======= Debug point 1 ======")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/hook", func(c *gin.Context) {
		log.Print("======= Debug point 2 ======")
		client := &http.Client{Timeout: time.Duration(15 * time.Second)}
		bot, err := linebot.New(lineChannelSecret, lineChannelAccessToken, linebot.WithHTTPClient(client))
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Print("======= Debug point 3 ======")
		received, err := bot.ParseRequest(c.Request)

		for _, event := range received {
			log.Print("======= Debug point 4 ======")
			log.Print(linebot.EventTypeMessage)
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					log.Print("======= Debug point 5 ======")
					source := event.Source
					if source.Type == linebot.EventSourceTypeUser {
						log.Print("======= Debug point 6 ======")
						if resMessage := getResMessage(message.Text); resMessage != "" {
							log.Print("======= Debug point 7 ======")
							postMessage := linebot.NewTextMessage(resMessage)
							if _, err = bot.ReplyMessage(event.ReplyToken, postMessage).Do(); err != nil {
								log.Print(err)
							}
							log.Print("======= Debug point 8 ======")
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
	return
}

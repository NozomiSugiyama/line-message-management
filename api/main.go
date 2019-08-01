package main

import (
	"log"
	"net/http"
	"os"

	"fmt"
	"github.com/jinzhu/gorm"
	"time"

	"api/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {

	var mode string
	if len(os.Args) > 0 {
		mode = os.Args[1]
	}
	var envFile string
	if mode == "dev" {
		envFile = ".env.development"
	} else {
		envFile = ".env.production"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Println("Failed load " + envFile + " file")
	}

	port := os.Getenv("PORT")
	lineChannelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineChannelAccessToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	providerWebOrigin := os.Getenv("PROVIDER_WEB_ORIGIN")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresDatabase := os.Getenv("POSTGRES_DATABASE")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresSSLMode := os.Getenv("POSTGRES_SSL_MODE")

	if port == "" {
		port = "8080"
	}

	var db *gorm.DB

	for i := 0; i < 10; i++ {
		db, err = database.Initialize(postgresHost, postgresPort, postgresUser, postgresDatabase, postgresPassword, postgresSSLMode)
		if err != nil {
			log.Println("Failed database initialize")
			log.Println(err)
			if i >= 9 {
				panic(err)
			}
			time.Sleep(10 * time.Second)
		} else {
			log.Println("Connected database")
			break
		}
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(database.Inject(db))

	router.POST("/hook", func(c *gin.Context) {
		client := &http.Client{Timeout: time.Duration(15 * time.Second)}
		bot, err := linebot.New(lineChannelSecret, lineChannelAccessToken, linebot.WithHTTPClient(client))
		if err != nil {
			fmt.Println(err)
			return
		}
		received, err := bot.ParseRequest(c.Request)

		for _, event := range received {
			switch event.Type {
			case linebot.EventTypeAccountLink:
				postMessage := linebot.NewTextMessage("replyToken :" + event.ReplyToken + "\n" + "link nonce :" + event.AccountLink.Nonce + "\n" + "result: " + string(event.AccountLink.Result))
				if _, err = bot.ReplyMessage(event.ReplyToken, postMessage).Do(); err != nil {
					log.Print(err)
				}
			case linebot.EventTypeMessage:
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

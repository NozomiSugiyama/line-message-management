package handler

import (
	"api/database/model"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"

	"fmt"
	"net/http"
	"time"
)

// HookHandler handler model
type HookHandler struct {
	userRepository         *model.UserRepository
	nonceRepository        *model.NonceRepository
	lineUserRepository     *model.LineUserRepository
	lineChannelSecret      string
	lineChannelAccessToken string
	providerWebOrigin      string
}

// NewHookHandler create new hook handler
func NewHookHandler(userRepo *model.UserRepository, nonceRepo *model.NonceRepository, lineUserRepo *model.LineUserRepository, lineChannelSecret string, lineChannelAccessToken string, providerWebOrigin string) *HookHandler {
	h := new(HookHandler)
	h.userRepository = userRepo
	h.nonceRepository = nonceRepo
	h.lineUserRepository = lineUserRepo
	h.lineChannelSecret = lineChannelSecret
	h.lineChannelAccessToken = lineChannelAccessToken
	h.providerWebOrigin = providerWebOrigin
	return h
}

func (h *HookHandler) PostHook(c *gin.Context) {
	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	bot, err := linebot.New(h.lineChannelSecret, h.lineChannelAccessToken, linebot.WithHTTPClient(client))
	if err != nil {
		fmt.Println(err)
		return
	}
	received, err := bot.ParseRequest(c.Request)

	for _, event := range received {
	switch event.Type {
	case linebot.EventTypeAccountLink:
		nonce, err := h.nonceRepository.FindNonceByNonce(event.AccountLink.Nonce)
		lineUser := model.LineUser{
			UserID:        nonce.UserID,
			LineID:        event.Source.UserID,
			LinkedAccount: nonce.LinkedAccount,
		}
		h.lineUserRepository.CreateLineUser(&lineUser)
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
						postMessage := linebot.NewTextMessage(h.providerWebOrigin + "/client-sign-in?link-line-token=" + res.LinkToken)
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
}

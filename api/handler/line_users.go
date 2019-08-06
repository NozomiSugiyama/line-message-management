package handler

import (
	"api/database/model"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type LineUserHandler struct {
	lineUserRepository     *model.LineUserRepository
	lineChannelSecret      string
	lineChannelAccessToken string
}

func NewLineUserHandler(lineUserRepo *model.LineUserRepository, lineChannelSecret string, lineChannelAccessToken string) *LineUserHandler {
	h := new(LineUserHandler)
	h.lineUserRepository = lineUserRepo
	h.lineChannelSecret = lineChannelSecret
	h.lineChannelAccessToken = lineChannelAccessToken
	return h
}

func (h *LineUserHandler) GetLineUsers(c *gin.Context) {

	lineUsers, err := h.lineUserRepository.ListLineUsersWithUser()
	if err != nil {
		if err == model.ErrRecordNotFound {
			c.Status(404)
			return
		}
		c.Error(err)
		return
	}

	bot, err := linebot.New(h.lineChannelSecret, h.lineChannelAccessToken)

	var lineUserResponses []LineUserResponse
	for _, lineUser := range lineUsers {
		userProfile, err := bot.GetProfile(lineUser.LineID).Do()
		if err == nil {
			lineUserResponses = append(
				lineUserResponses,
				LineUserResponse{
					ID:     lineUser.ID,
					UserID: lineUser.UserID,
					LineID: lineUser.LineID,
					DisplayName: &userProfile.DisplayName,
					User: UserResponse{
						ID:    lineUser.User.ID,
						Name:  lineUser.User.Name,
						Email: lineUser.User.Email,
					},
					LinkedAccount: lineUser.LinkedAccount,
				},
			)
		} else {
			lineUserResponses = append(
				lineUserResponses,
				LineUserResponse{
					ID:     lineUser.ID,
					UserID: lineUser.UserID,
					LineID: lineUser.LineID,
					User: UserResponse{
						ID:    lineUser.User.ID,
						Name:  lineUser.User.Name,
						Email: lineUser.User.Email,
					},
					LinkedAccount: lineUser.LinkedAccount,
				},
			)
		}
	}
	c.JSON(200, lineUserResponses)
}

func (h *LineUserHandler) GetLineUserByID(c *gin.Context) {
	lineUserID, err := strconv.Atoi(c.Param("lineuserid"))

	if err != nil {
		c.Status(400)
		return
	}

	lineUser, err := h.lineUserRepository.FindLineUserWithUserByID(lineUserID)
	if err != nil {
		if err == model.ErrRecordNotFound {
			c.Status(404)
			return
		}
		c.Error(err)
		return
	}

	lineUserResponse := LineUserResponse{
		ID:     lineUser.ID,
		UserID: lineUser.UserID,
		LineID: lineUser.LineID,
		User: UserResponse{
			ID:    lineUser.User.ID,
			Name:  lineUser.User.Name,
			Email: lineUser.User.Email,
		},
		LinkedAccount: lineUser.LinkedAccount,
	}

	bot, err := linebot.New(h.lineChannelSecret, h.lineChannelAccessToken)
	userProfile, err := bot.GetProfile(lineUser.LineID).Do()

	if err == nil {
		lineUserResponse.DisplayName = &userProfile.DisplayName
	}

	c.JSON(200, lineUserResponse)
}

// SendMessageByLineUserID a
func (h *LineUserHandler) SendMessageByLineUserID(c *gin.Context) {
	lineUserID := c.Param("lineuserid")
	var message Message
	c.BindJSON(&message)

	bot, err := linebot.New(h.lineChannelSecret, h.lineChannelAccessToken)

	_, err = bot.PushMessage(lineUserID, linebot.NewTextMessage(message.Messages)).Do()
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(200)
}

// SendTestMessagesByLineUserID a
func (h *LineUserHandler) SendTestMessagesByLineUserID(c *gin.Context) {
	lineUserID := c.Param("lineuserid")

	bot, err := linebot.New(h.lineChannelSecret, h.lineChannelAccessToken)

	sendingMessages := []linebot.SendingMessage{
		linebot.NewTextMessage("Hello, world"),
		linebot.NewImageMessage("http://line-message-management-ui.herokuapp.com/assets/img/no-image.png", "https://example.com/preview.jpg"),
		linebot.NewStickerMessage("1", "1"),
		linebot.NewLocationMessage("title", "address", 35.65910807942215, 139.70372892916203),
		linebot.NewTemplateMessage(
			"this is a buttons template",
			linebot.NewButtonsTemplate(
				"https://example.com/bot/images/image.jpg",
				"Menu",
				"Please select",
				NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
				NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
				NewURIAction("View detail", "https://example.com/page/123"),
			),
		),

	}

	for _, message := range sendingMessages {
		_, err = bot.PushMessage(lineUserID, message).Do()
		if err != nil {
			c.Error(err)
			return
		}
	}

	c.Status(200)
}

// LineUserResponse lineUser response model
type LineUserResponse struct {
	ID            int          `json:"id"`
	UserID        int          `json:"user_id"`
	User          UserResponse `json:"user"`
	LineID        string       `json:"line_id"`
	LinkedAccount string       `json:"linked_account"`
	DisplayName   *string      `json:"display_name"`
}

type Message struct {
	Messages string `json:"message"`
}

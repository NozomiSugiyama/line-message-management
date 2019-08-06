package handler

import (
	"fmt"
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
		linebot.NewImageMessage("https://picsum.photos/id/1/300/300", "https://picsum.photos/id/2/300/300"),
		linebot.NewStickerMessage("1", "1"),
		linebot.NewLocationMessage("title", "address", 35.65910807942215, 139.70372892916203),
		linebot.NewTemplateMessage(
			"this is a buttons template",
			linebot.NewButtonsTemplate(
				"https://picsum.photos/id/3/300/300",
				"Menu",
				"Please select",
				linebot.NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
				linebot.NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
				linebot.NewURIAction("View detail", "https://example.com/page/123"),
			),
		),
		linebot.NewTemplateMessage(
			"this is a buttons template",
			linebot.NewButtonsTemplate(
				"https://picsum.photos/id/4/300/300",
				"Menu",
				"Please select a date, time or datetime",
				linebot.NewDatetimePickerAction("Date", "action=sel&only=date", "date", "2017-09-01", "2017-09-03", ""),
				linebot.NewDatetimePickerAction("Time", "action=sel&only=time", "time", "", "23:59", "00:00"),
				linebot.NewDatetimePickerAction("DateTime", "action=sel", "datetime", "2017-09-01T12:00", "", ""),
			),
		),
		linebot.NewTemplateMessage(
			"this is a image carousel template",
			linebot.NewImageCarouselTemplate(
				linebot.NewImageCarouselColumn(
					"https://picsum.photos/id/5/300/300",
					linebot.NewURIAction("View detail", "https://example.com/page/111"),
				),
			),
		),
		linebot.NewFlexMessage(
			"this is a flex message",
			&linebot.BubbleContainer{
				Type: linebot.FlexContainerTypeBubble,
				Body: &linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeVertical,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type: linebot.FlexComponentTypeText,
							Text: "hello",
						},
						&linebot.TextComponent{
							Type: linebot.FlexComponentTypeText,
							Text: "world",
							Flex: linebot.IntPtr(0),
						},
						&linebot.SpacerComponent{
							Type: linebot.FlexComponentTypeSpacer,
						},
					},
				},
			},
		),
	}

	for _, message := range sendingMessages {
		_, err = bot.PushMessage(lineUserID, message).Do()
		if err != nil {
			fmt.Println(err)
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

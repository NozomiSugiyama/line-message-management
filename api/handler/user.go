package handler

import (
	"api/database/model"

	"github.com/gin-gonic/gin"
	"strconv"
)

type UserHandler struct {
	userRepository *model.UserRepository
}

func NewUserHandler(userRepo *model.UserRepository) *UserHandler {
	h := new(UserHandler)
	h.userRepository = userRepo
	return h
}

func (h *UserHandler) GetUsers(c *gin.Context) {

	users, err := h.userRepository.ListUsers()
	if err != nil {
		if err == model.ErrRecordNotFound {
			c.Status(404)
			return
		}
		c.Error(err)
		return
	}

	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(
			userResponses,
			UserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
		)
	}
	c.JSON(200, userResponses)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userid"))

	if err != nil {
		c.Status(400)
		return
	}

	user, err := h.userRepository.FindUserByID(userID)
	if err != nil {
		if err == model.ErrRecordNotFound {
			c.Status(404)
			return
		}
		c.Error(err)
		return
	}

	userResponse := UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(200, userResponse)
}

// UserResponse user response model
type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

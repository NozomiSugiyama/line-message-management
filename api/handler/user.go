package handler

import (
	"api/database/model"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepository *model.UserRepository
}

func NewUserHandler(repo *model.UserRepository) *UserHandler {
	h := new(UserHandler)
	h.userRepository = repo
	return h
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("userid")
	var user, err = h.userRepository.FindUserByID(userID)
	if err != nil {
		c.Error(err)
		return
	}

	userResponse := UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	c.JSON(200, userResponse)
}

// UserResponse user response model
type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

package handler

import (
	"api/database/model"
	crand "crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AuthHandler handler model
type AuthHandler struct {
	userRepository  *model.UserRepository
	nonceRepository *model.NonceRepository
}

// NewAuthHandler create new auth handler
func NewAuthHandler(userRepo *model.UserRepository, nonceRepo *model.NonceRepository) *AuthHandler {
	h := new(AuthHandler)
	h.userRepository = userRepo
	h.nonceRepository = nonceRepo
	return h
}

func (h *AuthHandler) ClientSignIn(c *gin.Context) {
	var requestCredential RequestCredential
	c.BindJSON(&requestCredential)
	linkLine := c.Query("link-line")

	user, err := h.userRepository.FindUserByEMail(requestCredential.Email)
	if err != nil {
		c.Error(err)
		return
	}

    err = passwordVerify(user.Password, requestCredential.Password)
    if err != nil {
		c.Error(err)
		fmt.Println("password verify error")
		return
    }

	var credential Credential
	if linkLine == "true" {
		fmt.Println("linkLine == 'true'")
		nonce := model.Nonce{
			UserID:        user.ID,
			Nonce:         secureRandomStr(16),
			LinkedAccount: "main",
		}

		err = h.nonceRepository.CreateNonce(&nonce)

		credential = Credential{
			Email:     user.Email,
			LineNonce: &nonce.Nonce,
		}
	} else {
		fmt.Println("linkLine == 'false'")
		credential = Credential{
			Email: user.Email,
		}
	}
	c.JSON(200, credential)

}

type RequestCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Credential struct {
	Email     string  `json:"email"`
	LineNonce *string `json:"line_nonce"`
}

func secureRandomStr(b int) string {
	k := make([]byte, b)
	if _, err := crand.Read(k); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", k)
}


func passwordHash(pw string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hash), err
}

func passwordVerify(hash, pw string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

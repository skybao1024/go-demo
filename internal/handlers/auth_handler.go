package handlers

import (
	"github.com/gin-gonic/gin"
	"myproject/internal/api/requests/authrequest"
	"myproject/internal/api/response"
	"myproject/internal/services/authservice"
)

type AuthHandler struct {
	authService authservice.ServiceInterface
}

func NewAuthHandler(authService authservice.ServiceInterface) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register handles user registration by binding JSON input and calling the authService to register the user.
//
// The function takes a gin Context as a parameter.
// It returns no value, but sends a response to the client.
func (h *AuthHandler) Register(c *gin.Context) {
	var input authrequest.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.authService.Register(input.Username, input.Password, input.Email); err != nil {
		response.Fail(c, 1001, "Failed to register user")
		return
	}

	response.SuccessWithoutData(c)
}

// Login handles user login by binding JSON input and calling the authService to generate a token.
//
// The function takes a gin Context as a parameter.
// It returns no value, but sends a response to the client containing the JWT token.
func (h *AuthHandler) Login(c *gin.Context) {
	var input authrequest.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	token, err := h.authService.Login(input.Username, input.Password)
	if err != nil {
		response.Unauthorized(c, "Invalid credentials")
		return
	}

	response.Success(c, gin.H{
		"access_token": token,
		"token_type":   "Bearer",
	})
}

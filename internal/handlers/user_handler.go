package handlers

import (
	"github.com/gin-gonic/gin"
	"myproject/internal/api/response"
	"myproject/internal/middleware"
	"myproject/internal/services/userservice"
)

type UserHandler struct {
	userService userservice.ServiceInterface
}

func NewUserHandler(userService userservice.ServiceInterface) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Profile(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		response.Unauthorized(c, "User not authenticated")
		return
	}

	user, err := h.userService.Profile(userID)
	if err != nil {
		response.InternalServerError(c, "Failed to get user information")
		return
	}

	response.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

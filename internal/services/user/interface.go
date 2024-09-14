package userservice

import "myproject/internal/models"

type ServiceInterface interface {
	Profile(id uint) (*models.User, error)
}

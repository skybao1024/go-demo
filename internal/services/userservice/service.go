package userservice

import (
	"errors"
	"gorm.io/gorm"
	"myproject/internal/models"
	"myproject/internal/services"
)

type Service struct {
	services.BaseService
}

func NewService(config *services.ServiceConfig) ServiceInterface {
	return &Service{
		BaseService: services.NewBaseService(config),
	}
}

func (s *Service) Profile(id uint) (*models.User, error) {
	var user models.User
	if err := s.Config.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

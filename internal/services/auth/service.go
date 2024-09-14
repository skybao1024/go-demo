package authservice

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func (s *Service) Register(username, password, email string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	return s.Config.DB.Create(&user).Error
}

func (s *Service) Login(username, password string) (string, error) {
	var user models.User
	if err := s.Config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	log.Println(user)
	token, err := s.generateToken(user)
	log.Println(token)
	log.Println(err)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) generateToken(user models.User) (string, error) {
	log.Println(s.Config.JWTExpires)
	log.Println(s.Config.JWTSecret)
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      jwt.NewNumericDate(time.Now().Add(s.Config.JWTExpires)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.Config.JWTSecret))
}

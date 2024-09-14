package app

import (
	"myproject/internal/config"
	"myproject/internal/db"
	"myproject/internal/services"
	"myproject/internal/services/auth"
	"myproject/internal/services/user"
)

type Services struct {
	Auth authservice.ServiceInterface
	User userservice.ServiceInterface
	// 其他服务...
}

func NewServices(cfg *config.Config, database db.Database, redisClient db.Redis) *Services {
	serviceConfig := &services.ServiceConfig{
		DB:         database,
		Redis:      redisClient,
		JWTSecret:  cfg.JWTSecret,
		JWTExpires: cfg.JWTExpires,
		// 其他配置...
	}

	return &Services{
		Auth: authservice.NewService(serviceConfig),
		User: userservice.NewService(serviceConfig),
		// 初始化其他服务...
	}
}

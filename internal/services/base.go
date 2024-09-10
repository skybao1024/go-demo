package services

import (
	"myproject/internal/db"
	"time"
)

type ServiceConfig struct {
	DB         db.Database
	Redis      db.Redis
	JWTSecret  string
	JWTExpires time.Duration
	// 可以添加其他全局配置
}

type BaseService struct {
	Config *ServiceConfig
}

func NewBaseService(config *ServiceConfig) BaseService {
	return BaseService{Config: config}
}

package db

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myproject/internal/config"
)

// NewMySQLConnection 创建一个新的 MySQL 数据库连接
func NewMySQLConnection(cfg config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// NewRedisClient 创建一个新的 Redis 客户端
func NewRedisClient(cfg config.RedisConfig) Redis {
	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:         cfg.Address,
			Password:     cfg.Password,
			DB:           cfg.Database,
			PoolSize:     cfg.PoolSize,
			MinIdleConns: cfg.MinIdleConns,
			MaxConnAge:   cfg.MaxConnAge,
		}),
	}
}

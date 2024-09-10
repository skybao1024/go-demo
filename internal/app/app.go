package app

import (
	"myproject/internal/config"
	"myproject/internal/db"
)

type App struct {
	Config   *config.Config
	DB       db.Database
	Redis    db.Redis
	Services *Services
}

func InitApp(cfg *config.Config) (*App, error) {
	database, err := db.NewMySQLConnection(cfg.DBConfig)
	if err != nil {
		return nil, err
	}

	redisClient := db.NewRedisClient(cfg.RedisConfig)

	app := &App{
		Config: cfg,
		DB:     db.NewGormDB(database),
		Redis:  redisClient,
	}

	app.Services = NewServices(cfg, app.DB, app.Redis)

	return app, nil
}

func (a *App) Close() error {
	if closer, ok := a.DB.(interface{ Close() error }); ok {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	if closer, ok := a.Redis.(interface{ Close() error }); ok {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

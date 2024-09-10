package main

import (
	"log"
	"myproject/internal/app"

	"myproject/internal/config"
	"myproject/internal/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	application, err := app.InitApp(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// 确保在程序结束时关闭数据库连接
	defer application.Close()

	r := router.SetupRouter(application)

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

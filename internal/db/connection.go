package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"short-url-api/internal/config"
	model "short-url-api/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	// Lấy DSN từ file config
	dsn := config.Load().DNS()

	// Cấu hình GORM
	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Kết nối MySQL
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), cfg)
	if err != nil {
		return fmt.Errorf("error opening DB connection: %w", err)
	}

	// Lấy sql.DB để cấu hình connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("error getting sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	// Kiểm tra kết nối với timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		sqlDB.Close()
		return fmt.Errorf("DB ping error: %w", err)
	}

	log.Println("Connected to MySQL successfully!")

	if err := DB.AutoMigrate(&model.Link{}, &model.LinkClick{}); err != nil {
		return fmt.Errorf("failed to migrate DB: %w", err)
	}

	log.Println("Database migrated successfully!")

	return nil
}

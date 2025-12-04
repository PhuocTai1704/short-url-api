package config

import (
	"fmt"
	"short-url-api/internal/utils"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Config struct {
	DB DatabaseConfig
}

func Load() *Config {

	return &Config{
		DB: DatabaseConfig{
			Host:     utils.GetEnv("DB_HOST", "localhost"),
			Port:     utils.GetEnv("DB_PORT", "3306"),
			User:     utils.GetEnv("DB_USER", "root"),
			Password: utils.GetEnv("DB_PASSWORD", ""),
			DBName:   utils.GetEnv("DB_NAME", "go_api"),
		},
	}
}

func (c *Config) DNS() string {
    return fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        c.DB.User,
        c.DB.Password,
        c.DB.Host,
        c.DB.Port,
        c.DB.DBName,
    )
}


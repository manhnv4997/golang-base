package mysql

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadConfig() Config {
	envPath := "env/.env.mysql"

	if err := godotenv.Load(envPath); err != nil {
		log.Println("Không tìm thấy file .env hoặc lỗi khi tải.")
	} else {
		_ = godotenv.Load(envPath)
	}

	return Config{
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", ""),
		Host:     getEnv("DB_HOST", "127.0.0.1"),
		Port:     getEnv("DB_PORT", "3306"),
		DBName:   getEnv("DB_NAME", "your_db_name"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

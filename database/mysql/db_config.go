package mysql

import (
	"demo/app/utils"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadDBConfig() DBConfig {
	utils.LoadEnv()

	return DBConfig{
		User:     utils.GetEnv("DB_USER", "root"),
		Password: utils.GetEnv("DB_PASSWORD", ""),
		Host:     utils.GetEnv("DB_HOST", "127.0.0.1"),
		Port:     utils.GetEnv("DB_PORT", "3306"),
		DBName:   utils.GetEnv("DB_NAME", "your_db_name"),
	}
}

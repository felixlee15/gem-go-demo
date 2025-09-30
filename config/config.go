package config

import "os"

type DBConfig struct {
	Host         string
	Port         string
	Name         string
	Username     string
	Password     string
	SSLMode      string
	MaxOpenConns *int
	MaxIdleConns *int
}

func GetDBConfig() DBConfig {
	return DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		Name:     getEnv("DB_NAME", "mydb"),
		Username: getEnv("DB_USERNAME", "myuser"),
		Password: getEnv("DB_PASSWORD", "mypassword"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func (c DBConfig) DSN() string {
	return "host=" + c.Host +
		" port=" + c.Port +
		" user=" + c.Username +
		" password=" + c.Password +
		" dbname=" + c.Name +
		" sslmode=" + c.SSLMode
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

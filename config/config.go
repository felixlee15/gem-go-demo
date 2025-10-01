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
		Host:     GetEnv("DB_HOST", "localhost"),
		Port:     GetEnv("DB_PORT", "5432"),
		Name:     GetEnv("DB_NAME", "mydb"),
		Username: GetEnv("DB_USERNAME", "myuser"),
		Password: GetEnv("DB_PASSWORD", "mypassword"),
		SSLMode:  GetEnv("DB_SSLMODE", "disable"),
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

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

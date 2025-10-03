package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	App      appConfig
	Database *DBConfig
}

type appConfig struct {
	Port               string
	TrustedOriginRegex string
}

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

var (
	AppConfigs *config
)

func init() {
	_ = godotenv.Load()

	AppConfigs = &config{
		App: appConfig{
			TrustedOriginRegex: os.Getenv("TRUSTED_ORIGIN_REGEX"),
			Port:               os.Getenv("PORT"),
		},
		Database: getDBConfig(),
	}
}

func getDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
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

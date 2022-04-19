package config

import (
	"github.com/spf13/viper"
)

type App struct {
	// Name is Application Name for this project
	Name string `json:"name"`
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string `json:"http_port"`
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string `json:"grpc_port"`
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int `json:"log_level"`
	// LogTimeFormat is print time format for logger e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string `json:"time_format"`
	// ContextTimeout is time limit on an event taking place
	ContextTimeout int `json:"context_timeout "`
	// JWTSecret is a private jwt secret key
	JWTSecret string `json:"jwt_secret"`
}

type Database struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Driver   string `json:"driver"`
}

type Config struct {
	App App
	db  Database
}

// LoadConfiguration will initialize fixed value for config
func NewConfig() Config {
	return Config{
		App: App{
			Name:           viper.GetString("APP_NAME"),
			HTTPPort:       viper.GetString("HTTP_PORT"),
			GRPCPort:       viper.GetString("GRPC_PORT"),
			LogLevel:       viper.GetInt("LOG_LEVEL"),
			LogTimeFormat:  viper.GetString("LOG_TIME_FORMAT"),
			ContextTimeout: viper.GetInt("CONTEXT_TIMEOUT"),
			JWTSecret:      viper.GetString("APP_JWT_SECRET"),
		},
		db: Database{
			Name:     viper.GetString("DB_NAME"),
			Host:     viper.GetString("DB_HOST"),
			User:     viper.GetString("DB_USER"),
			Port:     viper.GetString("DB_PORT"),
			Password: viper.GetString("DB_PASSWORD"),
			Driver:   viper.GetString("DRIVER"),
		},
	}
}

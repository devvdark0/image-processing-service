package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	App  AppConfig `yaml:"app"`
	DB   DBConfig
	Auth AuthConfig
}

type AppConfig struct {
	Env  string `yaml:"env"`
	Port string `yaml:"port"`
}

type DBConfig struct {
	Name     string `yaml:"db_name"`
	User     string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	Host     string `yaml:"db_host"`
	Port     string `yaml:"db_port"`
	SSLMode  string `yaml:"sslmode"`
}

type AuthConfig struct {
	SecretKey       string        `env:"SECRET_KEY" env-required:"true"`
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl"`
}

func MustLoad(configPath string) (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}

	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.Auth.SecretKey = os.Getenv("SECRET_KEY")

	return &cfg, nil
}

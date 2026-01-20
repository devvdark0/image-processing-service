package config

import (
	"fmt"
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
	Name     string `env:"DB_NAME" env-required:"true"`
	User     string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     string `env:"DB_PORT" env-required:"true"`
	SSLMode  string `env:"SSL_MODE" env-required:"true"`
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
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Port = os.Getenv("DB_PORT")
	cfg.DB.Name = os.Getenv("DB_NAME")
	cfg.DB.SSLMode = os.Getenv("SSL_MODE")
	cfg.Auth.SecretKey = os.Getenv("SECRET_KEY")

	return &cfg, nil
}

func (d *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		d.Host,
		d.Port,
		d.Name,
		d.User,
		d.Password,
		d.SSLMode,
	)
}

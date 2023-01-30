package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// API Configure Struct
type ApiConfig struct {
	ApiPort string
	ApiHost string
}

// DB Configure Struct
type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

// Token Config
type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

// Configure struct for API config and DB Config
type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
}

// Func for read config envirovment
func (c Config) readConfig() Config {
	c.DbConfig = DbConfig{
		// Get DB connection
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		// Driver:   os.Getenv("DB_DRIVER"),
	}
	// Set API connection
	c.ApiConfig = ApiConfig{
		ApiPort: "8889",
		ApiHost: "localhost",
	}
	// Set Token For login
	c.TokenConfig = TokenConfig{
		ApplicationName:     "INV_COINS",
		JwtSignatureKey:     "P@ssword",
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: 120 * time.Second,
	}
	return c
}

// For new Config
func NewConfig() Config {
	conf := Config{}
	return conf.readConfig()
}

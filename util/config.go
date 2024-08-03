package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment            string        `mapstructure:"ENVIRONMENT"`
	DBSource               string        `mapstructure:"DB_SOURCE"`
	MigrationURL           string        `mapstructure:"MIGRATION_URL"`
	RedisAddress           string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress      string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	TokenSymmetricKey      string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration    time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration   time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName        string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress     string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword    string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	AccessTokenCookieName  string        `mapstructure:"ACCESS_TOKEN_COOKIE_NAME"`
	JwtSecretKey           string        `mapstructure:"JWT_SECRET_KEY"`
	RefreshTokenCookieName string        `mapstructure:"REFRESH_TOKEN_COOKIE_NAME"`
	JwtRefreshSecretKey    string        `mapstructure:"JWT_REFRESH_SECRET_KEY"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

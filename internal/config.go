package internal

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBConn              string        `mapstructure:"DB_CONN"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	PostgresDB          string        `mapstructure:"POSTGRES_DB"`
	PostgresUser        string        `mapstructure:"POSTGRES_USER"`
	PostgresPassword    string        `mapstructure:"POSTGRES_PASSWORD"`
	ServerAddress       string        `mapstructure:"SERVER_ADD"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
}

func LoadConfig(p string) (config Config, err error) {
	viper.AddConfigPath(p)
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

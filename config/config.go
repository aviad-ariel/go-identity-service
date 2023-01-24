package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	Env *Config
)

type Config struct {
	DBConnectionPrefix     string `mapstructure:"DB_CONNECTION_PREFIX"`
	DBConnectionSuffix     string `mapstructure:"DB_CONNECTION_SUFFIX"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPassword             string `mapstructure:"DB_PASSWORD"`
	DBName                 string `mapstructure:"DB_NAME"`
	UsersCollectionName    string `mapstructure:"USERS_COLLECTION_NAME"`
	JwtSecret              string `mapstructure:"JWT_SECRET"`
	TokenExpirationInHours string `mapstructure:"TOKEN_EXPIRATIONS_IN_HOURS"`
	Port                   string `mapstructure:"PORT"`
}

func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	marshalError := viper.Unmarshal(&Env)
	if marshalError != nil {
		log.Fatal("cannot marshal config:", marshalError)
	}
}

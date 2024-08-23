package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type EnvModel struct {
	Port       string `mapstructure:"PORT"`
	DbName     string `mapstructure:"POSTGRES_DATABASE"`
	DbUser     string `mapstructure:"POSTGRES_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	DBUrl      string `mapstructure:"POSTGRES_URL"`
}

func InitConfig() (configs *EnvModel) {

	if os.Getenv("VERCEL") != "" {
		viper.AutomaticEnv() // Automatically read environment variables

		configs = &EnvModel{
			Port:       viper.GetString("PORT"),
			DbName:     viper.GetString("POSTGRES_DATABASE"),
			DbUser:     viper.GetString("POSTGRES_USER"),
			DbPassword: viper.GetString("DB_PASSWORD"),
			DbHost:     viper.GetString("DB_HOST"),
			DbPort:     viper.GetString("DB_PORT"),
			JWTSecret:  viper.GetString("JWT_SECRET"),
			DBUrl:      viper.GetString("POSTGRES_URL"),
		}
	} else {
		// Fallback for local development
		viper.AddConfigPath("./")
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Error loading env variables", err)
		}

		if err := viper.Unmarshal(&configs); err != nil {
			log.Fatal("Error while unmarshalling loaded variables into struct", err)
		}
	}

	return configs
}

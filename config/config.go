package config

import "os"

type EnvModel struct {
	Port       string `mapstructure:"PORT"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
}

func InitConfig() (configs *EnvModel) {

	// viper.AddConfigPath("./")
	// viper.SetConfigFile(".env")
	// viper.SetConfigType("env")

	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatal("Error loading env env variables", err)
	// }

	// if err := viper.Unmarshal(&configs); err != nil {
	// 	log.Fatal("Error while unmarshalling loaded variables into struct")
	// }

	configs.Port = os.Getenv("PORT")
	configs.DbName = os.Getenv("POSTGRES_DATABASE")
	configs.DbUser = os.Getenv("POSTGRES_USER")
	configs.DbPassword = os.Getenv("POSTGRES_PASSWORD")
	configs.DbHost = os.Getenv("POSTGRES_HOST")
	configs.DbPort = os.Getenv("DB_PORT")
	configs.JWTSecret = os.Getenv("JWT_SECRET")

	return
}

package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBHost string         `mapstructure:"DB_HOST"`
	DBPort string         `mapstructure:"DB_PORT"`
	DBUser string         `mapstructure:"DB_USER"`
	DBPass string         `mapstructure:"DB_PASS"`
	DBName string         `mapstructure:"DB_NAME"`
	Clerk_Secret string   `mapstructure:"CLERK_SECRET_KEY"`
}

func NewEnv() *Env {
	env := Env {}
	viper.SetConfigFile(".env")
	
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal(err)
	}

	return &env
}
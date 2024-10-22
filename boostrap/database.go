package bootstrap

import (
	"fmt"
	"log"

	"github.com/d1nnn/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresqlDatabase(env *Env) *gorm.DB {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	conString := fmt.Sprintf("postgresql://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&domain.Transaction{}, &domain.AppUser{}, &domain.Preset{})

	return db
}

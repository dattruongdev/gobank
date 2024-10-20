package repository

import (
	"context"
	"fmt"

	"github.com/d1nnn/domain"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) domain.UserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

	func(ur *PostgresUserRepository) Create(c context.Context, user domain.AppUser) error {


		tx := ur.db.Save(&user)

		return tx.Error
	}

	func(ur *PostgresUserRepository) Update(c context.Context, updateRequest domain.AppUser) error {
		tx := ur.db.Save(&updateRequest)

		return tx.Error
	}

	func(ur *PostgresUserRepository) GetAll(c context.Context, userId string) ([]domain.AppUser, error) {
		var users []domain.AppUser
		tx := ur.db.Preload("Transactions").Where("id <> ?", userId).Find(&users)

		return users, tx.Error
	}

	func(ur *PostgresUserRepository) GetByEmail(c context.Context, email string) (domain.AppUser, error) {
		user := domain.AppUser {
			Email: email,
		}

		tx := ur.db.Find(&user)

		return user, tx.Error
	}

	func(ur *PostgresUserRepository) GetByName(c context.Context, name string) ([]domain.AppUser, error) {
		var users []domain.AppUser
		tx := ur.db.Preload("Transactions").Where("fullname like ?", fmt.Sprintf("%%%v%%", name)).Find(&users)

		return users, tx.Error
	}

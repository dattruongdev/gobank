package repository

import (
	"context"

	"github.com/d1nnn/domain"
	"gorm.io/gorm"
)

type PostgresCategoryRepository struct {
	db *gorm.DB
}

func NewPostgresCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &PostgresCategoryRepository{
		db: db,
	}
}

func (cr *PostgresCategoryRepository) Create(c context.Context, category domain.Category) error {
	tx := cr.db.Create(&category)

	return tx.Error
}
func (cr *PostgresCategoryRepository) GetByName(c context.Context, catName string) (domain.Category, error) {
	category := domain.Category {
		Name: catName,
	}
	tx :=cr.db.Find(&category)

	return category, tx.Error
}
func (cr *PostgresCategoryRepository) GetAll(c context.Context) ([]domain.Category, error) {
	var categories []domain.Category
	tx := cr.db.Find(&categories)

	return categories, tx.Error
}
func (cr *PostgresCategoryRepository) Update(c context.Context, updateRequest domain.Category) error {
	category, err := cr.GetByName(c, updateRequest.Name)
	if err != nil {
		return err
	}

	updateRequest.ID = category.ID
	tx := cr.db.Save(&updateRequest)

	return tx.Error
}


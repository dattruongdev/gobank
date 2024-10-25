package repository

import (
	"context"

	"github.com/d1nnn/domain"
	"gorm.io/gorm"
)

type PostgresPresetRepository struct {
	db *gorm.DB
}

func NewPostgresPresetRepository(db *gorm.DB) domain.PresetRepository {
	return &PostgresPresetRepository{
		db: db,
	}
}

func (pr *PostgresPresetRepository) Create(c context.Context, p domain.Preset) error {
	tx := pr.db.Save(&p)

	return tx.Error
}
func (pr *PostgresPresetRepository) GetAll(c context.Context, userId string) ([]domain.Preset, error) {
	var presets []domain.Preset
	tx := pr.db.Preload("Payee").Where("payer_id = ?", userId).Find(&presets)
	return presets, tx.Error
}

func (pr *PostgresPresetRepository) Delete(c context.Context, payerId string, payeeIds ...string) error {
	tx := pr.db.Where("payee_id IN (?) and payer_id = ?", payeeIds, payerId).Delete(&domain.Preset{})

	return tx.Error
}

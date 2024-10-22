package usecase

import (
	"context"

	"github.com/d1nnn/domain"
)

type PresetUsecase struct {
	presetRepository domain.PresetRepository
}

func NewPresetUsecase(repo domain.PresetRepository) *PresetUsecase {
	return &PresetUsecase {
		presetRepository: repo,
	}
}

func (pu *PresetUsecase) GetAll(c context.Context, userId string) ([]domain.Preset, error) {
	return pu.presetRepository.GetAll(c, userId);
}

func (pu *PresetUsecase) Create(c context.Context, preset domain.Preset) error {
	return pu.presetRepository.Create(c, preset)
}

